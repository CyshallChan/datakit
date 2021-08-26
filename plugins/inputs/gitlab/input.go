package gitlab

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

const (
	inputName = "gitlab"

	sampleCfg = `
[[inputs.gitlab]]
    ## param type: string - default: http://127.0.0.1:80/-/metrics
    prometheus_url = "http://127.0.0.1:80/-/metrics"

    ## param type: string - optional: time units are "ms", "s", "m", "h" - default: 10s
    interval = "10s"

    [inputs.gitlab.tags]
    # some_tag = "some_value"
    # more_tag = "some_other_value"
`
)

var l = logger.DefaultSLogger(inputName)

func init() {
	inputs.Add(inputName, func() inputs.Input {
		return newInput()
	})
}

type Input struct {
	URL      string            `toml:"prometheus_url"`
	Interval string            `toml:"interval"`
	Tags     map[string]string `toml:"tags"`

	httpClient *http.Client
	duration   time.Duration

	paused     bool
	electionCh chan interface{}
}

func newInput() *Input {
	return &Input{
		Tags:     make(map[string]string),
		duration: time.Second * 10,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		electionCh: make(chan interface{}),
	}
}

func (*Input) SampleConfig() string {
	return sampleCfg
}

func (*Input) Catalog() string {
	return "gitlab"
}

func (this *Input) Run() {
	l = logger.SLogger(inputName)

	this.loadCfg()

	ticker := time.NewTicker(this.duration)
	defer ticker.Stop()

	defer close(this.electionCh)

	for {
		select {
		case <-datakit.Exit.Wait():
			l.Info("exit")
			return

		case x := <-this.electionCh:
			if x == PAUSE {
				this.paused = true
			} else {
				this.paused = false
			}

		case <-ticker.C:
			if !this.paused {
				this.gather()
			}
		}
	}
}

const (
	PAUSE  = 0
	RESUME = 1
)

func (this *Input) Pause() error {
	tick := time.NewTicker(time.Second * 3)
	defer tick.Stop()

	select {
	case this.electionCh <- PAUSE:
		return nil
	case <-tick.C:
		return fmt.Errorf("pause %s failed", inputName)
	}
}

func (this *Input) Resume() error {
	tick := time.NewTicker(time.Second * 3)
	defer tick.Stop()

	select {
	case this.electionCh <- RESUME:
		return nil
	case <-tick.C:
		return fmt.Errorf("resume %s failed", inputName)
	}
}

func (this *Input) loadCfg() {
	dur, err := time.ParseDuration(this.Interval)
	if err != nil {
		l.Warnf("parse interval error (use default 10s): %s", err)
		return
	}
	this.duration = dur
}

func (this *Input) gather() {
	start := time.Now()

	pts, err := this.gatherMetrics()
	if err != nil {
		l.Error(err)
		return
	}

	if err := io.Feed(inputName, datakit.Metric, pts, &io.Option{CollectCost: time.Since(start)}); err != nil {
		l.Error(err)
	}
}

func (this *Input) gatherMetrics() ([]*io.Point, error) {
	resp, err := this.httpClient.Get(this.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	metrics, err := promTextToMetrics(resp.Body)
	if err != nil {
		return nil, err
	}

	var points []*io.Point

	for _, m := range metrics {
		var measurement = inputName

		// 非常粗暴的筛选方式
		if len(m.tags) == 0 {
			measurement = inputName + "_base"
		}
		if _, ok := m.tags["method"]; ok {
			measurement = inputName + "_http"
		}

		for k, v := range this.Tags {
			m.tags[k] = v
		}

		point, err := io.MakePoint(measurement, m.tags, m.fields)
		if err != nil {
			l.Warn(err)
			continue
		}
		points = append(points, point)
	}

	return points, nil
}

func (*Input) AvailableArchs() []string {
	return datakit.AllArch
}

func (*Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{
		&gitlabMeasurement{},
		&gitlabBaseMeasurement{},
		&gitlabHTTPMeasurement{},
	}
}

type gitlabMeasurement struct{}
type gitlabBaseMeasurement struct{}
type gitlabHTTPMeasurement struct{}

func (*gitlabMeasurement) LineProto() (*io.Point, error)     { return nil, nil }
func (*gitlabBaseMeasurement) LineProto() (*io.Point, error) { return nil, nil }
func (*gitlabHTTPMeasurement) LineProto() (*io.Point, error) { return nil, nil }

func (*gitlabMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{Name: "gitlab",
		Desc: "GitLab 运行指标",
		Tags: map[string]interface{}{
			"action":           inputs.NewTagInfo("行为"),
			"controller":       inputs.NewTagInfo("管理"),
			"feature_category": inputs.NewTagInfo("类型特征"),
			"storage":          inputs.NewTagInfo("存储"),
		},
		Fields: map[string]interface{}{
			"transaction_cache_read_miss_count_total":             &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.NCount, Desc: "The counter for cache misses for Rails cache calls"},
			"transaction_cache_read_hit_count_total":              &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.NCount, Desc: "The counter for cache hits for Rails cache calls"},
			"transaction_db_count_total":                          &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.NCount, Desc: "The counter for db"},
			"transaction_db_cached_count_total":                   &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.NCount, Desc: "The counter for db cache"},
			"rack_requests_total":                                 &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.UnknownUnit, Desc: "The rack request count"},
			"cache_operations_total":                              &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.UnknownUnit, Desc: "The count of cache access time"},
			"cache_operation_duration_seconds_count":              &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The count of cache access time"},
			"cache_operation_duration_seconds_sum":                &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The count of cache access time"},
			"transaction_view_duration_total":                     &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.UnknownUnit, Desc: "The duration for views"},
			"transaction_new_redis_connections_total":             &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.UnknownUnit, Desc: "The counter for new Redis connections"},
			"sql_duration_seconds_count":                          &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The total SQL execution time, excluding SCHEMA operations and BEGIN / COMMIT"},
			"sql_duration_seconds_sum":                            &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The sum of SQL execution time, excluding SCHEMA operations and BEGIN / COMMIT"},
			"transaction_duration_seconds_count":                  &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The count of duration for all transactions (gitlab_transaction_* metrics)"},
			"transaction_duration_seconds_sum":                    &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The sum of duration for all transactions (gitlab_transaction_* metrics)"},
			"banzai_cacheless_render_real_duration_seconds_count": &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The count of duration of rendering Markdown into HTML when cached output exists"},
			"banzai_cacheless_render_real_duration_seconds_sum":   &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The sum of duration of rendering Markdown into HTML when cached output exists"},
			"cache_misses_total":                                  &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.UnknownUnit, Desc: "The cache read miss count"},
			"redis_client_requests_total":                         &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.UnknownUnit, Desc: "Number of Redis client requests"},
			"redis_client_requests_duration_seconds_count":        &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The count of redis request latency, excluding blocking commands"},
			"redis_client_requests_duration_seconds_sum":          &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The sum of redis request latency, excluding blocking commands"},
		},
	}
}

func (*gitlabBaseMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{Name: "gitlab_base",
		Desc: "GitLab 编程语言层面指标",
		Tags: nil,
		Fields: map[string]interface{}{
			"ruby_sampler_duration_seconds_total": &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The time spent collecting stats"},
			"ruby_gc_duration_seconds_sum":        &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The sum of time spent by Ruby in GC"},
			"ruby_gc_duration_seconds_count":      &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The count of time spent by Ruby in GC"},
			"rails_queue_duration_seconds_sum":    &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The sum for latency between GitLab Workhorse forwarding a request to Rails"},
			"rails_queue_duration_seconds_count":  &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The counter for latency between GitLab Workhorse forwarding a request to Rails"},
		},
	}
}

func (*gitlabHTTPMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{Name: "gitlab_http",
		Desc: "GitLab HTTP 相关指标",
		Tags: map[string]interface{}{
			"method": inputs.NewTagInfo("方法"),
			"status": inputs.NewTagInfo("状态码"),
		},
		Fields: map[string]interface{}{
			"http_request_duration_seconds_count": &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The counter for request duration"},
			"http_request_duration_seconds_sum":   &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.DurationSecond, Desc: "The sum for request duration"},
			"http_health_requests_total":          &inputs.FieldInfo{DataType: inputs.Float, Unit: inputs.UnknownUnit, Desc: "Number of health requests"},
		},
	}
}
