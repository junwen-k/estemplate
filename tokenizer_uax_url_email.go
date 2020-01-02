// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerUAXURLEmail Word Orientated Tokenizer which is like the standard tokenizer
// except that it recognises URLs and email addresses as single tokens.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-uaxurlemail-tokenizer.html
// for details.
type TokenizerUAXURLEmail struct {
	Tokenizer
	name string

	// fields specific to uax url email tokenizer
	maxTokenLength *int
}

// NewTokenizerUAXURLEmail initializes a new TokenizerUAXURLEmail.
func NewTokenizerUAXURLEmail(name string) *TokenizerUAXURLEmail {
	return &TokenizerUAXURLEmail{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (u *TokenizerUAXURLEmail) Name() string {
	return u.name
}

// MaxTokenLength sets the maximum token length and if a token is seen that exceeds this
// length then it is split at `max_token_length` intervals.
// Defaults to 255.
func (u *TokenizerUAXURLEmail) MaxTokenLength(maxTokenLength int) *TokenizerUAXURLEmail {
	u.maxTokenLength = &maxTokenLength
	return u
}

// Validate validates TokenizerUAXURLEmail.
func (u *TokenizerUAXURLEmail) Validate(includeName bool) error {
	var invalid []string
	if includeName && u.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (u *TokenizerUAXURLEmail) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "uax_url_email",
	// 		"max_token_length": 255
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "uax_url_email"

	if u.maxTokenLength != nil {
		options["max_token_length"] = u.maxTokenLength
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[u.name] = options
	return source, nil
}
