// Package container collect container metrics/loggings/objects.
package container

import (
	"fmt"
	"strconv"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/tailer"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var _ inputs.ReadEnv = (*Input)(nil)

const (
	objectInterval = time.Minute * 5
	metricInterval = time.Second * 20
)

type Input struct {
	DepercatedEndpoint string `toml:"endpoint"`
	DockerEndpoint     string `toml:"docker_endpoint"`
	ContainerdAddress  string `toml:"containerd_address"`

	LoggingRemoveAnsiEscapeCodes bool `toml:"logging_remove_ansi_escape_codes"`
	ExcludePauseContainer        bool `toml:"exclude_pause_container"`
	MaxLoggingLength             int  `toml:"max_logging_length"`

	ContainerIncludeMetric []string `toml:"container_include_metric"`
	ContainerExcludeMetric []string `toml:"container_exclude_metric"`
	ContainerIncludeLog    []string `toml:"container_include_log"`
	ContainerExcludeLog    []string `toml:"container_exclude_log"`

	K8sURL               string `toml:"kubernetes_url"`
	K8sBearerToken       string `toml:"bearer_token"`
	K8sBearerTokenString string `toml:"bearer_token_string"`
	DisableK8sEvents     bool   `toml:"disable_k8s_events"`

	Tags map[string]string `toml:"tags"`

	TLSCA              string `toml:"tls_ca"`
	TLSCert            string `toml:"tls_cert"`
	TLSKey             string `toml:"tls_key"`
	InsecureSkipVerify bool   `toml:"insecure_skip_verify"`
	DepercatedConf

	semStop *cliutils.Sem // start stop signal

	dockerInput     *dockerInput
	containerdInput *containerdInput
	k8sInput        *kubernetesInput

	chPause chan bool
	pause   bool
}

var (
	l = logger.DefaultSLogger(inputName)

	maxPauseCh = inputs.ElectionPauseChannelLength
)

func newInput() *Input {
	return &Input{
		DockerEndpoint:    dockerEndpoint,
		ContainerdAddress: containerdAddress,
		Tags:              make(map[string]string),
		chPause:           make(chan bool, maxPauseCh),
		semStop:           cliutils.NewSem(),
	}
}

func (*Input) SampleConfig() string { return sampleCfg }

func (*Input) Catalog() string { return catelog }

func (*Input) PipelineConfig() map[string]string { return nil }

func (*Input) GetPipeline() []*tailer.Option { return nil }

func (*Input) RunPipeline() { /*nil*/ }

func (*Input) AvailableArchs() []string { return []string{datakit.OSLinux} }

func (*Input) SampleMeasurement() []inputs.Measurement {
	res := []inputs.Measurement{}
	for _, mea := range measurements {
		res = append(res, mea)
	}
	return res
}

func (i *Input) Run() {
	l = logger.SLogger(inputName)

	l.Info("container input startd")
	io.FeedEventLog(&io.Reporter{Message: "container start ok, ready for collecting metrics.", Logtype: "event"})

	if i.setup() {
		return
	}

	objectTick := time.NewTicker(objectInterval)
	defer objectTick.Stop()
	metricTick := time.NewTicker(metricInterval)
	defer metricTick.Stop()

	if datakit.Docker && !i.DisableK8sEvents {
		go i.watchingK8sEventLog()
	}

	i.collectObject()

	for {
		select {
		case <-datakit.Exit.Wait():
			i.stop()
			l.Info("container exit success")
			return

		case <-i.semStop.Wait():
			l.Info("container exit return")
			return

		case <-metricTick.C:
			l.Debug("collect mertric")
			i.collectMetric()

		case <-objectTick.C:
			l.Debug("collect object")
			i.collectObject()

		case i.pause = <-i.chPause:
			globalPause.set(i.pause)
		}
	}
}

func (i *Input) stop() {
	if i.dockerInput != nil {
		i.dockerInput.stop()
	}
	if i.containerdInput != nil {
		i.containerdInput.stop()
	}
}

func (i *Input) collectObject() {
	l.Debug("collect object in func")
	if err := i.gatherDockerContainerObject(); err != nil {
		l.Errorf("failed to collect docker container object: %w", err)
	}

	if err := i.gatherContainerdObject(); err != nil {
		l.Errorf("failed to collect containerd object: %w", err)
	}

	if !datakit.Docker {
		return
	}
	if i.pause {
		l.Debug("not leader, skipped")
		return
	}

	if i.k8sInput == nil {
		l.Errorf("unrechable, k8s input is empty pointer")
		return
	}

	l.Debug("collect k8s resource")

	if err := i.gatherK8sResource(); err != nil {
		l.Errorf("failed fo collect k8s: %w", err)
	}
}

func (i *Input) collectMetric() {
	l.Debug("collect mertric in func")
	if err := i.gatherDockerContainerMetric(); err != nil {
		l.Errorf("failed to collect docker container metric: %w", err)
	}

	if err := i.gatherContainerdMetric(); err != nil {
		l.Errorf("failed to collect containerd metric: %w", err)
	}

	if err := i.watchNewDockerContainerLogs(); err != nil {
		l.Errorf("failed to watch container log: %w", err)
	}

	if !datakit.Docker {
		return
	}

	if i.pause {
		l.Debug("not leader, skipped")
		return
	}

	if i.k8sInput == nil {
		l.Errorf("unrechable, k8s input is empty pointer")
		return
	}

	l.Debug("collect k8s-pod metric")

	if err := i.gatherK8sPodMetrics(); err != nil {
		l.Errorf("failed to collect pod metric: %w", err)
	}
}

func (i *Input) gatherDockerContainerMetric() error {
	if i.dockerInput == nil {
		return nil
	}
	start := time.Now()

	res, err := i.dockerInput.gatherMetric()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		l.Debugf("container metric: no point")
		return nil
	}

	return inputs.FeedMeasurement("container-metric", datakit.Metric, res,
		&io.Option{CollectCost: time.Since(start)})
}

func (i *Input) gatherDockerContainerObject() error {
	if i.dockerInput == nil {
		return nil
	}
	start := time.Now()

	res, err := i.dockerInput.gatherObject()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		l.Debugf("container object: no point")
		return nil
	}

	return inputs.FeedMeasurement("container-object", datakit.Object, res,
		&io.Option{CollectCost: time.Since(start)})
}

func (i *Input) gatherContainerdMetric() error {
	if i.containerdInput == nil {
		return nil
	}
	start := time.Now()

	res, err := i.containerdInput.gatherMetric()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		l.Debugf("containerd metric: no point")
		return nil
	}

	return inputs.FeedMeasurement("containerd-metric", datakit.Metric, res,
		&io.Option{CollectCost: time.Since(start)})
}

func (i *Input) gatherContainerdObject() error {
	if i.containerdInput == nil {
		return nil
	}
	start := time.Now()

	res, err := i.containerdInput.gatherObject()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		l.Debugf("containerd object: no point")
		return nil
	}

	return inputs.FeedMeasurement("containerd-object", datakit.Object, res,
		&io.Option{CollectCost: time.Since(start)})
}

func (i *Input) gatherK8sResource() error {
	start := time.Now()

	metricMeas, objectMeas, err := i.k8sInput.gather()
	if err != nil {
		l.Errorf("k8s gather resource error: %w", err)
		return err
	}

	if err := inputs.FeedMeasurement("k8s-metric", datakit.Metric, metricMeas,
		&io.Option{CollectCost: time.Since(start)}); err != nil {
		l.Errorf("failed to feed k8s metrics: %w", err)
	}
	if err := inputs.FeedMeasurement("k8s-object", datakit.Object, objectMeas,
		&io.Option{CollectCost: time.Since(start)}); err != nil {
		l.Errorf("failed to feed k8s objects: %w", err)
	}

	return nil
}

func (i *Input) gatherK8sPodMetrics() error {
	start := time.Now()

	podMetrics, err := i.k8sInput.gatherPodMetrics()
	if err != nil {
		return err
	}

	return inputs.FeedMeasurement("k8s-pod", datakit.Metric, podMetrics,
		&io.Option{CollectCost: time.Since(start)})
}

func (i *Input) watchNewDockerContainerLogs() error {
	if i.dockerInput == nil {
		return nil
	}
	return i.dockerInput.watchNewContainerLogs()
}

func (i *Input) watchingK8sEventLog() {
	i.k8sInput.watchingEventLog(datakit.Exit.Wait())
}

func (i *Input) setup() bool {
	if i.DepercatedEndpoint != "" && i.DepercatedEndpoint != i.DockerEndpoint {
		i.DockerEndpoint = i.DepercatedEndpoint
	}

	for {
		select {
		case <-datakit.Exit.Wait():
			l.Info("exit")
			return true
		default:
			// nil
		}

		time.Sleep(time.Second)

		if d, err := newDockerInput(&dockerInputConfig{
			endpoint:               i.DockerEndpoint,
			excludePauseContainer:  i.ExcludePauseContainer,
			removeLoggingAnsiCodes: i.LoggingRemoveAnsiEscapeCodes,
			maxLoggingLength:       i.MaxLoggingLength,
			containerIncludeMetric: i.ContainerIncludeMetric,
			containerExcludeMetric: i.ContainerExcludeMetric,
			containerIncludeLog:    i.ContainerIncludeLog,
			containerExcludeLog:    i.ContainerExcludeLog,
			extraTags:              i.Tags,
		}); err != nil {
			l.Errorf("create docker input err: %w, skip", err)
		} else {
			i.dockerInput = d
		}

		if datakit.Docker {
			if k, err := newKubernetesInput(&kubernetesInputConfig{
				url:               i.K8sURL,
				bearerToken:       i.K8sBearerToken,
				bearerTokenString: i.K8sBearerTokenString,
				extraTags:         i.Tags,
			}); err != nil {
				l.Errorf("create k8s input err: %w", err)
				continue
			} else {
				i.k8sInput = k
				if i.dockerInput != nil {
					i.dockerInput.k8sClient = i.k8sInput.client
				}
			}
		}

		if c, err := newContainerdInput(&containerdInputConfig{
			endpoint:  i.ContainerdAddress,
			extraTags: i.Tags,
		}); err != nil {
			l.Warnf("create containerd input err: %w, skip", err)
		} else {
			i.containerdInput = c
		}

		break
	}

	return false
}

func (i *Input) Terminate() {
	if i.semStop != nil {
		i.semStop.Close()
	}
}

func (i *Input) Pause() error {
	tick := time.NewTicker(inputs.ElectionPauseTimeout)
	select {
	case i.chPause <- true:
		return nil
	case <-tick.C:
		return fmt.Errorf("pause %s failed", inputName)
	}
}

func (i *Input) Resume() error {
	tick := time.NewTicker(inputs.ElectionResumeTimeout)
	select {
	case i.chPause <- false:
		return nil
	case <-tick.C:
		return fmt.Errorf("resume %s failed", inputName)
	}
}

// ReadEnv , support envs：
//   ENV_INPUT_CONTAINER_DOCKER_ENDPOINT : string
//   ENV_INPUT_CONTAINER_CONTAINERD_ADDRESS : string
//   ENV_INPUT_CONTAINER_LOGGING_REMOVE_ANSI_ESCAPE_CODES : booler
//   ENV_INPUT_CONTAINER_TAGS : "a=b,c=d"
//   ENV_INPUT_CONTAINER_EXCLUDE_PAUSE_CONTAINER : booler
func (i *Input) ReadEnv(envs map[string]string) {
	if endpoint, ok := envs["ENV_INPUT_CONTAINER_DOCKER_ENDPOINT"]; ok {
		i.DockerEndpoint = endpoint
	}

	if address, ok := envs["ENV_INPUT_CONTAINER_CONTAINERD_ADDRESS"]; ok {
		i.ContainerdAddress = address
	}

	if remove, ok := envs["ENV_INPUT_CONTAINER_LOGGING_REMOVE_ANSI_ESCAPE_CODES"]; ok {
		b, err := strconv.ParseBool(remove)
		if err != nil {
			l.Warnf("parse ENV_INPUT_CONTAINER_LOGGING_REMOVE_ANSI_ESCAPE_CODES to bool: %s, ignore", err)
		} else {
			i.LoggingRemoveAnsiEscapeCodes = b
		}
	}

	if exclude, ok := envs["ENV_INPUT_CONTAINER_EXCLUDE_PAUSE_CONTAINER"]; ok {
		b, err := strconv.ParseBool(exclude)
		if err != nil {
			l.Warnf("parse ENV_INPUT_CONTAINER_EXCLUDE_PAUSE_CONTAINER to bool: %s, ignore", err)
		} else {
			i.ExcludePauseContainer = b
		}
	}

	if tagsStr, ok := envs["ENV_INPUT_CONTAINER_TAGS"]; ok {
		tags := config.ParseGlobalTags(tagsStr)
		for k, v := range tags {
			i.Tags[k] = v
		}
	}
}

//nolint:gochecknoinits
func init() {
	inputs.Add(inputName, func() inputs.Input {
		return newInput()
	})
}
