// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerKeyword Structured Text Tokenizer "noop" tokenizer that accepts whatever text
// it is given and outputs the exact same text as a single term. It can be combined with
// token filters to normalise output, e.g. lower-casing email addresses.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-keyword-tokenizer.html
// for details.
type TokenizerKeyword struct {
	Tokenizer
	name string

	// fields specific to keyword tokenizer
	bufferSize *int
}

// NewTokenizerKeyword initializes a new TokenizerKeyword.
func NewTokenizerKeyword(name string) *TokenizerKeyword {
	return &TokenizerKeyword{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (k *TokenizerKeyword) Name() string {
	return k.name
}

// BufferSize sets the number of characters read into the term buffer in a single
// pass. The term buffer will grow by this size until all the text has been consumed.
// It is advisable not to change this setting.
// Defaults to 256.
func (k *TokenizerKeyword) BufferSize(bufferSize int) *TokenizerKeyword {
	k.bufferSize = &bufferSize
	return k
}

// Validate validates TokenizerKeyword.
func (k *TokenizerKeyword) Validate(includeName bool) error {
	var invalid []string
	if includeName && k.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (k *TokenizerKeyword) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "keyword",
	// 		"buffer_size": 256
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "keyword"

	if k.bufferSize != nil {
		options["buffer_size"] = k.bufferSize
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[k.name] = options
	return source, nil
}
