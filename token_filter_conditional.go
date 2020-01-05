// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterConditional token filter that applies a set of token filters to tokens
// that match conditions in a provided predicate script.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-condition-tokenfilter.html
// for details.
type TokenFilterConditional struct {
	TokenFilter
	name string

	// fields specific to conditional token filter
	filter []string
	script *Script
}

// NewTokenFilterConditional initializes a new TokenFilterConditional.
func NewTokenFilterConditional(name string) *TokenFilterConditional {
	return &TokenFilterConditional{
		name:   name,
		filter: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (c *TokenFilterConditional) Name() string {
	return c.name
}

// Filter sets a list of token filters and if a token matches the predicate script
// in the `script` parameter, these filters are applied to the token in the order
// provided. These filters can include custom token filters defined in the index
// mapping.
func (c *TokenFilterConditional) Filter(filter ...string) *TokenFilterConditional {
	c.filter = append(c.filter, filter...)
	return c
}

// Script sets the script to be used to apply token filters. If a token matches this
// script, the filters in the `filter` parameter are applied to the token.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/modules-scripting-using.html#_script_parameters
// for details.
func (c *TokenFilterConditional) Script(script *Script) *TokenFilterConditional {
	c.script = script
	return c
}

// Validate validates TokenFilterConditional.
func (c *TokenFilterConditional) Validate(includeName bool) error {
	var invalid []string
	if includeName && c.name == "" {
		invalid = append(invalid, "Name")
	}
	if !(len(c.filter) > 0) {
		invalid = append(invalid, "Filter")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (c *TokenFilterConditional) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "condition",
	// 		"filter": ["lowercase", "asciifolding"],
	// 		"script": {
	// 			"source": "token.getTerm().length() < threshold",
	// 			"params": {
	// 				"threshold": 5
	// 			}
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "condition"

	if len(c.filter) > 0 {
		options["filter"] = c.filter
	}
	if c.script != nil {
		script, err := c.script.Source(false)
		if err != nil {
			return nil, err
		}
		options["script"] = script
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
