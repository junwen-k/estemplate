// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerLetter Word Orientated Tokenizer that breaks text into terms whenever it encounters a character
// which is not a letter. It does a reasonable job for most European languages, but does a
// terrible job for some Asian languages, where words are not separated by spaces.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-letter-tokenizer.html
// for details.
type TokenizerLetter struct {
	Tokenizer
	name string

	// fields specific to letter tokenizer
}

// NewTokenizerLetter initializes a new TokenizerLetter.
func NewTokenizerLetter(name string) *TokenizerLetter {
	return &TokenizerLetter{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (l *TokenizerLetter) Name() string {
	return l.name
}

// Validate validates TokenizerLetter.
func (l *TokenizerLetter) Validate(includeName bool) error {
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
func (l *TokenizerLetter) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "letter"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "letter"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[l.name] = options
	return source, nil
}
