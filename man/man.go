// Package man manages all datakit documents
package man

import (
	"bytes"
	"fmt"

	// nolint:typecheck
	"strings"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/git"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/funcs"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var (
	ManualBox = packr.New("manuals", "./manuals")
	OtherDocs = map[string]interface{}{
		// value not used, just document the markdown relative path
		// all manuals under man/manuals/
		"apis":                     "man/manuals/apis.md",
		"changelog":                "man/manuals/changelog.md",
		"datakit-arch":             "man/manuals/datakit-arch.md",
		"datakit-batch-deploy":     "man/manuals/datakit-batch-deploy.md",
		"datakit-conf-how-to":      "man/manuals/datakit-conf-how-to.md",
		"datakit-daemonset-deploy": "man/manuals/datakit-daemonset-deploy.md",
		"datakit-dql-how-to":       "man/manuals/datakit-dql-how-to.md",
		"datakit-how-to":           "man/manuals/datakit-how-to.md", // deprecated
		"datakit-install":          "man/manuals/datakit-install.md",
		"datakit-monitor":          "man/manuals/datakit-monitor.md",
		"datakit-offline-install":  "man/manuals/datakit-offline-install.md",
		"datakit-on-public":        "man/manuals/datakit-on-public.md",
		"datakit-pl-how-to":        "man/manuals/datakit-pl-how-to.md",
		"datakit-service-how-to":   "man/manuals/datakit-service-how-to.md",
		"datakit-tools-how-to":     "man/manuals/datakit-tools-how-to.md",
		"datakit-update":           "man/manuals/datakit-update.md",
		"datatypes":                "man/manuals/datatypes.md",
		"dataway":                  "man/manuals/dataway.md",
		"dca":                      "man/manuals/dca.md",
		"ddtrace-java":             "man/manuals/ddtrace-java.md",
		"ddtrace-python":           "man/manuals/ddtrace-python.md",
		"development":              "man/manuals/development.md",
		"dialtesting_json":         "man/manuals/dialtesting_json.md",
		"election":                 "man/manuals/election.md",
		"k8s-config-how-to":        "man/manuals/k8s-config-how-to.md",
		"kubernetes-prom":          "man/manuals/kubernetes-prom.md",
		"kubernetes-x":             "man/manuals/kubernetes-x.md",
		"logfwd":                   "man/manuals/logfwd.md",
		"logging_socket":           "man/manuals/logging_socket.md",
		"logging-pipeline-bench":   "man/manuals/logging-pipeline-bench.md",
		"pipeline":                 "man/manuals/pipeline.md",
		"prometheus":               "man/manuals/prometheus.md",
		"rum":                      "man/manuals/rum.md",
		"sec-checker":              "man/manuals/sec-checker.md",
		"telegraf":                 "man/manuals/telegraf.md",
		"why-no-data":              "man/manuals/why-no-data.md",
	}
	l = logger.DefaultSLogger("man")
)

type Params struct {
	InputName      string
	Catalog        string
	InputSample    string
	Version        string
	ReleaseDate    string
	Measurements   []*inputs.MeasurementInfo
	CSS            string
	AvailableArchs string
	PipelineFuncs  string
}

func Get(name string) (string, error) {
	return ManualBox.FindString(name + ".md")
}

type Option struct {
	WithCSS                       bool
	IgnoreMissing                 bool
	DisableMonofontOnTagFieldName bool
	ManVersion                    string
}

func BuildMarkdownManual(name string, opt *Option) ([]byte, error) {
	var p *Params

	css := MarkdownCSS
	ver := datakit.Version

	if !opt.WithCSS {
		css = ""
	}

	if opt.ManVersion != "" {
		ver = opt.ManVersion
	}

	if opt.DisableMonofontOnTagFieldName {
		inputs.MonofontOnTagFieldName = false
	}

	if _, ok := OtherDocs[name]; ok {
		p = &Params{
			Version:     ver,
			ReleaseDate: git.BuildAt,
			CSS:         css,
		}
		// Add pipeline functions doc.
		if name == "pipeline" {
			sb := strings.Builder{}
			for _, v := range funcs.PipelineFunctionDocs {
				sb.WriteString(v.Doc)
			}
			p.PipelineFuncs = sb.String()
		}
	} else {
		c, ok := inputs.Inputs[name]
		if !ok {
			return nil, fmt.Errorf("input %s not found", name)
		}

		input := c()
		switch i := input.(type) {
		case inputs.InputV2:
			sampleMeasurements := i.SampleMeasurement()
			p = &Params{
				InputName:      name,
				InputSample:    i.SampleConfig(),
				Catalog:        i.Catalog(),
				Version:        ver,
				ReleaseDate:    git.BuildAt,
				CSS:            css,
				AvailableArchs: strings.Join(i.AvailableArchs(), ","),
			}
			for _, m := range sampleMeasurements {
				p.Measurements = append(p.Measurements, m.Info())
			}

		default:
			l.Warnf("incomplete input: %s", name)

			return nil, nil
		}
	}

	md, err := Get(name)
	if err != nil {
		if !opt.IgnoreMissing {
			return nil, err
		} else {
			l.Warn(err)
			return nil, nil
		}
	}
	temp, err := template.New(name).Parse(md)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := temp.Execute(&buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
