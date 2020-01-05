// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterCommonGrams token filter that generates bigrams for a specified set of
// common words.
// For example, this filter converts [the, quick, fox, is, brown] to [the, the_quick,
// quick, fox, fox_is, is, is_brown, brown] if your specify "is" and "the" as the
// common words.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-common-grams-tokenfilter.html
// for details.
type TokenFilterCommonGrams struct {
	TokenFilter
	name string

	// fields specific to common grams token filter
	commonWords     []string
	commonWordsPath string
	ignoreCase      *bool
	queryMode       *bool
}

// NewTokenFilterCommonGrams initializes a new TokenFilterCommonGrams.
func NewTokenFilterCommonGrams(name string) *TokenFilterCommonGrams {
	return &TokenFilterCommonGrams{
		name:        name,
		commonWords: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (g *TokenFilterCommonGrams) Name() string {
	return g.name
}

// CommonWords sets a list of tokens for the filter to generate bigrams.
// Either this or the `common_words_path` parameter is required.
func (g *TokenFilterCommonGrams) CommonWords(commonWords ...string) *TokenFilterCommonGrams {
	g.commonWords = append(g.commonWords, commonWords...)
	return g
}

// CommonWordsPath sets a path to a file containing a list of tokens for the filter to generate bigrams.
// This path must be absolute or relative to the `config` location. The file must be UTF-8 encoded.
// Each token in the file must be separated by a line break.
// Either this or the `common_words` parameter is required.
func (g *TokenFilterCommonGrams) CommonWordsPath(commonWordsPath string) *TokenFilterCommonGrams {
	g.commonWordsPath = commonWordsPath
	return g
}

// IgnoreCase sets whether matches for common words matching should be case-insensitive or not.
// Defaults to false.
func (g *TokenFilterCommonGrams) IgnoreCase(ignoreCase bool) *TokenFilterCommonGrams {
	g.ignoreCase = &ignoreCase
	return g
}

// QueryMode sets whether the filter should exclude the following tokens from the output:
// - Unigrams for common words
// - Unigrams for terms followed by common words
// Recommend enabling `query_mode` for search analyzers.
// Defaults to false.
func (g *TokenFilterCommonGrams) QueryMode(queryMode bool) *TokenFilterCommonGrams {
	g.queryMode = &queryMode
	return g
}

// Validate validates TokenFilterCommonGrams.
func (g *TokenFilterCommonGrams) Validate(includeName bool) error {
	var invalid []string
	if includeName && g.name == "" {
		invalid = append(invalid, "Name")
	}
	if !(len(g.commonWords) > 0) && g.commonWordsPath == "" {
		invalid = append(invalid, "CommonWords || CommonWordsPath")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (g *TokenFilterCommonGrams) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "common_grams",
	// 		"common_words": ["a", "is", "the"],
	// 		"common_words_path": "common_words.txt",
	// 		"ignore_case": true,
	// 		"query_mode": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "common_grams"

	if len(g.commonWords) > 0 {
		options["common_words"] = g.commonWords
	}
	if g.commonWordsPath != "" {
		options["common_words_path"] = g.commonWordsPath
	}
	if g.ignoreCase != nil {
		options["ignore_case"] = g.ignoreCase
	}
	if g.queryMode != nil {
		options["query_mode"] = g.queryMode
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[g.name] = options
	return source, nil
}
