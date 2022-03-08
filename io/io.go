// Package io implements datakits data transfer among inputs.
package io

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/dataway"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/sender"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/sink/sinkcommon"
)

const (
	minGZSize   = 1024
	maxKodoPack = 10 * 1000 * 1000
)

var (
	testAssert                 = false
	highFreqCleanInterval      = time.Millisecond * 500
	datawayListIntervalDefault = 50
	heartBeatIntervalDefault   = 40
	log                        = logger.DefaultSLogger("io")

	DisableLogFilter            bool
	DisableHeartbeat            bool
	DisableDatawayList          bool
	FlagDebugDisableDatawayList bool
)

type Option struct {
	CollectCost time.Duration
	HighFreq    bool
	Version     string
	HTTPHost    string
	PostTimeout time.Duration
	Sample      func(points []*Point) []*Point
}

type lastError struct {
	from, err string
	ts        time.Time
}

func (e *lastError) Error() string {
	return fmt.Sprintf("%s [%s] %s", e.ts, e.from, e.err)
}

func NewLastError(from, err string) *lastError {
	return &lastError{
		from: from,
		err:  err,
		ts:   time.Now(),
	}
}

type IO struct {
	FeedChanSize              int
	HighFreqFeedChanSize      int
	MaxCacheCount             int64
	CacheDumpThreshold        int64
	MaxDynamicCacheCount      int64
	DynamicCacheDumpThreshold int64
	FlushInterval             time.Duration
	OutputFile                string
	OutputFileInput           []string
	EnableCache               bool

	dw *dataway.DataWayCfg

	in        chan *iodata
	in2       chan *iodata // high-freq chan
	inLastErr chan *lastError

	SentBytes int

	inputstats map[string]*InputsStat
	lock       sync.RWMutex

	cache        map[string][]*Point
	dynamicCache map[string][]*Point

	fd *os.File

	cacheCnt        int64
	dynamicCacheCnt int64
	droppedTotal    int64
	outputFileSize  int64
	sender          *sender.Sender
}

type IoStat struct {
	SentBytes int `json:"sent_bytes"`
}

func NewIO() *IO {
	x := &IO{
		FeedChanSize:         1024,
		HighFreqFeedChanSize: 2048,
		MaxCacheCount:        1024,
		MaxDynamicCacheCount: 1024,
		FlushInterval:        10 * time.Second,
		in:                   make(chan *iodata, 128),
		in2:                  make(chan *iodata, 128*8),
		inLastErr:            make(chan *lastError, 128),

		inputstats: map[string]*InputsStat{},

		cache:        map[string][]*Point{},
		dynamicCache: map[string][]*Point{},
	}

	log.Debugf("IO: %+#v", x)

	return x
}

type iodata struct {
	category, name string
	opt            *Option
	pts            []*Point
}

func TestOutput() {
	testAssert = true
}

func SetTest() {
	testAssert = true
}

//nolint:gocyclo
func (x *IO) DoFeed(pts []*Point, category, name string, opt *Option) error {
	if testAssert {
		return nil
	}

	ch := x.in
	if opt != nil && opt.HighFreq {
		ch = x.in2
	}

	switch category {
	case datakit.MetricDeprecated:
	case datakit.Metric:
	case datakit.Network:
	case datakit.KeyEvent:
	case datakit.Object:
	case datakit.CustomObject:
	case datakit.Logging:
		if x.dw.ClientsCount() == 1 {
			if !DisableLogFilter {
				pts = defLogfilter.filter(pts)
			}
		} else {
			// TODO: add multiple dataway config support
			log.Infof("multiple dataway config %d for log filter not support yet", x.dw.ClientsCount())
		}
	case datakit.Tracing:
	case datakit.Security:
	case datakit.RUM:
	default:
		return fmt.Errorf("invalid category `%s'", category)
	}

	log.Debugf("io feed %s", name)

	select {
	case ch <- &iodata{
		category: category,
		pts:      pts,
		name:     name,
		opt:      opt,
	}:
	case <-datakit.Exit.Wait():
		log.Warnf("%s/%s feed skipped on global exit", category, name)
	}

	return nil
}

func (x *IO) ioStop() {
	if x.fd != nil {
		if err := x.fd.Close(); err != nil {
			log.Error(err)
		}
	}
	// stop sender
	if err := x.sender.Stop(); err != nil {
		log.Error(err)
	}
}

func (x *IO) updateLastErr(e *lastError) {
	x.lock.Lock()
	defer x.lock.Unlock()

	stat, ok := x.inputstats[e.from]
	if !ok {
		stat = &InputsStat{
			First: time.Now(),
			Last:  time.Now(),
		}
		x.inputstats[e.from] = stat
	}

	stat.LastErr = e.err
	stat.LastErrTS = e.ts
}

func (x *IO) updateStats(d *iodata) {
	now := time.Now()
	stat, ok := x.inputstats[d.name]

	if !ok {
		stat = &InputsStat{
			Total: int64(len(d.pts)),
			First: now,
		}
		x.inputstats[d.name] = stat
	}

	stat.Total += int64(len(d.pts))
	stat.Count++
	stat.Last = now
	stat.Category = d.category

	if (stat.Last.Unix() - stat.First.Unix()) > 0 {
		stat.Frequency = fmt.Sprintf("%.02f/min",
			float64(stat.Count)/(float64(stat.Last.Unix()-stat.First.Unix())/60))
	}
	stat.AvgSize = (stat.Total) / stat.Count

	if d.opt != nil {
		stat.Version = d.opt.Version
		stat.totalCost += d.opt.CollectCost
		stat.AvgCollectCost = (stat.totalCost) / time.Duration(stat.Count)
		if d.opt.CollectCost > stat.MaxCollectCost {
			stat.MaxCollectCost = d.opt.CollectCost
		}
	}
}

func (x *IO) ifMatchOutputFileInput(feedName string) bool {
	for _, v := range x.OutputFileInput {
		if v == feedName {
			return true
		}
	}
	return false
}

func (x *IO) cacheData(d *iodata, tryClean bool) {
	if d == nil {
		log.Warn("get empty data, ignored")
		return
	}

	log.Debugf("get iodata(%d points) from %s|%s", len(d.pts), d.category, d.name)

	x.updateStats(d)

	if d.opt != nil && d.opt.HTTPHost != "" {
		x.dynamicCache[d.opt.HTTPHost] = append(x.dynamicCache[d.opt.HTTPHost], d.pts...)
		x.dynamicCacheCnt += int64(len(d.pts))
	} else {
		x.cache[d.category] = append(x.cache[d.category], d.pts...)
		x.cacheCnt += int64(len(d.pts))
	}

	if x.OutputFile != "" {
		bodies, err := x.buildBody(d.pts)
		if err != nil {
			log.Errorf("build iodata bodies failed: %s", err)
		}
		for _, body := range bodies {
			if len(x.OutputFileInput) == 0 || x.ifMatchOutputFileInput(d.name) {
				if err := x.fileOutput(body.buf); err != nil {
					log.Error("fileOutput: %s, ignored", err.Error())
				}
			}
		}
	}

	if (tryClean && x.MaxCacheCount > 0 && x.cacheCnt > x.MaxCacheCount) ||
		(x.MaxDynamicCacheCount > 0 && x.dynamicCacheCnt > x.MaxDynamicCacheCount) {
		x.flushAll()
	}
}

func (x *IO) cleanHighFreqIOData() {
	if len(x.in2) > 0 {
		log.Debugf("clean %d cache on high-freq-chan", len(x.in2))
	}

	for {
		select {
		case d := <-x.in2: // eat all cached data
			x.cacheData(d, false)
		default:
			return
		}
	}
}

func (x *IO) init() error {
	if x.OutputFile != "" {
		f, err := os.OpenFile(x.OutputFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o644) //nolint:gosec
		if err != nil {
			log.Error(err)
			return err
		}

		x.fd = f
	}

	return nil
}

func (x *IO) StartIO(recoverable bool) {
	if sender, err := sender.NewSender(
		&sender.Option{
			Cache:              x.EnableCache,
			FlushCacheInterval: x.FlushInterval,
			Write:              x.dw.Write,
			ErrorCallback: func(err error) {
				addReporter(Reporter{Status: "error", Message: err.Error()})
			},
		}); err != nil {
		log.Errorf("init sender error: %s", err.Error())
	} else {
		x.sender = sender
	}

	g := datakit.G("io")
	g.Go(func(ctx context.Context) error {
		if err := x.init(); err != nil {
			log.Errorf("init io err %v", err)
			return nil
		}

		defer x.ioStop()

		tick := time.NewTicker(x.FlushInterval)
		defer tick.Stop()

		highFreqRecvTicker := time.NewTicker(highFreqCleanInterval)
		defer highFreqRecvTicker.Stop()

		heartBeatTick := time.NewTicker(time.Second * time.Duration(heartBeatIntervalDefault))
		defer heartBeatTick.Stop()

		datawaylistTick := time.NewTicker(time.Second * time.Duration(datawayListIntervalDefault))
		defer datawaylistTick.Stop()

		for {
			select {
			case d := <-x.in:
				x.cacheData(d, true)

			case e := <-x.inLastErr:
				x.updateLastErr(e)

			case <-highFreqRecvTicker.C:
				x.cleanHighFreqIOData()

			case <-heartBeatTick.C:
				log.Debugf("### enter heartBeat")
				if !DisableHeartbeat {
					heartBeatInterval, err := x.dw.HeartBeat()
					if err != nil {
						log.Warnf("dw.HeartBeat: %s, ignored", err.Error())
					}
					if heartBeatInterval != heartBeatIntervalDefault {
						heartBeatTick.Reset(time.Second * time.Duration(heartBeatInterval))
						heartBeatIntervalDefault = heartBeatInterval
					}
				}

			case <-datawaylistTick.C:
				log.Debugf("### enter dataway list")
				if !DisableDatawayList {
					var dws []string
					var err error
					var datawayListInterval int
					dws, datawayListInterval, err = x.dw.DatawayList()
					if err != nil {
						log.Warnf("DatawayList(): %s, ignored", err)
					}
					dataway.AvailableDataways = dws
					if datawayListInterval != datawayListIntervalDefault {
						datawaylistTick.Reset(time.Second * time.Duration(datawayListInterval))
						datawayListIntervalDefault = datawayListInterval
					}
				}

			case <-tick.C:
				x.flushAll()

			case <-datakit.Exit.Wait():
				log.Info("io exit on exit")
				return nil
			}
		}
	})

	// start log filter
	if !DisableLogFilter {
		defLogfilter.start()
	}

	log.Info("starting...")
}

func (x *IO) flushAll() {
	x.flush()

	if x.cacheCnt > 0 {
		log.Warnf("post failed cache count: %d", x.cacheCnt)
	}

	// dump cache pts
	if x.CacheDumpThreshold > 0 && x.cacheCnt > x.CacheDumpThreshold {
		log.Warnf("failed cache count reach max limit(%d), cleanning cache...", x.MaxCacheCount)
		for k := range x.cache {
			x.cache[k] = nil
		}
		x.droppedTotal += x.cacheCnt
		x.cacheCnt = 0
	}
	// dump dynamic cache pts
	if x.DynamicCacheDumpThreshold > 0 && x.dynamicCacheCnt > x.DynamicCacheDumpThreshold {
		log.Warnf("failed dynamicCache count reach max limit(%d), cleanning cache...", x.MaxDynamicCacheCount)
		for k := range x.dynamicCache {
			x.dynamicCache[k] = nil
		}
		x.droppedTotal += x.dynamicCacheCnt
		x.dynamicCacheCnt = 0
	}
}

func (x *IO) flush() {
	for k, v := range x.cache {
		if err := x.doFlush(v, k); err != nil {
			log.Errorf("post %d to %s failed", len(v), k)
			continue
		}

		if len(v) > 0 {
			x.cacheCnt -= int64(len(v))
			log.Debugf("clean %d cache on %s, remain: %d", len(v), k, x.cacheCnt)
			x.cache[k] = nil
		}
	}

	// flush dynamic cache: __not__ post to default dataway
	for k, v := range x.dynamicCache {
		if err := x.doFlush(v, k); err != nil {
			log.Errorf("post %d to %s failed", len(v), k)
			continue
		}

		if len(v) > 0 {
			x.dynamicCacheCnt -= int64(len(v))
			log.Debugf("clean %d dynamicCache on %s, remain: %d", len(v), k, x.dynamicCacheCnt)
			x.dynamicCache[k] = nil
		}
	}
}

type body struct {
	buf  []byte
	gzon bool
}

var lines = bytes.Buffer{}

func (x *IO) buildBody(pts []*Point) ([]*body, error) {
	var (
		gz = func(lines []byte) (*body, error) {
			var (
				body = &body{buf: lines}
				err  error
			)
			log.Debugf("### io body size before GZ: %dM %dK", len(body.buf)/1000/1000, len(body.buf)/1000)
			if len(lines) > minGZSize && x.OutputFile == "" {
				if body.buf, err = datakit.GZip(body.buf); err != nil {
					log.Errorf("gz: %s", err.Error())

					return nil, err
				}
				body.gzon = true
			}

			return body, nil
		}
		// lines  bytes.Buffer
		bodies []*body
	)
	lines.Reset()
	for _, pt := range pts {
		ptstr := pt.String()
		if lines.Len()+len(ptstr)+1 >= maxKodoPack {
			if body, err := gz(lines.Bytes()); err != nil {
				return nil, err
			} else {
				bodies = append(bodies, body)
			}
			lines.Reset()
		}
		lines.WriteString(ptstr)
		lines.WriteString("\n")
	}
	if body, err := gz(lines.Bytes()); err != nil {
		return nil, err
	} else {
		return append(bodies, body), nil
	}
}

func (x *IO) doFlush(pts []*Point, category string) error {
	if x.sender == nil {
		return fmt.Errorf("io sender is not initialized")
	}

	points := []sinkcommon.ISinkPoint{}

	for _, pt := range pts {
		points = append(points, pt)
	}

	return x.sender.Write(category, points)
}

func (x *IO) fileOutput(body []byte) error {
	if _, err := x.fd.Write(append(body, '\n')); err != nil {
		return err
	}

	x.outputFileSize += int64(len(body))
	if x.outputFileSize > 4*1024*1024 {
		if err := x.fd.Truncate(0); err != nil {
			return err
		}
		x.outputFileSize = 0
	}

	return nil
}

func (x *IO) DroppedTotal() int64 {
	// NOTE: not thread-safe
	return x.droppedTotal
}
