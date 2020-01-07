// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterPatternCapture token filter that emits a oktne for every capture
// group in the regular expression. Patterns are not anchored to the beginning
// and end of the string, so each pattern can match multiple times, and matches
// are allowed to overlap.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-pattern-capture-tokenfilter.html
// for details.
type TokenFilterPatternCapture struct {
	TokenFilter
	name string

	// fields specific to pattern capture token filter
	preserveOriginal *bool
	patterns         []string
}

// NewTokenFilterPatternCapture initializes a new TokenFilterPatternCapture.
func NewTokenFilterPatternCapture(name string) *TokenFilterPatternCapture {
	return &TokenFilterPatternCapture{
		name:     name,
		patterns: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (c *TokenFilterPatternCapture) Name() string {
	return c.name
}

// PreserveOriginal sets whether to emit the original token.
// Defaults to true.
func (c *TokenFilterPatternCapture) PreserveOriginal(preserveOriginal bool) *TokenFilterPatternCapture {
	c.preserveOriginal = &preserveOriginal
	return c
}

// Patterns sets the regular expressions for the filter.
func (c *TokenFilterPatternCapture) Patterns(patterns ...string) *TokenFilterPatternCapture {
	c.patterns = append(c.patterns, patterns...)
	return c
}

// Validate validates TokenFilterPatternCapture.
func (c *TokenFilterPatternCapture) Validate(includeName bool) error {
	var invalid []string
	if includeName && c.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (c *TokenFilterPatternCapture) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "pattern_capture",
	// 		"preserve_original": true,
	// 		"patterns": ["(\\p{Ll}+|\\p{Lu}\\p{Ll}+|\\p{Lu}+)", "(\\d+)"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "pattern_capture"

	if c.preserveOriginal != nil {
		options["preserve_original"] = c.preserveOriginal
	}
	if len(c.patterns) > 0 {
		options["patterns"] = c.patterns
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
