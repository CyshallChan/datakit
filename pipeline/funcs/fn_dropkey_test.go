// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"testing"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
)

func TestDropKey(t *testing.T) {
	cases := []struct {
		name, pl, in string
		key          string
		fail         bool
	}{
		{
			name: "value type: string",
			in:   `162.62.81.1 - - [29/Nov/2021:07:30:50 +0000] "POST /?signature=b8d8ea&timestamp=1638171049 HTTP/1.1" 200 413 "-" "Mozilla/4.0"`,
			pl: `
grok(_, "%{IPORHOST:client_ip} %{NOTSPACE} %{NOTSPACE} \\[%{HTTPDATE:time}\\] \"%{DATA} %{GREEDYDATA} HTTP/%{NUMBER}\" %{INT:status_code} %{INT:bytes}")
drop_key(client_ip)
`,
			key:  "client_ip",
			fail: false,
		},
	}

	for idx, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runner, err := NewTestingRunner(tc.pl)
			if err != nil {
				if tc.fail {
					t.Logf("[%d]expect error: %s", idx, err)
				} else {
					t.Errorf("[%d] failed: %s", idx, err)
				}
				return
			}
			pt, _ := io.MakePoint("test", map[string]string{},
				map[string]interface{}{
					"message": tc.in,
				}, time.Now())
			ret, err := runner.Run(pt)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(ret)
			if v, ok := ret.Fields[tc.key]; ok {
				t.Errorf("[%d] failed: key `%s` value `%v`", idx, tc.key, v)
			} else {
				t.Logf("[%d] PASS", idx)
			}
		})
	}
}
