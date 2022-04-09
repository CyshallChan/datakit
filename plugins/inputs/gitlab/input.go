// Package gitlab collect GitLab metrics
package gitlab

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	dhttp "gitlab.jiagouyun.com/cloudcare-tools/datakit/http"
	ihttp "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/http"
	iod "gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var _ inputs.ElectionInput = (*Input)(nil)

const (
	inputName         = "gitlab"
	catalog           = "gitlab"
	pipelineHook      = "Pipeline Hook"
	jobHook           = "Job Hook"
	gitlabEventHeader = "X-Gitlab-Event"
	sampleCfg         = `
[[inputs.gitlab]]
    ## param type: string - default: http://127.0.0.1:80/-/metrics
    prometheus_url = "http://127.0.0.1:80/-/metrics"

    ## param type: string - optional: time units are "ms", "s", "m", "h" - default: 10s
    interval = "10s"

    enable_ci_visibility = true

    [inputs.gitlab.tags]
    # some_tag = "some_value"
    # more_tag = "some_other_value"
`
)

var l = logger.DefaultSLogger(inputName)

func init() { //nolint:gochecknoinits
	inputs.Add(inputName, func() inputs.Input {
		return newInput()
	})
}

type Input struct {
	URL                string            `toml:"prometheus_url"`
	Interval           string            `toml:"interval"`
	Tags               map[string]string `toml:"tags"`
	EnableCIVisibility bool              `toml:"enable_ci_visibility"`

	httpClient *http.Client
	duration   time.Duration

	pause   bool
	pauseCh chan bool

	semStop *cliutils.Sem // start stop signal
}

func (ipt *Input) RegHTTPHandler() {
	if ipt.EnableCIVisibility {
		l.Infof("start listening gitlab webhooks")
		dhttp.RegHTTPHandler("POST", "/v1/gitlab", ihttp.ProtectedHandlerFunc(ipt.ServeHTTP, l))
	}
}

var maxPauseCh = inputs.ElectionPauseChannelLength

func newInput() *Input {
	return &Input{
		Tags:       make(map[string]string),
		pauseCh:    make(chan bool, maxPauseCh),
		duration:   time.Second * 10,
		httpClient: &http.Client{Timeout: 5 * time.Second},

		semStop: cliutils.NewSem(),
	}
}

func (ipt *Input) Run() {
	l = logger.SLogger(inputName)

	ipt.loadCfg()

	ticker := time.NewTicker(ipt.duration)
	defer ticker.Stop()

	for {
		select {
		case <-datakit.Exit.Wait():
			l.Info("exit")
			return

		case <-ipt.semStop.Wait():
			l.Info("gitlab return")
			return

		case <-ticker.C:
			if ipt.pause {
				l.Debugf("not leader, skipped")
				continue
			}
			ipt.gather()

		case ipt.pause = <-ipt.pauseCh:
			// nil
		}
	}
}

func (ipt *Input) Terminate() {
	if ipt.semStop != nil {
		ipt.semStop.Close()
	}
}

func (ipt *Input) Pause() error {
	tick := time.NewTicker(inputs.ElectionPauseTimeout)
	defer tick.Stop()
	select {
	case ipt.pauseCh <- true:
		return nil
	case <-tick.C:
		return fmt.Errorf("pause %s failed", inputName)
	}
}

func (ipt *Input) Resume() error {
	tick := time.NewTicker(inputs.ElectionResumeTimeout)
	defer tick.Stop()
	select {
	case ipt.pauseCh <- false:
		return nil
	case <-tick.C:
		return fmt.Errorf("resume %s failed", inputName)
	}
}

func (ipt *Input) loadCfg() {
	dur, err := time.ParseDuration(ipt.Interval)
	if err != nil {
		l.Warnf("parse interval error (use default 10s): %s", err)
		return
	}
	ipt.duration = dur
}

func (ipt *Input) gather() {
	start := time.Now()

	pts, err := ipt.gatherMetrics()
	if err != nil {
		l.Error(err)
		return
	}

	if err := iod.Feed(inputName, datakit.Metric, pts, &iod.Option{CollectCost: time.Since(start)}); err != nil {
		l.Error(err)
	}
}

func (ipt *Input) gatherMetrics() ([]*iod.Point, error) {
	resp, err := ipt.httpClient.Get(ipt.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	metrics, err := promTextToMetrics(resp.Body)
	if err != nil {
		return nil, err
	}

	var points []*iod.Point

	for _, m := range metrics {
		measurement := inputName

		// 非常粗暴的筛选方式
		if len(m.tags) == 0 {
			measurement = inputName + "_base"
		}
		if _, ok := m.tags["method"]; ok {
			measurement = inputName + "_http"
		}

		for k, v := range ipt.Tags {
			m.tags[k] = v
		}

		point, err := iod.MakePoint(measurement, m.tags, m.fields)
		if err != nil {
			l.Warn(err)
			continue
		}
		points = append(points, point)
	}

	return points, nil
}

func (*Input) SampleConfig() string { return sampleCfg }

func (*Input) Catalog() string { return catalog }

func (*Input) AvailableArchs() []string { return datakit.AllArch }

func (*Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{
		&gitlabMeasurement{},
		&gitlabBaseMeasurement{},
		&gitlabHTTPMeasurement{},
		&gitlabPipelineMeasurement{},
		&gitlabJobMeasurement{},
	}
}
