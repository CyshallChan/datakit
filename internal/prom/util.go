// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package prom

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"strings"
	"time"

	dto "github.com/prometheus/client_model/go"
	iod "gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
)

func (p *Prom) getMetricTypeName(familyType dto.MetricType) string {
	var metricTypeName string
	switch familyType {
	case dto.MetricType_COUNTER:
		metricTypeName = "counter"
	case dto.MetricType_GAUGE:
		metricTypeName = "gauge"
	case dto.MetricType_HISTOGRAM:
		metricTypeName = "histogram"
	case dto.MetricType_SUMMARY:
		metricTypeName = "summary"
	case dto.MetricType_UNTYPED:
		metricTypeName = "untyped"
	}
	return metricTypeName
}

func (p *Prom) validMetricType(familyType dto.MetricType) bool {
	if len(p.opt.MetricTypes) == 0 {
		return true
	}
	typeName := p.getMetricTypeName(familyType)
	for _, mt := range p.opt.MetricTypes {
		if strings.ToLower(mt) == typeName {
			return true
		}
	}
	return false
}

func (p *Prom) validMetricName(name string) bool {
	if len(p.opt.MetricNameFilter) == 0 {
		return true
	}
	for _, p := range p.opt.MetricNameFilter {
		match, err := regexp.MatchString(p, name)
		if err != nil {
			continue
		}
		if match {
			return true
		}
	}
	return false
}

// getNames prioritizes naming rules as follows:
// 1. Check if any measurement rule is matched.
// 2. Check if measurement name is configured.
// 3. Check if measurement/field name can be split by the first '_' met.
// 4. If no term above matches, set both measurement name and field name to name.
func (p *Prom) getNames(name string) (measurementName string, fieldName string) {
	measurementName, fieldName = p.doGetNames(name)
	if measurementName == "" {
		measurementName = "prom"
	}
	return p.opt.MeasurementPrefix + measurementName, fieldName
}

func (p *Prom) doGetNames(name string) (measurementName string, fieldName string) {
	// Check if it matches custom rules.
	if mName, fName, matchAny := p.getNamesByRules(name); matchAny {
		return mName, fName
	}

	// Check if measurement name is set.
	if len(p.opt.MeasurementName) > 0 {
		return p.opt.MeasurementName, name
	}

	if mName, fName, matchAny := p.getNamesByDefault(name); matchAny {
		return mName, fName
	}

	return name, name
}

func (p *Prom) getNamesByRules(name string) (measurementName string, fieldName string, matchAny bool) {
	for _, rule := range p.opt.Measurements {
		if len(rule.Prefix) > 0 && strings.HasPrefix(name, rule.Prefix) {
			if rule.Name != "" {
				measurementName = rule.Name
			} else {
				// If rule name is not set, use rule prefix as measurement name but remove all trailing _.
				measurementName = strings.TrimRight(rule.Prefix, "_")
			}
			return measurementName, name[len(rule.Prefix):], true
		}
	}
	return
}

func (p *Prom) getNamesByDefault(name string) (measurementName string, fieldName string, matchAny bool) {
	// By default, measurement name and metric name are split according to the first '_' met.
	pattern := "(^[^_]+)_(.*)$"
	reg := regexp.MustCompile(pattern)
	if reg != nil {
		result := reg.FindAllStringSubmatch(name, -1)
		if len(result) == 1 {
			return result[0][1], result[0][2], true
		}
	}
	return
}

func (p *Prom) tagKVMatched(tags map[string]string) bool {
	if p.opt.IgnoreTagKV == nil {
		return false
	}

	for k, v := range tags {
		if res, ok := p.opt.IgnoreTagKV[k]; ok {
			for _, re := range res {
				if re.MatchString(v) {
					return true
				}
			}
		}
	}

	return false
}

func (p *Prom) getTags(labels []*dto.LabelPair) map[string]string {
	tags := map[string]string{}

	// Add custom tags.
	for k, v := range p.opt.Tags {
		tags[k] = v
	}

	// Add prometheus labels as tags.
	for _, lab := range labels {
		tags[lab.GetName()] = lab.GetValue()
	}

	p.removeIgnoredTags(tags)
	p.renameTags(tags)

	return tags
}

func (p *Prom) removeIgnoredTags(tags map[string]string) {
	for t := range tags {
		for _, ignoredTag := range p.opt.TagsIgnore {
			if t == ignoredTag {
				delete(tags, t)
			}
		}
	}
}

func (p *Prom) renameTags(tags map[string]string) {
	if tags == nil || p.opt.RenameTags == nil {
		return
	}

	for oldKey, newKey := range p.opt.RenameTags.Mapping {
		if v, ok := tags[oldKey]; ok { // rename the tag
			if _, exists := tags[newKey]; exists && !p.opt.RenameTags.OverwriteExistTags {
				continue
			}

			delete(tags, oldKey)
			tags[newKey] = v
		}
	}
}

func (p *Prom) filterMetricFamilies(metricFamilies map[string]*dto.MetricFamily) map[string]*dto.MetricFamily {
	filteredMetricFamilies := make(map[string]*dto.MetricFamily)
	for name, value := range metricFamilies {
		if p.validMetricName(name) && p.validMetricType(value.GetType()) {
			filteredMetricFamilies[name] = value
		}
	}
	return filteredMetricFamilies
}

// Text2Metrics converts raw prometheus metric text to line protocol point.
func (p *Prom) Text2Metrics(in io.Reader) (pts []*iod.Point, lastErr error) {
	metricFamilies, err := p.parser.TextToMetricFamilies(in)
	if err != nil {
		return nil, err
	}

	startTime := time.Now()

	filteredMetricFamilies := p.filterMetricFamilies(metricFamilies)

	for name, value := range filteredMetricFamilies {
		measurementName, fieldName := p.getNames(name)

		switch value.GetType() {
		case dto.MetricType_GAUGE, dto.MetricType_UNTYPED, dto.MetricType_COUNTER:
			for _, m := range value.GetMetric() {
				v := getValue(m, value.GetType())
				if math.IsInf(v, 0) || math.IsNaN(v) {
					continue
				}

				fields := map[string]interface{}{
					fieldName: v,
				}
				tags := p.getTags(m.GetLabel())

				if !p.tagKVMatched(tags) {
					pt, err := iod.MakePoint(measurementName, tags, fields, getTimestampS(m, startTime))
					if err != nil {
						lastErr = err
					} else {
						pts = append(pts, pt)
					}
				}
			}

		case dto.MetricType_SUMMARY:
			for _, m := range value.GetMetric() {
				fields := map[string]interface{}{
					fieldName + "_count": float64(m.GetSummary().GetSampleCount()),
					fieldName + "_sum":   m.GetSummary().GetSampleSum(),
				}

				tags := p.getTags(m.GetLabel())

				if !p.tagKVMatched(tags) {
					pt, err := iod.MakePoint(measurementName, tags, fields, getTimestampS(m, startTime))
					if err != nil {
						lastErr = err
					} else {
						pts = append(pts, pt)
					}
				}

				for _, q := range m.GetSummary().Quantile {
					fields := map[string]interface{}{
						fieldName: q.GetValue(),
					}

					tags := p.getTags(m.GetLabel())
					tags["quantile"] = fmt.Sprint(q.GetQuantile())

					if !p.tagKVMatched(tags) {
						pt, err := iod.MakePoint(measurementName, tags, fields, getTimestampS(m, startTime))
						if err != nil {
							lastErr = err
						} else {
							pts = append(pts, pt)
						}
					}
				}
			}

		case dto.MetricType_HISTOGRAM:
			for _, m := range value.GetMetric() {
				fields := map[string]interface{}{
					fieldName + "_count": float64(m.GetHistogram().GetSampleCount()),
					fieldName + "_sum":   m.GetHistogram().GetSampleSum(),
				}

				tags := p.getTags(m.GetLabel())

				if !p.tagKVMatched(tags) {
					pt, err := iod.MakePoint(measurementName, tags, fields, getTimestampS(m, startTime))
					if err != nil {
						lastErr = err
					} else {
						pts = append(pts, pt)
					}
				}

				for _, b := range m.GetHistogram().GetBucket() {
					fields := map[string]interface{}{
						fieldName + "_bucket": b.GetCumulativeCount(),
					}

					tags := p.getTags(m.GetLabel())
					tags["le"] = fmt.Sprint(b.GetUpperBound())

					if !p.tagKVMatched(tags) {
						pt, err := iod.MakePoint(measurementName, tags, fields, getTimestampS(m, startTime))
						if err != nil {
							lastErr = err
						} else {
							pts = append(pts, pt)
						}
					}
				}
			}
		}
	}

	return pts, lastErr
}

func getValue(m *dto.Metric, metricType dto.MetricType) float64 {
	switch metricType { //nolint:exhaustive
	case dto.MetricType_GAUGE:
		return m.GetGauge().GetValue()
	case dto.MetricType_UNTYPED:
		return m.GetUntyped().GetValue()
	case dto.MetricType_COUNTER:
		return m.GetCounter().GetValue()
	default:
		// Shouldn't get here.
		return 0
	}
}

func getTimestampS(m *dto.Metric, startTime time.Time) time.Time {
	if m.GetTimestampMs() != 0 {
		return time.Unix(m.GetTimestampMs()/1000, 0)
	}
	return startTime
}
