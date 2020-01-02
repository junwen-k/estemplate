// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerClassic Word Orientated Tokenizer grammer based tokenizer that is good for English
// language documents. This tokenizer has heuristics for special treatment of acronyms, company names,
// email addresses, and internet host names. However, these rules don’t always work, and the tokenizer
// doesn’t work well for most languages other than English:
//
// - It splits words at most punctuation characters, removing punctuation. However, a dot that’s not
//   followed by whitespace is considered part of a token.
//
// - It splits words at hyphens, unless there’s a number in the token, in which case the whole token is
//   interpreted as a product number and is not split.
//
// - It recognizes email addresses and internet hostnames as one token.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-classic-tokenizer.html
// for details.
type TokenizerClassic struct {
	Tokenizer
	name string

	// fields specific to classic tokenizer
	maxTokenLength *int
}

// NewTokenizerClassic initializes a new TokenizerClassic.
func NewTokenizerClassic(name string) *TokenizerClassic {
	return &TokenizerClassic{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (c *TokenizerClassic) Name() string {
	return c.name
}

// MaxTokenLength sets the maximum token length and if a token is seen that exceeds this
// length then it is split at `max_token_length` intervals.
// Defaults to 255.
func (c *TokenizerClassic) MaxTokenLength(maxTokenLength int) *TokenizerClassic {
	c.maxTokenLength = &maxTokenLength
	return c
}

// Validate validates TokenizerClassic.
func (c *TokenizerClassic) Validate(includeName bool) error {
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
func (c *TokenizerClassic) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "classic",
	// 		"max_token_length": 255
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "classic"

	if c.maxTokenLength != nil {
		options["max_token_length"] = c.maxTokenLength
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
