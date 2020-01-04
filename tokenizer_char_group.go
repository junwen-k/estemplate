// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerCharGroup Structured Text Tokenizer that breaks text into terms whever it encounters a
// character which is in a defined set. It is mostly useful for cases where a simple custom tokenization
// is desired, and the overhead of use of the `pattern` tokenizer is not acceptable.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-chargroup-tokenizer.html
// for details.
type TokenizerCharGroup struct {
	Tokenizer
	name string

	// fields specific to char group tokenizer
	tokenizeOnChars []string
}

// NewTokenizerCharGroup initializes a new TokenizerCharGroup.
func NewTokenizerCharGroup(name string) *TokenizerCharGroup {
	return &TokenizerCharGroup{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (g *TokenizerCharGroup) Name() string {
	return g.name
}

// TokenizeOnChars sets a list containing a list of characters to tokenize the string on.
// This accepts either a single character like e.g. "-", or character groups (whitespace, letter
// digit, punctuation, symbol).
func (g *TokenizerCharGroup) TokenizeOnChars(tokenizeOnChars ...string) *TokenizerCharGroup {
	g.tokenizeOnChars = append(g.tokenizeOnChars, tokenizeOnChars...)
	return g
}

// Validate validates TokenizerCharGroup.
func (g *TokenizerCharGroup) Validate(includeName bool) error {
	var invalid []string
	if includeName && g.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (g *TokenizerCharGroup) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "char_group",
	// 		"tokenize_on_chars": [
	// 			"whitespace",
	// 			"-",
	// 			"\n"
	// 		]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "char_group"

	if len(g.tokenizeOnChars) > 0 {
		options["tokenize_on_chars"] = g.tokenizeOnChars
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[g.name] = options
	return source, nil
}
