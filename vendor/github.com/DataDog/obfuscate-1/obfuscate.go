// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package obfuscate implements quantizing and obfuscating of tags and resources for
// a set of spans matching a certain criteria.
package obfuscate

import (
	"bytes"
	"sync/atomic"

	"github.com/DataDog/datadog-go/statsd"
)

// Obfuscator quantizes and obfuscates spans. The obfuscator is not safe for
// concurrent use.
type Obfuscator struct {
	opts                 *Config
	es                   *jsonObfuscator // nil if disabled
	mongo                *jsonObfuscator // nil if disabled
	sqlExecPlan          *jsonObfuscator // nil if disabled
	sqlExecPlanNormalize *jsonObfuscator // nil if disabled
	// sqlLiteralEscapes reports whether we should treat escape characters literally or as escape characters.
	// A non-zero value means 'yes'. Different SQL engines behave in different ways and the tokenizer needs
	// to be generic.
	// Not safe for concurrent use.
	sqlLiteralEscapes int32
	// queryCache keeps a cache of already obfuscated queries.
	queryCache *measuredCache
	log        Logger
}

// SetSQLLiteralEscapes sets whether or not escape characters should be treated literally by the SQL obfuscator.
func (o *Obfuscator) SetSQLLiteralEscapes(ok bool) {
	if ok {
		atomic.StoreInt32(&o.sqlLiteralEscapes, 1)
	} else {
		atomic.StoreInt32(&o.sqlLiteralEscapes, 0)
	}
}

// SQLLiteralEscapes reports whether escape characters should be treated literally by the SQL obfuscator.
func (o *Obfuscator) SQLLiteralEscapes() bool {
	return atomic.LoadInt32(&o.sqlLiteralEscapes) == 1
}

// NewObfuscator creates a new obfuscator
func NewObfuscator(cfg *Config) *Obfuscator {
	if cfg == nil {
		cfg = &Config{
			Statsd: &statsd.NoOpClient{},
			Log:    noOpLogger{},
		}
	}
	cache := new(measuredCache) // no-op as is
	if cfg.SQL.Cache {
		cache = newMeasuredCache(cfg.Statsd)
	}
	o := Obfuscator{
		opts:       cfg,
		queryCache: cache,
	}
	if cfg.ES.Enabled {
		o.es = newJSONObfuscator(&cfg.ES, &o)
	}
	if cfg.Mongo.Enabled {
		o.mongo = newJSONObfuscator(&cfg.Mongo, &o)
	}
	if cfg.SQLExecPlan.Enabled {
		o.sqlExecPlan = newJSONObfuscator(&cfg.SQLExecPlan, &o)
	}
	if cfg.SQLExecPlanNormalize.Enabled {
		o.sqlExecPlanNormalize = newJSONObfuscator(&cfg.SQLExecPlanNormalize, &o)
	}
	return &o
}

// Stop cleans up after a finished Obfuscator.
func (o *Obfuscator) Stop() { o.queryCache.Close() }

// Obfuscate may obfuscate span's properties based on its type and on the Obfuscator's
// configuration.
func (o *Obfuscator) Obfuscate(typ, q string) (*ObfuscatedQuery, error) {
	switch typ {
	case "sql", "cassandra":
		out, err := o.ObfuscateSQLString(q)
		if err != nil {
			return &ObfuscatedQuery{Query: nonParsableResource}, err
		}
		return out, nil
	default:
		return &ObfuscatedQuery{Query: o.obfuscateNonDB(typ, q)}, nil
	}
}

func (o *Obfuscator) obfuscateNonDB(typ, q string) string {
	switch typ {
	case "redis":
		if o.opts.Redis.Enabled {
			return o.obfuscateRedis(q)
		}
	case "memcached":
		if o.opts.Memcached.Enabled {
			return o.obfuscateMemcached(q)
		}
	case "web", "http":
		return o.obfuscateHTTP(q)
	case "mongodb":
		return o.obfuscateJSON(q, o.mongo)
	case "elasticsearch":
		return o.obfuscateJSON(q, o.es)
	}
	return q
}

// ObfuscateStatsGroup obfuscates the given stats bucket group.
func (o *Obfuscator) ObfuscateStatsGroup(typ, q string) string {
	switch typ {
	case "sql", "cassandra":
		oq, err := o.ObfuscateSQLString(q)
		if err != nil {
			o.opts.Log.Errorf("Error obfuscating stats group resource %q: %v", q, err)
			return nonParsableResource
		} else {
			return oq.Query
		}
	case "redis":
		return o.QuantizeRedisString(q)
	default:
		return q
	}
}

// compactWhitespaces compacts all whitespaces in t.
func compactWhitespaces(t string) string {
	n := len(t)
	r := make([]byte, n)
	spaceCode := uint8(32)
	isWhitespace := func(char uint8) bool { return char == spaceCode }
	nr := 0
	offset := 0
	for i := 0; i < n; i++ {
		if isWhitespace(t[i]) {
			copy(r[nr:], t[nr+offset:i])
			r[i-offset] = spaceCode
			nr = i + 1 - offset
			for j := i + 1; j < n; j++ {
				if !isWhitespace(t[j]) {
					offset += j - i - 1
					i = j
					break
				} else if j == n-1 {
					offset += j - i
					i = j
					break
				}
			}
		}
	}
	copy(r[nr:], t[nr+offset:n])
	r = r[:n-offset]
	return string(bytes.Trim(r, " "))
}

// replaceDigits replaces consecutive sequences of digits with '?',
// example: "jobs_2020_1597876964" --> "jobs_?_?"
func replaceDigits(buffer []byte) []byte {
	scanningDigit := false
	filtered := buffer[:0]
	for _, b := range buffer {
		// digits are encoded as 1 byte in utf8
		if isDigit(rune(b)) {
			if scanningDigit {
				continue
			}
			scanningDigit = true
			filtered = append(filtered, byte('?'))
			continue
		}
		scanningDigit = false
		filtered = append(filtered, b)
	}
	return filtered
}
