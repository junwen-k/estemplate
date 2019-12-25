// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// AnalyzerCustom custom analyzer which uses the appropriate combination of:
// - zero or more character filters
// - a tokenizer
// - zero or more token filters
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-custom-analyzer.html
// for details.
type AnalyzerCustom struct {
	Analyzer
	name string

	// fields specific to custom analyzer
	tokenizer            string
	charFilter           []string
	filter               []string
	positionIncrementGap *int
}

// NewAnalyzerCustom initializes a new AnalyzerCustom.
func NewAnalyzerCustom(name, tokenizer string) *AnalyzerCustom {
	return &AnalyzerCustom{
		name:      name,
		tokenizer: tokenizer,
	}
}

// Name returns field key for the Analyzer.
func (c *AnalyzerCustom) Name() string {
	return c.name
}

// Tokenizer sets the tokenizer (built-in or customised) for this analyzer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-tokenizers.html
// for details.
func (c *AnalyzerCustom) Tokenizer(tokenizer string) *AnalyzerCustom {
	c.tokenizer = tokenizer
	return c
}

// CharFilter sets the character filters (built-in or customised) for this analzyer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-charfilters.html
// for details.
func (c *AnalyzerCustom) CharFilter(charFilter ...string) *AnalyzerCustom {
	c.charFilter = append(c.charFilter, charFilter...)
	return c
}

// Filter sets the token filters (built-in or customised) for this analzyer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-tokenfilters.html
// for details.
func (c *AnalyzerCustom) Filter(filter ...string) *AnalyzerCustom {
	c.filter = append(c.filter, filter...)
	return c
}

// PositionIncrementGap sets the number of fake term position which should be inserted between
// each element of an array of strings.
// Defaults to the settings of `position_increment_gap` which defaults to 100.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/position-increment-gap.html
// for details.
func (c *AnalyzerCustom) PositionIncrementGap(positionIncrementGap int) *AnalyzerCustom {
	c.positionIncrementGap = &positionIncrementGap
	return c
}

// Validate validates AnalyzerCustom.
func (c *AnalyzerCustom) Validate(includeName bool) error {
	var invalid []string
	if includeName && c.name == "" {
		invalid = append(invalid, "Name")
	}
	if c.tokenizer == "" {
		invalid = append(invalid, "Tokenizer")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (c *AnalyzerCustom) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "custom",
	// 		"tokenizer": "standard",
	// 		"char_filter": ["html_strip"],
	// 		"filter": ["lowercase", "asciifolding"],
	// 		"position_increment_gap": 1
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "custom"

	if c.tokenizer != "" {
		options["tokenizer"] = c.tokenizer
	}
	if len(c.charFilter) > 0 {
		options["char_filter"] = c.charFilter
	}
	if len(c.filter) > 0 {
		options["filter"] = c.filter
	}
	if c.positionIncrementGap != nil {
		options["position_increment_gap"] = c.positionIncrementGap
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
