// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterEdgeNGram token filter that forms an n-gram of a specified length
// from the beginning of a token. For example, you can use the `edge_ngram` token
// filter to change "quick" to "qu".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-edgengram-tokenfilter.html
// for details.
type TokenFilterEdgeNGram struct {
	TokenFilter
	name string

	// fields specific to edge ngram token filter
	maxGram *int
	minGram *int
	side    string
}

// NewTokenFilterEdgeNGram initializes a new TokenFilterEdgeNGram.
func NewTokenFilterEdgeNGram(name string) *TokenFilterEdgeNGram {
	return &TokenFilterEdgeNGram{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (g *TokenFilterEdgeNGram) Name() string {
	return g.name
}

// MaxGram sets the maximum character length of a gram.
// Defaults to 2 for custom token filters.
// Defaults to 1 for built-in `edge_ngram` filter.
func (g *TokenFilterEdgeNGram) MaxGram(maxGram int) *TokenFilterEdgeNGram {
	g.maxGram = &maxGram
	return g
}

// MinGram sets the minimum character length of a gram.
// Defaults to 1.
func (g *TokenFilterEdgeNGram) MinGram(minGram int) *TokenFilterEdgeNGram {
	g.minGram = &minGram
	return g
}

// Side sets whether to truncate tokens from the `front` or `back`.
// Can be set to the following values:
// "front"
// "back"
// ! Deprecated, use `reverse` token filter before and after the `edge_ngram`
// filter to achieve the same results.
func (g *TokenFilterEdgeNGram) Side(side string) *TokenFilterEdgeNGram {
	g.side = side
	return g
}

// Validate validates TokenFilterEdgeNGram.
func (g *TokenFilterEdgeNGram) Validate(includeName bool) error {
	var invalid []string
	if includeName && g.name == "" {
		invalid = append(invalid, "Name")
	}
	if g.side != "" {
		if _, valid := map[string]bool{
			"front": true,
			"back":  true,
		}[g.side]; !valid {
			invalid = append(invalid, "Side")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (g *TokenFilterEdgeNGram) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "edge_ngram",
	// 		"max_gram": 1,
	// 		"min_gram": 1,
	// 		"side": "front"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "edge_ngram"

	if g.maxGram != nil {
		options["max_gram"] = g.maxGram
	}
	if g.minGram != nil {
		options["min_gram"] = g.minGram
	}
	if g.side != "" {
		options["side"] = g.side
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[g.name] = options
	return source, nil
}
