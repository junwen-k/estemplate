// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterNGram token filter that forms n-grams of specified lengths from a token.
// For example, you can use the `ngram` token filter to change "fox" to ["f", "fo", "o",
// "ox", "x"].
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-ngram-tokenfilter.html
// for details.
type TokenFilterNGram struct {
	TokenFilter
	name string

	// fields specific to ngram token filter
	maxGram *int
	minGram *int
}

// NewTokenFilterNGram initializes a new TokenFilterNGram.
func NewTokenFilterNGram(name string) *TokenFilterNGram {
	return &TokenFilterNGram{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (g *TokenFilterNGram) Name() string {
	return g.name
}

// MaxGram sets the maximum length of characters in a gram.
// Defaults to 2.
func (g *TokenFilterNGram) MaxGram(maxGram int) *TokenFilterNGram {
	g.maxGram = &maxGram
	return g
}

// MinGram sets the minimum length of characters in a gram.
// Defaults to 1.
func (g *TokenFilterNGram) MinGram(minGram int) *TokenFilterNGram {
	g.minGram = &minGram
	return g
}

// Validate validates TokenFilterNGram.
func (g *TokenFilterNGram) Validate(includeName bool) error {
	var invalid []string
	if includeName && g.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (g *TokenFilterNGram) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "ngram",
	// 		"max_gram": 5,
	// 		"min_gram": 3
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "ngram"

	if g.maxGram != nil {
		options["max_gram"] = g.maxGram
	}
	if g.minGram != nil {
		options["min_gram"] = g.minGram
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[g.name] = options
	return source, nil
}
