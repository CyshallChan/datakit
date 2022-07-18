// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package skywalking

import (
	"testing"

	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/trace"
)

func TestSkyWalkingAgent(t *testing.T) {
	afterGatherRun = itrace.AfterGatherFunc(func(inputName string, dktraces itrace.DatakitTraces, strikMod bool) {})
}

func TestSkyWalkingMetrics(t *testing.T) {}
