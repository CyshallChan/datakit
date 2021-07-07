package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/git"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/election"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/man"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

type enabledInput struct {
	Input     string `json:"input"`
	Instances int    `json:"instances"`
	Panics    int    `json:"panic"`
}

type DatakitStats struct {
	InputsStats     map[string]*io.InputsStat `json:"inputs_status"`
	EnabledInputs   []*enabledInput           `json:"enabled_inputs"`
	AvailableInputs []string                  `json:"available_inputs"`

	Version string `json:"version"`
	BuildAt string `json:"build_at"`
	Branch  string `json:"branch"`
	Uptime  string `json:"uptime"`
	OSArch  string `json:"os_arch"`

	Reload     time.Time `json:"reload"`
	ReloadCnt  int       `json:"reload_cnt"`
	ReloadInfo string    `json:"reload_info"`

	WithinDocker bool   `json:"docker"`
	IOChanStat   string `json:"io_chan_stats"`
	Elected      string `json:"elected"`
	AutoUpdate   bool   `json:"auto_update"`

	// markdown options
	DisableMonofont bool `json:"-"`

	CSS string `json:"-"`
}

var (
	part1 = `
- 版本       : {{.Version}}
- 运行时间   : {{.Uptime}}
- 发布日期   : {{.BuildAt}}
- 分支       : {{.Branch}}
- 系统类型   : {{.OSArch}}
- 容器运行   : {{.WithinDocker}}
- Reload 次数: {{.ReloadInfo}}
- IO 消耗统计: {{.IOChanStat}}
- 自动更新   ：{{.AutoUpdate}}
- 选举状态   ：{{.Elected}}
	`

	part2 = `
{{.InputsStatsTable}}
`

	part3 = `
{{.InputsConfTable}}
`

	fullMonitorTmpl = `
{{.CSS}}

# DataKit 运行展示
` + `
## 基本信息
` + part1 + `
## 采集器运行情况
` + part2 + `
## 采集器配置情况
` + part3
)

var (
	categoryMap = map[string]string{
		datakit.MetricDeprecated: "M",
		datakit.Metric:           "M",
		datakit.KeyEvent:         "E",
		datakit.Object:           "O",
		datakit.Logging:          "L",
		datakit.Tracing:          "T",
		datakit.Rum:              "R",
		datakit.Security:         "S",
	}
)

func (x *DatakitStats) InputsConfTable() string {
	const (
		tblHeader = `
| 采集器 | 实例个数 | 奔溃次数 |
| ----   | :----:   |  :----:  |
`
	)

	var rowFmt = "|`%s`|%d|%d|"
	if x.DisableMonofont {
		rowFmt = "|%s|%d|%d|"
	}

	if len(x.EnabledInputs) == 0 {
		return "没有开启任何采集器"
	}

	rows := []string{}
	for _, v := range x.EnabledInputs {
		rows = append(rows, fmt.Sprintf(rowFmt,
			v.Input,
			v.Instances,
			v.Panics,
		))
	}

	sort.Strings(rows)
	return tblHeader + strings.Join(rows, "\n")
}

func (x *DatakitStats) InputsStatsTable() string {

	const (
		tblHeader = `
| 采集器 | 数据类型 | 频率   | 平均 IO 大小 | 总次数 | 点数  | 首次采集 | 最近采集 | 平均采集消耗 | 最大采集消耗 | 当前错误(时间) |
| ----   | :----:   | :----: | :----:       | :----: | :---: | :----:   | :---:    | :----:       | :---:        | :----:         |
`
	)

	var rowFmt = "|`%s`|`%s`|%s|%d|%d|%d|%s|%s|%s|%s|`%s`(%s)|"
	if x.DisableMonofont {
		rowFmt = "|%s|%s|%s|%d|%d|%d|%s|%s|%s|%s|%s(%s)|"
	}

	if len(x.InputsStats) == 0 {
		return "暂无采集器统计数据"
	}

	now := time.Now()

	rows := []string{}

	for k, s := range x.InputsStats {

		firstIO := humanize.RelTime(s.First, now, "ago", "")
		lastIO := humanize.RelTime(s.Last, now, "ago", "")

		lastErr := "-"
		if s.LastErr != "" {
			lastErr = s.LastErr
		}

		lastErrTime := "-"
		if s.LastErr != "" {
			lastErrTime = humanize.RelTime(s.LastErrTS, now, "ago", "")
		}

		freq := "-"
		if s.Frequency != "" {
			freq = s.Frequency
		}

		category := "-"
		if s.Category != "" {
			category = categoryMap[s.Category]
		}

		rows = append(rows,
			fmt.Sprintf(rowFmt,
				k,
				category,
				freq,
				s.AvgSize,
				s.Count,
				s.Total,
				firstIO,
				lastIO,
				s.AvgCollectCost,
				s.MaxCollectCost,
				lastErr,
				lastErrTime,
			))
	}

	sort.Strings(rows)
	return tblHeader + strings.Join(rows, "\n")
}

func GetStats() (*DatakitStats, error) {

	now := time.Now()
	stats := &DatakitStats{
		Version:      datakit.Version,
		BuildAt:      git.BuildAt,
		Branch:       git.Branch,
		Uptime:       fmt.Sprintf("%v", now.Sub(uptime)),
		OSArch:       fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		ReloadCnt:    reloadCnt,
		ReloadInfo:   "0",
		WithinDocker: datakit.Docker,
		IOChanStat:   io.ChanStat(),
		Elected:      election.Elected(),
		AutoUpdate:   datakit.AutoUpdate,
	}

	if reloadCnt > 0 {
		stats.Reload = reload
		stats.ReloadInfo = fmt.Sprintf("%d(%s)", stats.ReloadCnt, humanize.RelTime(stats.Reload, now, "ago", ""))
	}

	var err error

	stats.InputsStats, err = io.GetStats(time.Second * 5) // get all inputs stats
	if err != nil {
		return nil, err
	}

	for k := range inputs.Inputs {
		if !datakit.Enabled(k) {
			continue
		}

		n := inputs.InputEnabled(k)
		npanic := inputs.GetPanicCnt(k)
		if n > 0 {
			stats.EnabledInputs = append(stats.EnabledInputs, &enabledInput{Input: k, Instances: n, Panics: npanic})
		}
	}

	for k := range inputs.Inputs {
		if !datakit.Enabled(k) {
			continue
		}
		stats.AvailableInputs = append(stats.AvailableInputs, fmt.Sprintf("[D] %s", k))
	}

	// add available inputs(datakit) stats
	stats.AvailableInputs = append(stats.AvailableInputs, fmt.Sprintf("tatal %d, datakit %d",
		len(stats.AvailableInputs), len(inputs.Inputs)))

	sort.Strings(stats.AvailableInputs)
	return stats, nil
}

func (ds *DatakitStats) Markdown(css string) ([]byte, error) {

	tmpl := fullMonitorTmpl

	temp, err := template.New("").Parse(tmpl)
	if err != nil {
		return nil, fmt.Errorf("parse markdown template failed: %s", err.Error())
	}

	if css != "" {
		ds.CSS = css
	}

	var buf bytes.Buffer
	if err := temp.Execute(&buf, ds); err != nil {
		return nil, fmt.Errorf("execute markdown template failed: %s", err.Error())
	}

	return buf.Bytes(), nil
}

func apiGetDatakitMonitor(c *gin.Context) {
	s, err := GetStats()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	mdbytes, err := s.Markdown(man.MarkdownCSS)
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	mdext := parser.CommonExtensions
	psr := parser.NewWithExtensions(mdext)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{Flags: htmlFlags}
	//opts := html.RendererOptions{Flags: htmlFlags, Head: headerScript}
	renderer := html.NewRenderer(opts)

	out := markdown.ToHTML(mdbytes, psr, renderer)

	c.Data(http.StatusOK, "text/html; charset=UTF-8", out)
}

func apiGetDatakitStats(c *gin.Context) {

	s, err := GetStats()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	body, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}

	c.Data(http.StatusOK, "application/json", body)
}
