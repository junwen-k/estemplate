// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerSimplePatternSplit Structured Text Tokenizer that uses a regular expression to split
// the input into terms at pattern matches. The set of regular expression features it supports is
// more limited than the `pattern` tokenizer, but the tokenization is generally faster.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-simplepatternsplit-tokenizer.html
// for details.
type TokenizerSimplePatternSplit struct {
	Tokenizer
	name string

	// fields specific to simple pattern split tokenizer
	pattern string
}

// NewTokenizerSimplePatternSplit initializes a new TokenizerSimplePatternSplit.
func NewTokenizerSimplePatternSplit(name string) *TokenizerSimplePatternSplit {
	return &TokenizerSimplePatternSplit{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (s *TokenizerSimplePatternSplit) Name() string {
	return s.name
}

// Pattern sets the Lucene regular expression for the tokenizer.
// Defaults to "" (empty string).
//
// See http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/util/automaton/RegExp.html
// for details.
func (s *TokenizerSimplePatternSplit) Pattern(pattern string) *TokenizerSimplePatternSplit {
	s.pattern = pattern
	return s
}

// Validate validates TokenizerSimplePatternSplit.
func (s *TokenizerSimplePatternSplit) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *TokenizerSimplePatternSplit) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "simple_pattern_split",
	// 		"pattern": "_"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "simple_pattern_split"

	if s.pattern != "" {
		options["pattern"] = s.pattern
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
