// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterKeywordMarker token filter that protects words from being modified by stemmers.
// Must be placed before any stemming filters.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-keyword-marker-tokenfilter.html
// for details.
type TokenFilterKeywordMarker struct {
	TokenFilter
	name string

	// fields specific to keyword marker token filter
	keywords        []string
	keywordsPath    string
	keywordsPattern string
	ignoreCase      *bool
}

// NewTokenFilterKeywordMarker initializes a new TokenFilterKeywordMarker.
func NewTokenFilterKeywordMarker(name string) *TokenFilterKeywordMarker {
	return &TokenFilterKeywordMarker{
		name:     name,
		keywords: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (m *TokenFilterKeywordMarker) Name() string {
	return m.name
}

// Keywords sets a list of words to use.
func (m *TokenFilterKeywordMarker) Keywords(keywords ...string) *TokenFilterKeywordMarker {
	m.keywords = append(m.keywords, keywords...)
	return m
}

// KeywordsPath sets a path to a file containing a list of words. This path must be absolute
// or relative to the `config` location, and the file must be UTF-8 encoded. Each word
// in the file must be separated by a line break.
func (m *TokenFilterKeywordMarker) KeywordsPath(keywordsPath string) *TokenFilterKeywordMarker {
	m.keywordsPath = keywordsPath
	return m
}

// KeywordsPattern sets a regular expression to match against words in the text.
func (m *TokenFilterKeywordMarker) KeywordsPattern(keywordsPattern string) *TokenFilterKeywordMarker {
	m.keywordsPattern = keywordsPattern
	return m
}

// IgnoreCase sets whether to lowercase all words first.
// Defaults to false.
func (m *TokenFilterKeywordMarker) IgnoreCase(ignoreCase bool) *TokenFilterKeywordMarker {
	m.ignoreCase = &ignoreCase
	return m
}

// Validate validates TokenFilterKeywordMarker.
func (m *TokenFilterKeywordMarker) Validate(includeName bool) error {
	var invalid []string
	if includeName && m.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (m *TokenFilterKeywordMarker) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "keyword_marker",
	// 		"keywords": ["one", "two", "three"],
	// 		"keywords_path": "analysis/example_word_list.txt",
	// 		"keywords_pattern": "^profit_\\d+$",
	// 		"ignore_case": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "keyword_marker"

	if len(m.keywords) > 0 {
		options["keywords"] = m.keywords
	}
	if m.keywordsPath != "" {
		options["keywords_path"] = m.keywordsPath
	}
	if m.keywordsPattern != "" {
		options["keywords_pattern"] = m.keywordsPattern
	}
	if m.ignoreCase != nil {
		options["ignore_case"] = m.ignoreCase
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[m.name] = options
	return source, nil
}
