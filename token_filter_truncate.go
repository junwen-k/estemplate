// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterTruncate token filter that truncates tokens that exceed a specified character
// limit.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-truncate-tokenfilter.html
// for details.
type TokenFilterTruncate struct {
	TokenFilter
	name string

	// fields specific to truncate token filter
	limit *int
}

// NewTokenFilterTruncate initializes a new TokenFilterTruncate.
func NewTokenFilterTruncate(name string) *TokenFilterTruncate {
	return &TokenFilterTruncate{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (t *TokenFilterTruncate) Name() string {
	return t.name
}

// Limit sets the character limit for each token. Tokens exceeding this limit are truncated.
// Defaults to 10.
func (t *TokenFilterTruncate) Limit(limit int) *TokenFilterTruncate {
	t.limit = &limit
	return t
}

// Validate validates TokenFilterTruncate.
func (t *TokenFilterTruncate) Validate(includeName bool) error {
	var invalid []string
	if includeName && t.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (t *TokenFilterTruncate) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "truncate",
	// 		"limit": 10
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "truncate"

	if t.limit != nil {
		options["limit"] = t.limit
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[t.name] = options
	return source, nil
}
