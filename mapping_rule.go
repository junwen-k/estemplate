// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"fmt"
	"strings"
)

// MappingRule standard mapping rule of Elasticsearch in the form of `key => value` pair,
// `key, key => value, value, value`, "key, key" or "value, value".
type MappingRule struct {
	key   []string
	value []string
}

// NewMappingRule initializes a new MappingRule.
func NewMappingRule(key, value string) *MappingRule {
	r := &MappingRule{
		key:   make([]string, 0),
		value: make([]string, 0),
	}
	if key != "" {
		r.key = []string{key}
	}
	if value != "" {
		r.value = []string{value}
	}
	return r
}

// Key sets the key for the mapping rule.
func (r *MappingRule) Key(key ...string) *MappingRule {
	r.key = append(r.key, key...)
	return r
}

// Value sets the value for the mapping rule.
func (r *MappingRule) Value(value ...string) *MappingRule {
	r.value = append(r.value, value...)
	return r
}

// Source returns the serializable JSON for the source builder.
func (r *MappingRule) Source() (interface{}, error) {
	// "value, value"
	// "key, key"
	// "key, key => value, value, value"
	var source interface{}

	var key interface{}
	switch {
	case len(r.key) > 1:
		key = strings.Join(r.key, ", ")
		break
	case len(r.key) == 1:
		key = r.key[0]
		break
	default:
		key = ""
	}

	var value interface{}
	switch {
	case len(r.value) > 1:
		value = strings.Join(r.value, ", ")
		break
	case len(r.value) == 1:
		value = r.value[0]
		break
	default:
		value = ""
	}

	if key != "" {
		source = key
	}
	if value != "" {
		source = value
	}
	if key != "" && value != "" {
		source = fmt.Sprintf("%s => %s", key, value)
	}

	return source, nil
}
