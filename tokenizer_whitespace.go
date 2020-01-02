// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerWhitespace Word Orientated Tokenizer that breaks text into terms whenever it encounters
// a whitespace character.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-whitespace-tokenizer.html
// for details.
type TokenizerWhitespace struct {
	Tokenizer
	name string

	// fields specific to whitespace tokenizer
	maxTokenLength *int
}

// NewTokenizerWhitespace initializes a new TokenizerWhitespace.
func NewTokenizerWhitespace(name string) *TokenizerWhitespace {
	return &TokenizerWhitespace{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (w *TokenizerWhitespace) Name() string {
	return w.name
}

// MaxTokenLength sets the maximum token length and if a token is seen that exceeds this
// length then it is split at `max_token_length` intervals.
// Defaults to 255.
func (w *TokenizerWhitespace) MaxTokenLength(maxTokenLength int) *TokenizerWhitespace {
	w.maxTokenLength = &maxTokenLength
	return w
}

// Validate validates TokenizerWhitespace.
func (w *TokenizerWhitespace) Validate(includeName bool) error {
	var invalid []string
	if includeName && w.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (w *TokenizerWhitespace) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "whitespace",
	// 		"max_token_length": 255
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "whitespace"

	if w.maxTokenLength != nil {
		options["max_token_length"] = w.maxTokenLength
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[w.name] = options
	return source, nil
}
