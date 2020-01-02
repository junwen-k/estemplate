// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerThai Word Orientated Tokenizer that segments Thai text into words, using the Thai
// segmentation algorithm included with Java. Text in other languages in general will be treated
// the same as the standard tokenizer.
// ! This tokenizer may not be supported by all JREs.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-thai-tokenizer.html
// for details.
type TokenizerThai struct {
	Tokenizer
	name string

	// fields specific to letter thai
}

// NewTokenizerThai initializes a new TokenizerThai.
func NewTokenizerThai(name string) *TokenizerThai {
	return &TokenizerThai{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (t *TokenizerThai) Name() string {
	return t.name
}

// Validate validates TokenizerThai.
func (t *TokenizerThai) Validate(includeName bool) error {
	var invalid []string
	if includeName && t.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (t *TokenizerThai) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "thai"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "thai"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[t.name] = options
	return source, nil
}
