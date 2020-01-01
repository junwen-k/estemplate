// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// NormalizerCustom custom normalizer which are similar to analyzers except that they may
// only emit a single token. As a consequence, they do not have a tokenizer and only accept
// a subset of available char filters and token filters.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-normalizers.html
// for details.
type NormalizerCustom struct {
	Normalizer
	name string

	// fields specific to custom normalizer
	charFilter []string
	filter     []string
}

// NewNormalizerCustom initializes a new NormalizerCustom.
func NewNormalizerCustom(name string) *NormalizerCustom {
	return &NormalizerCustom{
		name: name,
	}
}

// Name returns field key for the Normalizer.
func (c *NormalizerCustom) Name() string {
	return c.name
}

// CharFilter sets the character filters (built-in or customised) for this normalizer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-charfilters.html
// for details.
func (c *NormalizerCustom) CharFilter(charFilter ...string) *NormalizerCustom {
	c.charFilter = append(c.charFilter, charFilter...)
	return c
}

// Filter sets the token filters (built-in or customised) for this normalizer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-tokenfilters.html
// for details.
func (c *NormalizerCustom) Filter(filter ...string) *NormalizerCustom {
	c.filter = append(c.filter, filter...)
	return c
}

// Validate validates NormalizerCustom.
func (c *NormalizerCustom) Validate(includeName bool) error {
	var invalid []string
	if includeName && c.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (c *NormalizerCustom) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "custom",
	// 		"char_filter": ["quote"],
	// 		"filter": ["lowercase", "asciifolding"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "custom"

	if len(c.charFilter) > 0 {
		options["char_filter"] = c.charFilter
	}
	if len(c.filter) > 0 {
		options["filter"] = c.filter
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
