// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterMultiplexer token filter that will emit multiple tokens at the same position, each
// version of the token having been run through a different filter. Identical output tokens at the
// same position will be removed.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-multiplexer-tokenfilter.html
// for details.
type TokenFilterMultiplexer struct {
	TokenFilter
	name string

	// fields specific to multiplexer token filter
	filters          []string
	preserveOriginal *bool
}

// NewTokenFilterMultiplexer initializes a new TokenFilterMultiplexer.
func NewTokenFilterMultiplexer(name string) *TokenFilterMultiplexer {
	return &TokenFilterMultiplexer{
		name:    name,
		filters: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (m *TokenFilterMultiplexer) Name() string {
	return m.name
}

// Filters sets a list of token filters to apply to incoming tokens. These can be any token filters
// defined elsewhere in the index mappings. Filters can be chained using comma-delimited string,
// so for example "lowercase, porter_stem" would apply `lowercase` filter and then the `porter_stem`
// filter to a single token.
func (m *TokenFilterMultiplexer) Filters(filters ...string) *TokenFilterMultiplexer {
	m.filters = append(m.filters, filters...)
	return m
}

// PreserveOriginal sets whether to emit the original token in addition to the filtered tokens.
// Defaults to true.
func (m *TokenFilterMultiplexer) PreserveOriginal(preserveOriginal bool) *TokenFilterMultiplexer {
	m.preserveOriginal = &preserveOriginal
	return m
}

// Validate validates TokenFilterMultiplexer.
func (m *TokenFilterMultiplexer) Validate(includeName bool) error {
	var invalid []string
	if includeName && m.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (m *TokenFilterMultiplexer) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "multiplexer",
	// 		"filters": ["lowercase", "lowercase, porter_stem"],
	// 		"preserve_original": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "multiplexer"

	if len(m.filters) > 0 {
		options["filters"] = m.filters
	}
	if m.preserveOriginal != nil {
		options["preserve_original"] = m.preserveOriginal
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[m.name] = options
	return source, nil
}
