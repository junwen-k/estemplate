// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenizerEdgeNGram Partial Word Tokenizer first breaks text down into words whenever it
// encounters one of a list of specified characters, then it emits N-grams of each word
// where the start of the N-gram is anchored to the beginning of the word.
// Edge N-Grams are useful for search-as-you-type queries.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-edgengram-tokenizer.html
// for details.
type TokenizerEdgeNGram struct {
	Tokenizer
	name string

	// fields specific to edge ngram tokenizer
	minGram    *int
	maxGram    *int
	tokenChars []string
}

// NewTokenizerEdgeNGram initializes a new TokenizerEdgeNGram.
func NewTokenizerEdgeNGram(name string) *TokenizerEdgeNGram {
	return &TokenizerEdgeNGram{
		name:       name,
		tokenChars: make([]string, 0),
	}
}

// Name returns field key for the Tokenizer.
func (e *TokenizerEdgeNGram) Name() string {
	return e.name
}

// MinGram sets the minimum length of characters in a gram.
// Defaults to 1.
func (e *TokenizerEdgeNGram) MinGram(minGram int) *TokenizerEdgeNGram {
	e.minGram = &minGram
	return e
}

// MaxGram sets the maximum length of characters in a gram.
// Defaults to 2.
//
// ! Limitations of the `max_gram` parameter for Edge N-Gram tokenizer.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/max-gram-limits.html
// for details.
func (e *TokenizerEdgeNGram) MaxGram(maxGram int) *TokenizerEdgeNGram {
	e.maxGram = &maxGram
	return e
}

// TokenChars sets the character classes that should be included in a token.
// Elasticsearch will split on characters that don't belong to the classes
// specified.
// Can be set to the following values:
// letter - ex: (a / b / ï / 京)
// digit - ex: (3 / 7)
// whitespace - ex: (" " / "\n")
// punctuation - ex: (! / ")
// symbol - ex: ($ / √)
//
// Defaults to [] (keep all characters).
func (e *TokenizerEdgeNGram) TokenChars(tokenChars ...string) *TokenizerEdgeNGram {
	e.tokenChars = append(e.tokenChars, tokenChars...)
	return e
}

// Validate validates TokenizerEdgeNGram.
func (e *TokenizerEdgeNGram) Validate(includeName bool) error {
	var invalid []string
	if includeName && e.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(e.tokenChars) > 0 {
		for _, c := range e.tokenChars {
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
func (e *TokenizerEdgeNGram) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "edge_ngram",
	// 		"min_gram": 3,
	// 		"max_gram": 3,
	// 		"token_chars": ["letter", "digit"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "edge_ngram"

	if e.minGram != nil {
		options["min_gram"] = e.minGram
	}
	if e.maxGram != nil {
		options["max_gram"] = e.maxGram
	}
	if len(e.tokenChars) > 0 {
		options["token_chars"] = e.tokenChars
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[e.name] = options
	return source, nil
}
