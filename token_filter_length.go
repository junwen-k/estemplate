// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterLength token filter that removes tokens shorter or longer than specified
// character lengths. For example, you can use the `length` filter to exclude tokens
// shorter than 2 characters and tokens longer than 5 characters.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-length-tokenfilter.html
// for details.
type TokenFilterLength struct {
	TokenFilter
	name string

	// fields specific to length token filter
	min *int
	max *int
}

// NewTokenFilterLength initializes a new TokenFilterLength.
func NewTokenFilterLength(name string) *TokenFilterLength {
	return &TokenFilterLength{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (l *TokenFilterLength) Name() string {
	return l.name
}

// Min sets the minimum character length of a token. Shorter tokens are excluded from
// the output.
// Defaults to 0.
func (l *TokenFilterLength) Min(min int) *TokenFilterLength {
	l.min = &min
	return l
}

// Max sets the maximum character length of a token. Longer tokens are excluded from
// the output.
// Defaults to `Integer.MAX_VALUE`, which is 2^31-1 or 2147483647.
func (l *TokenFilterLength) Max(max int) *TokenFilterLength {
	l.max = &max
	return l
}

// Validate validates TokenFilterLength.
func (l *TokenFilterLength) Validate(includeName bool) error {
	var invalid []string
	if includeName && l.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (l *TokenFilterLength) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "length",
	// 		"min": 2,
	// 		"max": 10
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "length"

	if l.min != nil {
		options["min"] = l.min
	}
	if l.max != nil {
		options["max"] = l.max
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[l.name] = options
	return source, nil
}
