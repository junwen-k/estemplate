// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerLowercase Word Orientated Tokenizer which like the letter tokenizer that breaks text into terms whenever
// it encounters a character which is not a letter, but also lowercases all terms. It does a reasonable job
// for most European languages, but does a terrible job for some Asian languages, where words are not separated
// by spaces.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-lowercase-tokenizer.html
// for details.
type TokenizerLowercase struct {
	Tokenizer
	name string

	// fields specific to lowercase tokenizer
}

// NewTokenizerLowercase initializes a new TokenizerLowercase.
func NewTokenizerLowercase(name string) *TokenizerLowercase {
	return &TokenizerLowercase{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (l *TokenizerLowercase) Name() string {
	return l.name
}

// Validate validates TokenizerLowercase.
func (l *TokenizerLowercase) Validate(includeName bool) error {
	var invalid []string
	if includeName && l.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (l *TokenizerLowercase) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "lowercase"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "lowercase"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[l.name] = options
	return source, nil
}
