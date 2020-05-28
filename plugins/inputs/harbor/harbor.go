package harbor

import (
	"context"
	"log"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/selfstat"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/models"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

type HarborMonitor struct {
	Harbor           []*HarborCfg
	runningInstances []*runningInstance
	ctx              context.Context
	cancelFun        context.CancelFunc
	accumulator      telegraf.Accumulator
	logger           *models.Logger
}

type runningInstance struct {
	cfg        *HarborCfg
	agent      *HarborMonitor
	logger     *models.Logger
	metricName string
}

func (_ *HarborMonitor) SampleConfig() string {
	return baiduIndexConfigSample
}

func (_ *HarborMonitor) Description() string {
	return ""
}

func (_ *HarborMonitor) Gather(telegraf.Accumulator) error {
	return nil
}

func (h *HarborMonitor) Start(acc telegraf.Accumulator) error {
	if len(b.Harbor) == 0 {
		log.Printf("W! [HarborMonitor] no configuration found")
		return nil
	}

	h.logger = &models.Logger{
		Errs: selfstat.Register("gather", "errors", nil),
		Name: `HarborMonitor`,
	}

	log.Printf("HarborMonitor cdn start")

	h.accumulator = acc

	for _, instCfg := range h.Harbor {
		r := &runningInstance{
			cfg:    instCfg,
			agent:  h,
			logger: h.logger,
		}

		r.metricName = instCfg.MetricName
		if r.metricName == "" {
			r.metricName = "baiduIndex"
		}

		if r.cfg.Interval.Duration == 0 {
			r.cfg.Interval.Duration = time.Minute * 10
		}

		h.runningInstances = append(h.runningInstances, r)

		go r.run(h.ctx)
	}

	return nil
}

func (h *HarborMonitor) Stop() {
	h.cancelFun()
}

func (r *runningInstance) run(ctx context.Context) error {
	defer func() {
		if e := recover(); e != nil {

		}
	}()

	for {
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}

		internal.SleepContext(ctx, r.cfg.Interval.Duration)
	}

	return nil
}

func init() {
	inputs.Add("harborMonitor", func() telegraf.Input {
		ac := &HarborMonitor{}
		ac.ctx, ac.cancelFun = context.WithCancel(context.Background())
		return ac
	})
}
