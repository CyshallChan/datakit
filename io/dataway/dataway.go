// Package dataway implement all dataway API request.
package dataway

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	ihttp "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/http"
)

var (
	apis = []string{
		datakit.CustomObject,
		datakit.DatakitPull,
		datakit.Election,
		datakit.ElectionHeartbeat,
		datakit.HeartBeat,
		datakit.KeyEvent,
		datakit.ListDataWay,
		datakit.LogFilter,
		datakit.LogUpload,
		datakit.Logging,
		datakit.Metric,
		datakit.MetricDeprecated,
		datakit.Network,
		datakit.Object,
		datakit.ObjectLabel,
		datakit.PipelinePull,
		datakit.QueryRaw,
		datakit.RUM,
		datakit.Security,
		datakit.Tracing,
		datakit.Workspace,
	}

	ExtraHeaders               = map[string]string{}
	AvailableDataways          = []string{}
	log                        = logger.DefaultSLogger("dataway")
	datawayListIntervalDefault = 60
	heartBeatIntervalDefault   = 30
)

type DataWayCfg struct {
	URLs      []string `toml:"urls"`
	endPoints []*endPoint

	DeprecatedURL    string `toml:"url,omitempty"`
	HTTPTimeout      string `toml:"timeout"`
	HTTPProxy        string `toml:"http_proxy"`
	Hostname         string `toml:"-"`
	DeprecatedHost   string `toml:"host,omitempty"`
	DeprecatedScheme string `toml:"scheme,omitempty"`
	DeprecatedToken  string `toml:"token,omitempty"`

	TimeoutDuration time.Duration `toml:"-"`
	httpCli         *http.Client

	MaxFails int `toml:"max_fail"`

	Proxy  bool `toml:"proxy,omitempty"`
	ontest bool
}

type endPoint struct {
	url         string
	host        string
	scheme      string
	proxy       string
	urlValues   url.Values
	categoryURL map[string]string
	ontest      bool
	fails       int
	dw          *DataWayCfg // reference
}

func (dw *DataWayCfg) String() string {
	arr := []string{fmt.Sprintf("dataways: [%s]", strings.Join(dw.URLs, ","))}

	for _, x := range dw.endPoints {
		arr = append(arr, "---------------------------------")
		for k, v := range x.categoryURL {
			arr = append(arr, fmt.Sprintf("% 24s: %s", k, v))
		}
	}

	return strings.Join(arr, "\n")
}

func (dw *DataWayCfg) ClientsCount() int {
	return len(dw.endPoints)
}

func (dw *DataWayCfg) GetToken() []string {
	resToken := []string{}
	for _, ep := range dw.endPoints {
		if ep.urlValues != nil {
			token := ep.urlValues.Get("token")
			if token != "" {
				resToken = append(resToken, token)
			}
		}
	}

	return resToken
}

func (dw *DataWayCfg) CheckToken(token string) (err error) {
	err = fmt.Errorf("token invalid format")

	tokenFormatMap := map[string]int{
		"token_": 32,
		"tkn_":   32,
		"tokn_":  24,
	}

	parts := strings.Split(token, "_")

	if len(parts) == 2 {
		prefix := parts[0] + "_"
		tokenVal := parts[1]

		if tokenLen, ok := tokenFormatMap[prefix]; ok {
			if len(tokenVal) == tokenLen {
				err = nil
			}
		}
	}

	return
}

func (dw *DataWayCfg) Apply() error {
	log = logger.SLogger("dataway")

	// 如果 env 已传入了 dataway 配置, 则不再追加老的 dataway 配置,
	// 避免俩边配置了同样的 dataway, 造成数据混乱
	if dw.DeprecatedURL != "" && len(dw.URLs) == 0 {
		dw.URLs = []string{dw.DeprecatedURL}
	}

	if len(dw.URLs) == 0 {
		return fmt.Errorf("dataway not set")
	}

	if dw.HTTPTimeout == "" {
		dw.HTTPTimeout = "5s"
	}

	if dw.MaxFails == 0 {
		dw.MaxFails = 20
	}

	timeout, err := time.ParseDuration(dw.HTTPTimeout)
	if err != nil {
		return err
	}

	dw.TimeoutDuration = timeout

	if err := dw.initHTTP(); err != nil {
		return err
	}

	dw.endPoints = dw.endPoints[:0]

	for _, httpurl := range dw.URLs {
		ep, err := dw.initEndpoint(httpurl)
		if err != nil {
			log.Errorf("init dataway url %s failed: %s", httpurl, err.Error())
			return err
		}

		dw.endPoints = append(dw.endPoints, ep)
	}

	return nil
}

func (dw *DataWayCfg) initEndpoint(httpurl string) (*endPoint, error) {
	u, err := url.ParseRequestURI(httpurl)
	if err != nil {
		log.Errorf("parse dataway url %s failed: %s", httpurl, err.Error())
		return nil, err
	}

	cli := &endPoint{
		url:         httpurl,
		scheme:      u.Scheme,
		urlValues:   u.Query(),
		host:        u.Host,
		categoryURL: map[string]string{},
		ontest:      dw.ontest,
		proxy:       dw.HTTPProxy,
		dw:          dw, // reference
	}

	for _, api := range apis {
		if cli.urlValues.Encode() != "" {
			cli.categoryURL[api] = fmt.Sprintf("%s://%s%s?%s",
				cli.scheme,
				cli.host,
				api,
				cli.urlValues.Encode())
		} else {
			cli.categoryURL[api] = fmt.Sprintf("%s://%s%s",
				cli.scheme,
				cli.host,
				api)
		}
	}

	return cli, nil
}

func (dw *DataWayCfg) initHTTP() error {
	cliopts := &ihttp.Options{
		DialTimeout: dw.TimeoutDuration,
	}

	if dw.HTTPProxy != "" { // set proxy
		if u, err := url.ParseRequestURI(dw.HTTPProxy); err != nil {
			log.Warnf("parse http proxy failed err: %s, ignored", err.Error())
		} else {
			cliopts.ProxyURL = u
			log.Infof("set dataway proxy to %s ok", dw.HTTPProxy)
		}
	}

	dw.httpCli = ihttp.Cli(cliopts)
	log.Debugf("httpCli: %p", dw.httpCli.Transport)

	return nil
}
