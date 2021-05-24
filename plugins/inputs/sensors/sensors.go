// +build linux

package sensors

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var (
	ErrSensorsTimeout = errors.New("Command 'sensors' process timeout")
)

var (
	defCommand   = "sensors"
	defPath      = "/usr/bin/sensors"
	defInterval  = datakit.Duration{Duration: 10 * time.Second}
	defTimeout   = datakit.Duration{Duration: 3 * time.Second}
	inputName    = "sensors"
	sampleConfig = `
[[inputs.sensors]]
	## Command path of 'senssor' usually under /usr/bin/sensors
	# path = "/usr/bin/senssors"

	## Gathering interval
	# interval = "10s"

	## Command timeout
	# timeout = "3s"

	## Customer tags, if set will be seen with every metric.
	[inputs.sensors.tags]
		# "key1" = "value1"
		# "key2" = "value2"
`
	l = logger.SLogger(inputName)
)

type Input struct {
	Path     string            `toml:"path"`
	Interval datakit.Duration  `toml:"interval"`
	Timeout  datakit.Duration  `toml:timeout`
	Tags     map[string]string `toml:"tags"`
}

func (*Input) Catalog() string {
	return inputName
}

func (*Input) SampleConfig() string {
	return sampleConfig
}

func (*Input) AvailableArchs() []string {
	return []string{datakit.OSLinux}
}

func (*Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{}
}

func (s *Input) Run() {
	l.Info("sensors input started")

	if finfo, err := os.Stat(s.Path); err != nil || finfo.IsDir() {
		l.Error(err.Error())
		if s.Path, err = exec.LookPath(defCommand); err != nil {
			l.Errorf("Can not find executable sensor command, install 'lm-sensors' first.")

			return
		}
		l.Infof("Fallback: use %q for gathering instead.", s.Path)
	}

	tick := time.NewTicker(s.Interval.Duration)
	for {
		select {
		case <-tick.C:
			if err := s.gather(); err != nil {
				l.Error(err.Error())
				io.FeedLastError(inputName, err.Error())
				continue
			}
		case <-datakit.Exit.Wait():
			l.Info("sensors input exits")

			return
		}
	}
}

func (s *Input) gather() error {
	var (
		cmd   = exec.Command(s.Path, "-u")
		buf   bytes.Buffer
		start = time.Now()
	)
	cmd.Stdout = &buf
	cmd.Stderr = nil
	err := cmd.Start()
	if err != nil {
		return err
	}

	timeout := time.AfterFunc(s.Timeout.Duration, func() {
		if err := cmd.Process.Kill(); err != nil {
			l.Error(err.Error())
		}
	})

	err = cmd.Wait()

	if !timeout.Stop() && err == nil {
		err = ErrSensorsTimeout
	}
	if err != nil {
		return err
	}

	if cache, err := parse(string(buf.Bytes())); err != nil {
		return err
	} else {
		return inputs.FeedMeasurement(inputName, datakit.Metric, cache, &io.Option{CollectCost: time.Now().Sub(start)})
	}
}

func (s *Input) getCustomerTags() map[string]string {
	if len(s.Tags) != 0 {
		return s.Tags
	} else {
		return make(map[string]string)
	}
}

func (s *Input) parse(output string) ([]inputs.Measurement, error) {
	var (
		lines  = strings.Split(strings.TrimSpace(output), "\n")
		tags   = s.getCustomerTags()
		fields = make(map[string]interface{})
		cache  []inputs.Measurement
	)
	for _, line := range lines {
		if line == "" {
			cache = append(cache, &sensorsMeasurement{
				name:   inputName,
				tags:   tags,
				fields: fields,
				ts:     time.Now(),
			})
			tags = s.getCustomerTags()
			fields = make(map[string]interface{})
		} else {
			if strings.Contains(line, ":") {
				parts := strings.Split(line, ":")
				if strings.HasSuffix(line, ":") {
					if len(fields) != 0 {
						cache = append(cache, &sensorsMeasurement{
							name:   inputName,
							tags:   tags,
							fields: fields,
							ts:     time.Now(),
						})

						tmp := make(map[string]string)
						for k, v := range tags {
							tmp[k] = v
						}
						tags = tmp
						fields = make(map[string]interface{})
					}
					tags["feature"] = strings.ToLower(strings.Replace(strings.TrimSpace(parts[0]), " ", "_", -1))
				} else if strings.HasPrefix(parts[0], " ") {
					if value, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64); err != nil {
						log.Println(err.Error())

						return nil, err
					} else {
						fields[strings.ToLower(strings.TrimSpace(parts[0]))] = value
					}
				} else {
					tags[strings.ToLower(parts[0])] = strings.TrimSpace(parts[1])
				}
			} else {
				tags["chip"] = line
			}
		}
	}
	cache = append(cache, &sensorsMeasurement{
		name:   inputName,
		tags:   tags,
		fields: fields,
		ts:     time.Now(),
	})

	return cache, nil
}

func init() {
	inputs.Add(inputName, func() inputs.Input {
		return &Input{
			Path:     defPath,
			Interval: defInterval,
			Timeout:  defTimeout,
		}
	})
}
