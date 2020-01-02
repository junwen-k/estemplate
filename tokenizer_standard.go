// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerStandard Word Orientated Tokenizer that provides grammar based tokenization (based on the Unicode
// Text Segmentation algorithm, as specified in Unicode Standard Annex #29) and works well for most
// languages.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-standard-tokenizer.html
// for details.
type TokenizerStandard struct {
	Tokenizer
	name string

	// fields specific to standard tokenizer
	maxTokenLength *int
}

// NewTokenizerStandard initializes a new TokenizerStandard.
func NewTokenizerStandard(name string) *TokenizerStandard {
	return &TokenizerStandard{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (s *TokenizerStandard) Name() string {
	return s.name
}

// MaxTokenLength sets the maximum token length and if a token is seen that exceeds this
// length then it is split at `max_token_length` intervals.
// Defaults to 255.
func (s *TokenizerStandard) MaxTokenLength(maxTokenLength int) *TokenizerStandard {
	s.maxTokenLength = &maxTokenLength
	return s
}

// Validate validates TokenizerStandard.
func (s *TokenizerStandard) Validate(includeName bool) error {
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
func (s *TokenizerStandard) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "standard",
	// 		"max_token_length": 255
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "standard"

	if s.maxTokenLength != nil {
		options["max_token_length"] = s.maxTokenLength
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
