// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SlowlogThreshold shard level slow search log allows log slow search (query and fetch phases) into a
// dedicated log file.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-slowlog.html
// for details.
type SlowlogThreshold struct {
	slowlogType string
	phase       string
	level       string
	value       string
}

// NewSlowlogThreshold initalizes a new SlowlogThreshold.
func NewSlowlogThreshold(slowlogType, phase, level, value string) *SlowlogThreshold {
	return &SlowlogThreshold{
		slowlogType: slowlogType,
		phase:       phase,
		level:       level,
		value:       value,
	}
}

// NewSearchSlowlogThreshold initalizes a new search SlowlogThreshold.
func NewSearchSlowlogThreshold(phase, level, value string) *SlowlogThreshold {
	return NewSlowlogThreshold("search", phase, level, value)
}

// NewIndexSlowlogThreshold initalizes a new index SlowlogThreshold.
func NewIndexSlowlogThreshold(phase, level, value string) *SlowlogThreshold {
	return NewSlowlogThreshold("indexing", phase, level, value)
}

// Level sets the level for this slowlog threshold.
// Can be set to the following values:
// "warn"
// "info"
// "debug"
// "trace"
func (t *SlowlogThreshold) Level(level string) *SlowlogThreshold {
	t.level = level
	return t
}

// Phase sets the phase for this slowlog threshold.
// Can be set to the following values:
// "query" - query phase of the execution
// "fetch" - fetch phase
// "index" - indexing phase
func (t *SlowlogThreshold) Phase(phase string) *SlowlogThreshold {
	t.phase = phase
	return t
}

// SlowlogType sets the slowlog type for this slowlog threshold.
// Can be set to the following values:
// "search" - For search slowlog.
// "indexing" - For indexing slowlog.
func (t *SlowlogThreshold) SlowlogType(slowlogType string) *SlowlogThreshold {
	t.slowlogType = slowlogType
	return t
}

// Value sets the value for this slowlog threshold.
// Defaults to "-1" which disable the slowlog.
func (t *SlowlogThreshold) Value(value string) *SlowlogThreshold {
	t.value = value
	return t
}

// Source returns the serializable JSON for the source builder.
func (t *SlowlogThreshold) Source(includeName bool) (interface{}, error) {
	// {
	// 	"search": {
	// 		"slowlog.threshold.query.warn": "10s",
	// 		"slowlog.threshold.query.info": "5s",
	// 		"slowlog.threshold.query.query": "2s",
	// 		"slowlog.threshold.query.trace": "500ms"
	// 	}
	// }
	options := make(map[string]interface{})

	if t.phase != "" && t.level != "" && t.value != "" {
		options[fmt.Sprintf("slowlog.threshold.%s.%s", t.phase, t.level)] = t.value
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[t.slowlogType] = options
	return source, nil
}
