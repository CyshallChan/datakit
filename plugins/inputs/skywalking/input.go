// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package skywalking handle SkyWalking tracing metrics.
package skywalking

import (
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var (
	inputName    = "skywalking"
	sampleConfig = `
[[inputs.skywalking]]
  ## skywalking grpc server listening on address
  address = "localhost:13800"
	## customer tags
  # [inputs.skywalking.V3.tags]
    # tag1 = "tag1"
    # tag2 = "tag2"
    # ...
`
	log                = logger.DefaultSLogger(inputName)
	_   inputs.InputV2 = &Input{}
)

var (
	defAddr = "localhost:13800"
	tags    map[string]string
)

// deprecated.
type skywalkingConfig struct {
	Address string            `toml:"address"`
	Tags    map[string]string `toml:"tags"`
}

type Input struct {
	V2      *skywalkingConfig `toml:"V2"` // deprecated
	V3      *skywalkingConfig `toml:"V3"` // deprecated
	Address string            `toml:"address"`
	Tags    map[string]string `toml:"tags"`
}

func (*Input) Catalog() string {
	return inputName
}

func (*Input) SampleConfig() string {
	return sampleConfig
}

func (*Input) AvailableArchs() []string {
	return datakit.AllArch
}

func (i *Input) SampleMeasurement() []inputs.Measurement {
	return nil
} // TODO

func (i *Input) Run() {
	log = logger.SLogger(inputName)
	log.Infof("%s input started...", inputName)

	if len(i.Address) == 0 {
		i.Address = defAddr
	}
	if len(i.Tags) != 0 {
		tags = i.Tags
	}

	log.Debug("start skywalking grpc v3 server")
	go runServerV3(i.Address)
}

func init() { //nolint:gochecknoinits
	inputs.Add(inputName, func() inputs.Input {
		return &Input{}
	})
}
