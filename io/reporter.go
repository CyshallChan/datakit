// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package io

import (
	"context"
	"regexp"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
)

type Reporter struct {
	Status   string `json:"status"` // info warning error
	Message  string `json:"message"`
	Category string `josn:"category"`
	Logtype  string `josn:"Logtype"`
	feed     func(string, string, []*Point, *Option) error
}

func (e *Reporter) Tags() map[string]string {
	tags := map[string]string{
		"source":   "datakit",
		"status":   "info",
		"category": "default",
		"log_type": "",
	}

	if len(e.Status) > 0 {
		tags["status"] = e.Status
	}

	if len(e.Category) > 0 {
		tags["category"] = e.Category
	}

	if len(e.Logtype) > 0 {
		tags["log_type"] = e.Logtype
	}

	return tags
}

func (e *Reporter) Fields() map[string]interface{} {
	fields := map[string]interface{}{
		"message": "",
	}

	if len(e.Message) > 0 {
		fields["message"] = e.escape(e.Message)
	}

	return fields
}

func (e *Reporter) escape(message string) string {
	p := regexp.MustCompile(`token=(\w+)`)

	escapedMessage := p.ReplaceAllString(message, "token=xxxxxx")

	return escapedMessage
}

// addReporter report log, should not block.
func addReporter(reporter Reporter) {
	tags := reporter.Tags()
	fields := reporter.Fields()

	now := time.Now()
	pt, err := MakePoint("datakit", tags, fields, now)

	if reporter.feed == nil {
		reporter.feed = Feed
	}

	if err == nil {
		g := datakit.G("io")
		g.Go(func(ctx context.Context) error {
			if err := reporter.feed("datakit", datakit.Logging, []*Point{pt}, nil); err != nil {
				log.Debugf("feed logging error: %s", err.Error())
			}
			return nil
		})
	} else {
		log.Debugf("make point error: %s", err.Error())
	}
}
