// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerSimplePattern Structured Text Tokenizer that uses a regular expression to capture
// matching text as terms. The set of regular expression features it supports is more limited than
// the `pattern` tokenizer, but the tokenization is generally faster.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-simplepattern-tokenizer.html
// for details.
type TokenizerSimplePattern struct {
	Tokenizer
	name string

	// fields specific to simple pattern tokenizer
	pattern string
}

// NewTokenizerSimplePattern initializes a new TokenizerSimplePattern.
func NewTokenizerSimplePattern(name string) *TokenizerSimplePattern {
	return &TokenizerSimplePattern{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (p *TokenizerSimplePattern) Name() string {
	return p.name
}

// Pattern sets the Lucene regular expression for the tokenizer, pattern should be always
// configured with a non-default pattern.
// Defaults to "" (empty string).
//
// See http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/util/automaton/RegExp.html
// for details.
func (p *TokenizerSimplePattern) Pattern(pattern string) *TokenizerSimplePattern {
	p.pattern = pattern
	return p
}

// Validate validates TokenizerSimplePattern.
func (p *TokenizerSimplePattern) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *TokenizerSimplePattern) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "simple_pattern",
	// 		"pattern": "[0123456789]{3}"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "simple_pattern"

	if p.pattern != "" {
		options["pattern"] = p.pattern
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
