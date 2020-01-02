// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerNGram Partial Word Tokenizer first breaks text down into words whenever it
// encounters one of a list of specified characters, then it emits N-grams of each word
// of the specified length. They are useful for querying languages that don’t use spaces
// or that have long compound words, like German.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-ngram-tokenizer.html
// for details.
type TokenizerNGram struct {
	Tokenizer
	name string

	// fields specific to ngram tokenizer
	minGram    *int
	maxGram    *int
	tokenChars []string
}

// NewTokenizerNGram initializes a new TokenizerNGram.
func NewTokenizerNGram(name string) *TokenizerNGram {
	return &TokenizerNGram{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (n *TokenizerNGram) Name() string {
	return n.name
}

// MinGram sets the minimum length of characters in a gram.
// It usually makes sense to set `min_gram` and `max_gram` to the same value.
// The smaller the length, the more documents will match but the lower the
// quality of the matches. The longer the length, the more specific the
// matches. A tri-gram (length 3) is a good place to start.
// Defaults to 1.
func (n *TokenizerNGram) MinGram(minGram int) *TokenizerNGram {
	n.minGram = &minGram
	return n
}

// MaxGram sets the maximum length of characters in a gram.
// It usually makes sense to set `min_gram` and `max_gram` to the same value.
// The smaller the length, the more documents will match but the lower the
// quality of the matches. The longer the length, the more specific the
// matches. A tri-gram (length 3) is a good place to start.
// Defaults to 2.
func (n *TokenizerNGram) MaxGram(maxGram int) *TokenizerNGram {
	n.maxGram = &maxGram
	return n
}

// TokenChars sets the character classes that should be included in a token.
// Elasticsearch will split on characters that don't belong to the classes
// specified.
// Can be set to the following values:
// - letter. ex: (a / b / ï / 京)
// - digit. ex: (3 / 7)
// - whitespace. ex: (" " / "\n")
// - punctuation. ex: (! / ")
// - symbol. ex: ($ / √)
//
// Defaults to [] (keep all characters).
func (n *TokenizerNGram) TokenChars(tokenChars ...string) *TokenizerNGram {
	n.tokenChars = append(n.tokenChars, tokenChars...)
	return n
}

// Validate validates TokenizerNGram.
func (n *TokenizerNGram) Validate(includeName bool) error {
	var invalid []string
	if includeName && n.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(n.tokenChars) > 0 {
		for _, c := range n.tokenChars {
			if _, ok := map[string]bool{
				"letter":      true,
				"digit":       true,
				"whitespace":  true,
				"punctuation": true,
				"symbol":      true,
			}[c]; !ok {
				invalid = append(invalid, "TokenChars")
				break
			}
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (n *TokenizerNGram) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "ngram",
	// 		"min_gram": 3,
	// 		"max_gram": 3,
	// 		"token_chars": ["letter", "digit"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "ngram"

	if n.minGram != nil {
		options["min_gram"] = n.minGram
	}
	if n.maxGram != nil {
		options["max_gram"] = n.maxGram
	}
	if len(n.tokenChars) > 0 {
		options["token_chars"] = n.tokenChars
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[n.name] = options
	return source, nil
}
