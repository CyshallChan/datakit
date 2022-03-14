// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package tailer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/encoding"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/multiline"
	dkio "gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/worker"
)

const (
	defaultSleepDuration = time.Second
	readBuffSize         = 1024 * 4
	timeoutDuration      = time.Second * 3

	firstMessage = "[DataKit-logging] First Message. filename: %s, source: %s"
)

type Single struct {
	opt                    *Option
	file                   *os.File
	filename, baseFilename string

	decoder *encoding.Decoder
	mult    *multiline.Multiline

	readBuff []byte

	tags   map[string]string
	stopCh chan struct{}
}

func NewTailerSingle(filename string, opt *Option) (*Single, error) {
	if opt == nil {
		return nil, fmt.Errorf("option cannot be null pointer")
	}

	t := &Single{
		stopCh: make(chan struct{}, 1),
		opt:    opt,
	}

	var err error
	if opt.CharacterEncoding != "" {
		t.decoder, err = encoding.NewDecoder(opt.CharacterEncoding)
		if err != nil {
			return nil, err
		}
	}
	t.mult, err = multiline.New(opt.MultilineMatch, opt.MultilineMaxLines)
	if err != nil {
		return nil, err
	}

	t.file, err = os.Open(filename) //nolint:gosec
	if err != nil {
		return nil, err
	}

	if !opt.FromBeginning {
		if _, err := t.file.Seek(0, os.SEEK_END); err != nil {
			return nil, err
		}
	}

	t.readBuff = make([]byte, readBuffSize)
	t.filename = t.file.Name()
	t.baseFilename = filepath.Base(t.filename)
	t.tags = t.buildTags(opt.GlobalTags)

	return t, nil
}

func (t *Single) Run() {
	defer t.Close()
	t.forwardMessage()
}

func (t *Single) Close() {
	t.stopCh <- struct{}{}
	t.opt.log.Infof("closing %s", t.filename)
}

//nolint:cyclop
func (t *Single) forwardMessage() {
	var (
		b       = &buffer{}
		timeout = time.NewTicker(timeoutDuration)
		lines   []string
		readNum int
		err     error
	)
	defer timeout.Stop()

	if !t.opt.DisableSendEvent {
		dkio.FeedEventLog(&dkio.Reporter{Message: fmt.Sprintf(firstMessage, t.filename, t.opt.Source), Logtype: "event"})
	}
	for {
		select {
		case <-t.stopCh:
			if err := t.file.Close(); err != nil {
				t.opt.log.Warnf("Close(): %s, ignored", err.Error())
			}
			t.opt.log.Infof("stop reading data from file %s", t.filename)
			return
		case <-timeout.C:
			if str := t.mult.FlushString(); str != "" {
				t.send(str)
			}
		default:
			// nil
		}

		b.buf, readNum, err = t.read()
		if err != nil {
			t.opt.log.Warnf("failed to read data from file %s, error: %s", t.filename, err)
			return
		}
		if readNum == 0 {
			t.wait()
			continue
		}

		lines = b.split()
		var pending []worker.TaskData
		for _, line := range lines {
			if line == "" {
				continue
			}

			var text string
			text, err = t.decode(line)
			if err != nil {
				t.opt.log.Debugf("decode '%s' error: %s", t.opt.CharacterEncoding, err)
				if !t.opt.DisableSendEvent {
					dkio.FeedEventLog(&dkio.Reporter{Message: line, Logtype: "event", Status: "warning"}) // event:warning
				}
			}

			text = t.multiline(text)
			if text == "" {
				continue
			}

			if t.opt.ForwardFunc != nil {
				t.sendToForwardCallback(text)
				continue
			}
			pending = append(pending, &SocketTaskData{Source: t.opt.Source, Log: text, Tag: t.tags})
		}
		if len(pending) > 0 {
			t.sendToPipeline(pending)
		}
	}
}

func (t *Single) send(text string) {
	if t.opt.ForwardFunc != nil {
		t.sendToForwardCallback(text)
		return
	}
	t.sendToPipeline([]worker.TaskData{&SocketTaskData{Source: t.opt.Source, Log: text, Tag: t.tags}})
}

func (t *Single) sendToForwardCallback(text string) {
	err := t.opt.ForwardFunc(t.baseFilename, text)
	if err != nil {
		t.opt.log.Warnf("failed to forward text from file %s, error: %s", t.filename, err)
	}
}

func (t *Single) sendToPipeline(pending []worker.TaskData) {
	task := &worker.Task{
		TaskName:   "logging/" + t.opt.Pipeline,
		ScriptName: t.opt.Pipeline,
		Source:     t.opt.Source,
		Data:       pending,
		Opt: &worker.TaskOpt{
			IgnoreStatus:          t.opt.IgnoreStatus,
			DisableAddStatusField: t.opt.DisableAddStatusField,
		},
		TS:            time.Now(),
		MaxMessageLen: maxFieldsLength,
	}

	err := worker.FeedPipelineTaskBlock(task)
	if err != nil {
		t.opt.log.Warnf("pipline feed err = %v", err)
		return
	}
}

func (t *Single) currentOffset() int64 {
	if t.file == nil {
		return -1
	}
	offset, err := t.file.Seek(0, os.SEEK_CUR)
	if err != nil {
		return -1
	}
	return offset
}

func (t *Single) read() ([]byte, int, error) {
	n, err := t.file.Read(t.readBuff)
	if err != nil && err != io.EOF {
		// an unexpected error occurred, stop the tailor
		t.opt.log.Warnf("Unexpected error occurred while reading file: %s", err)
		return nil, 0, err
	}

	return t.readBuff[:n], n, nil
}

func (t *Single) wait() {
	time.Sleep(defaultSleepDuration)
}

func (t *Single) buildTags(globalTags map[string]string) map[string]string {
	tags := make(map[string]string)
	for k, v := range globalTags {
		tags[k] = v
	}
	if _, ok := tags["filename"]; !ok {
		tags["filename"] = t.baseFilename
	}
	return tags
}

func (t *Single) decode(text string) (str string, err error) {
	if t.decoder == nil {
		return text, nil
	}
	return t.decoder.String(text)
}

func (t *Single) multiline(text string) string {
	if t.mult == nil {
		return text
	}
	return t.mult.ProcessLineString(text)
}

type buffer struct {
	buf           []byte
	previousBlock []byte
}

func (b *buffer) split() []string {
	// 以换行符 split
	lines := bytes.Split(b.buf, []byte{'\n'})
	if len(lines) == 0 {
		return nil
	}

	var res []string

	// block 不为空时，将其内容添加到 lines 首元素前端
	// block 置空
	if len(b.previousBlock) != 0 {
		lines[0] = append(b.previousBlock, lines[0]...)
		b.previousBlock = b.previousBlock[:0]
	}

	// 当 lines 最后一个元素不为空时，说明这段内容并不包含换行符，将其暂存到 previousBlock
	if len(lines[len(lines)-1]) != 0 {
		// 将 lines 尾元素 append previousBlock，避免占用此 slice 造成内存泄漏
		b.previousBlock = append(b.previousBlock, lines[len(lines)-1]...)
		lines = lines[:len(lines)-1]
	}

	for _, line := range lines {
		res = append(res, string(line))
	}

	return res
}
