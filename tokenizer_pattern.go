// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"fmt"
	"strings"
)

// TokenizerPattern Structured Text Tokenizer that uses a regular expression to either split text
// into terms whenever it matches a word separator, or to capture matching text as terms. The regular
// expression defaults to `\W+`, which splits text whenever it encounters non-word characters.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-pattern-tokenizer.html
// for details.
type TokenizerPattern struct {
	Tokenizer
	name string

	// fields specific to pattern tokenizer
	pattern string
	flags   []string
	group   *int
}

// NewTokenizerPattern initializes a new TokenizerPattern.
func NewTokenizerPattern(name string) *TokenizerPattern {
	return &TokenizerPattern{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (p *TokenizerPattern) Name() string {
	return p.name
}

// Pattern sets the Java regular expression for the tokenizer.
// Defaults to `\W+`.
//
// See https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html
// for details.
func (p *TokenizerPattern) Pattern(pattern string) *TokenizerPattern {
	p.pattern = pattern
	return p
}

// Flags sets Java regular expression flags. eg "CASE_INSENSITIVE|COMMENTS".
//
// See https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html#field.summary
// for details.
func (p *TokenizerPattern) Flags(flags ...string) *TokenizerPattern {
	p.flags = append(p.flags, flags...)
	return p
}

// Group sets which capture group to extract as tokens.
// Defaults to -1 (split).
func (p *TokenizerPattern) Group(group int) *TokenizerPattern {
	p.group = &group
	return p
}

// Validate validates TokenizerPattern.
func (p *TokenizerPattern) Validate(includeName bool) error {
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
func (p *TokenizerPattern) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "pattern",
	// 		"max_token_length": 255
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "pattern"

	if p.pattern != "" {
		options["pattern"] = p.pattern
	}
	if len(p.flags) > 0 {
		options["flags"] = strings.Join(p.flags, "|")
	}
	if p.group != nil {
		options["group"] = p.group
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
