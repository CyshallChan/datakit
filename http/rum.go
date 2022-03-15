// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package http

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-sourcemap/sourcemap"
	influxm "github.com/influxdata/influxdb1-client/models"
	lp "gitlab.jiagouyun.com/cloudcare-tools/cliutils/lineproto"
	uhttp "gitlab.jiagouyun.com/cloudcare-tools/cliutils/network/http"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/funcs"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/ip2isp"
)

var (
	sourcemapCache = make(map[string]map[string]*sourcemap.Consumer)
	sourcemapLock  sync.Mutex

	rumMetricNames = map[string]bool{
		`view`:      true,
		`resource`:  true,
		`error`:     true,
		`long_task`: true,
		`action`:    true,
	}

	rumMetricAppID = "app_id"
)

func geoTags(srcip string) map[string]string {
	if srcip == "" {
		return nil
	}

	ipInfo, err := funcs.Geo(srcip)

	l.Debugf("ipinfo(%s): %+#v", srcip, ipInfo)

	if err != nil {
		l.Warnf("geo failed: %s, ignored", err)
		return nil
	}

	switch ipInfo.Country { // #issue 354
	case "TW":
		ipInfo.Country = "CN"
		ipInfo.Region = "Taiwan"
	case "MO":
		ipInfo.Country = "CN"
		ipInfo.Region = "Macao"
	case "HK":
		ipInfo.Country = "CN"
		ipInfo.Region = "Hong Kong"
	}

	// 无脑填充 geo 数据
	tags := map[string]string{
		"city":     ipInfo.City,
		"province": ipInfo.Region,
		"country":  ipInfo.Country,
		"isp":      ip2isp.SearchIsp(srcip),
		"ip":       srcip,
	}

	return tags
}

func doHandleRUMBody(body []byte,
	precision string,
	isjson bool,
	extraTags map[string]string,
	appIDWhiteList []string) ([]*io.Point, error) {
	if isjson {
		opt := lp.NewDefaultOption()
		opt.Precision = precision
		opt.ExtraTags = extraTags
		rumpts, err := jsonPoints(body, opt)
		if err != nil {
			return nil, err
		}
		for _, p := range rumpts {
			tags := p.Tags()
			if tags != nil {
				if !contains(tags[rumMetricAppID], appIDWhiteList) {
					return nil, ErrRUMAppIDNotInWhiteList
				}
			}
		}
		return rumpts, nil
	}

	rumpts, err := lp.ParsePoints(body, &lp.Option{
		Time:      time.Now(),
		Precision: precision,
		ExtraTags: extraTags,
		Strict:    true,

		// 由于 RUM 数据需要分别处理，故用回调函数来区分
		Callback: func(p influxm.Point) (influxm.Point, error) {
			name := string(p.Name())

			if !contains(p.Tags().GetString(rumMetricAppID), appIDWhiteList) {
				return nil, ErrRUMAppIDNotInWhiteList
			}

			if _, ok := rumMetricNames[name]; !ok {
				return nil, uhttp.Errorf(ErrUnknownRUMMeasurement, "unknow RUM measurement: %s", name)
			}

			// handle sourcemap
			if name == "error" {
				sdkName := p.Tags().GetString("sdk_name")
				if sdkName == "df_web_rum_sdk" { // only support web now
					err := handleSourcemap(p)
					if err != nil {
						l.Debugf("handle source map failed: %s", err.Error())
					}
				}
			}

			return p, nil
		},
	})
	if err != nil {
		l.Warnf("doHandleRUMBody: %s", err)
		return nil, err
	}

	return io.WrapPoint(rumpts), nil
}

func contains(str string, list []string) bool {
	if len(list) == 0 {
		return true
	}
	for _, a := range list {
		if a == str {
			return true
		}
	}
	return false
}

func getSrcIP(ac *APIConfig, req *http.Request) (ip string) {
	if ac != nil {
		ip = req.Header.Get(ac.RUMOriginIPHeader)
		l.Debugf("get ip from %s: %s", ac.RUMOriginIPHeader, ip)
		if ip == "" {
			for k, v := range req.Header {
				l.Debugf("%s: %s", k, strings.Join(v, ","))
			}
		}
	} else {
		l.Info("apiConfig not set")
	}

	if ip != "" {
		l.Debugf("header remote addr: %s", ip)
		parts := strings.Split(ip, ",")
		if len(parts) > 0 {
			ip = parts[0] // 注意：此处只取第一个 IP 作为源 IP
			return
		}
	} else { // 默认取 http 框架带进来的 IP
		l.Debugf("gin remote addr: %s", req.RemoteAddr)
		host, _, err := net.SplitHostPort(req.RemoteAddr)
		if err == nil {
			ip = host
			return
		} else {
			l.Warnf("net.SplitHostPort(%s): %s, ignored", req.RemoteAddr, err)
		}
	}

	return ip
}

func handleRUMBody(body []byte,
	precision string,
	isjson bool,
	geoInfo map[string]string,
	list []string) ([]*io.Point, error) {
	return doHandleRUMBody(body, precision, isjson, geoInfo, list)
}

func handleSourcemap(p influxm.Point) error {
	fields, err := p.Fields()
	if err != nil {
		return fmt.Errorf("parse field error: %w", err)
	}
	errStack, ok := fields["error_stack"]

	// if error_stack exists
	if ok {
		errStackStr := fmt.Sprintf("%v", errStack)

		appID := p.Tags().GetString("app_id")
		env := p.Tags().GetString("env")
		version := p.Tags().GetString("version")

		if len(appID) > 0 && (len(env) > 0) && (len(version) > 0) {
			zipFile := GetSourcemapZipFileName(appID, env, version)
			if sourcemapItem, ok := sourcemapCache[zipFile]; ok {
				errorStackSource := getSourcemap(errStackStr, sourcemapItem)
				errorStackSourceBase64 := base64.StdEncoding.EncodeToString([]byte(errorStackSource)) // tag cannot have '\n'
				p.AddTag("error_stack_source_base64", errorStackSourceBase64)
			}
		}
	}

	return nil
}

func getSourcemap(errStack string, sourcemapItem map[string]*sourcemap.Consumer) string {
	reg := regexp.MustCompile(`@ .*:\d+:\d+`)

	replaceStr := reg.ReplaceAllStringFunc(errStack, func(str string) string {
		return str[0:2] + getSourceMapString(str[2:], sourcemapItem)
	})
	return replaceStr
}

func getSourceMapString(str string, sourcemapItem map[string]*sourcemap.Consumer) string {
	parts := strings.Split(str, ":")
	partsLen := len(parts)
	if partsLen < 3 {
		return str
	}
	rowNumber, err := strconv.Atoi(parts[partsLen-2])
	if err != nil {
		return str
	}
	colNumber, err := strconv.Atoi(parts[partsLen-1])
	if err != nil {
		return str
	}

	path := strings.Join(parts[:partsLen-2], ":") // http://localhost:5500/dist/bundle.js

	urlObj, err := url.Parse(path)
	if err != nil {
		l.Debugf("parse url failed, %s, %s", path, err.Error())
		return str
	}

	urlPath := strings.TrimPrefix(urlObj.Path, "/")
	sourceMapFileName := urlPath + ".map"

	smap, ok := sourcemapItem[sourceMapFileName]
	if !ok {
		l.Debugf("no sourcemap: %s", sourceMapFileName)
		return str
	}

	file, _, line, col, ok := smap.Source(rowNumber, colNumber)

	if ok {
		return fmt.Sprintf("%s:%v:%v", file, line, col)
	}

	return str
}

// GetSourcemapZipFileName  zip file name.
func GetSourcemapZipFileName(applicatinID, env, version string) string {
	fileName := fmt.Sprintf("%s-%s-%s.zip", applicatinID, env, version)

	return strings.ReplaceAll(fileName, string(filepath.Separator), "__")
}

func GetRumSourcemapDir() string {
	rumDir := filepath.Join(datakit.DataDir, "rum")
	return rumDir
}

func loadSourcemapFile() {
	rumDir := GetRumSourcemapDir()
	files, err := ioutil.ReadDir(rumDir)
	if err != nil {
		l.Errorf("load rum sourcemap dir failed: %s", err.Error())
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.HasSuffix(fileName, ".zip") {
				sourcemapItem, err := loadZipFile(filepath.Join(rumDir, fileName))
				if err != nil {
					l.Debugf("load zip file %s failed, %s", fileName, err.Error())
					continue
				}

				sourcemapCache[fileName] = sourcemapItem
			}
		}
	}
}

func loadZipFile(zipFile string) (map[string]*sourcemap.Consumer, error) {
	sourcemapItem := make(map[string]*sourcemap.Consumer)

	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close() //nolint:errcheck

	for _, f := range zipReader.File {
		if !f.FileInfo().IsDir() && strings.HasSuffix(f.Name, ".map") {
			file, err := f.Open()
			if err != nil {
				l.Debugf("ignore sourcemap %s, %s", f.Name, err.Error())
				continue
			}
			defer file.Close() // nolint:errcheck

			content, err := ioutil.ReadAll(file)
			if err != nil {
				l.Debugf("ignore sourcemap %s, %s", f.Name, err.Error())
				continue
			}

			smap, err := sourcemap.Parse(f.Name, content)
			if err != nil {
				l.Debugf("sourcemap parse failed, %s", err.Error())
				continue
			}

			sourcemapItem[f.Name] = smap
		}
	}

	return sourcemapItem, nil
}

func updateSourcemapCache(zipFile string) {
	fileName := filepath.Base(zipFile)
	if strings.HasSuffix(fileName, ".zip") {
		if sourcemapItem, err := loadZipFile(zipFile); err != nil {
			l.Debugf("load zip file error: %s", err.Error())
		} else {
			sourcemapLock.Lock()
			sourcemapCache[fileName] = sourcemapItem
			sourcemapLock.Unlock()
			l.Debugf("load sourcemap: %s", fileName)
		}
	}
}

func deleteSourcemapCache(zipFile string) {
	fileName := filepath.Base(zipFile)
	if strings.HasSuffix(fileName, ".zip") {
		sourcemapLock.Lock()
		delete(sourcemapCache, fileName)
		sourcemapLock.Unlock()
	}
}
