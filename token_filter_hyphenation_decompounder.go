// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterHyphenationDecompounder token filter that uses XML-based hyphenation patterns
// to find potential subwords in compound words. These subwords are then checked against the
// specified word list. Subwords not in the list are excluded from the token output.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-hyp-decomp-tokenfilter.html
// for details.
type TokenFilterHyphenationDecompounder struct {
	TokenFilter
	name string

	// fields specific to hyphenation decompounder token filter
	hyphenationPatternsPath string
	wordList                []string
	wordListPath            string
	maxSubwordSize          *int
	minSubwordSize          *int
	minWordSize             *int
	onlyLongestMatch        *bool
}

// NewTokenFilterHyphenationDecompounder initializes a new TokenFilterHyphenationDecompounder.
func NewTokenFilterHyphenationDecompounder(name string) *TokenFilterHyphenationDecompounder {
	return &TokenFilterHyphenationDecompounder{
		name:     name,
		wordList: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (d *TokenFilterHyphenationDecompounder) Name() string {
	return d.name
}

// HyphenationPatternsPath sets a path to an Apache FOP (Formatting Objects Processor) XML
// hyphenation pattern file. This path must be absolute or relative to the `config` location.
// Only FOP v1.2 compatible files are supported.
func (d *TokenFilterHyphenationDecompounder) HyphenationPatternsPath(hyphenationPatternsPath string) *TokenFilterHyphenationDecompounder {
	d.hyphenationPatternsPath = hyphenationPatternsPath
	return d
}

// WordList sets a list of subwords. Subwords found using the hyphenation pattern but not in this
// list are excluded from the token output.
// Either this or `word_list_path` parameter must be specified.
func (d *TokenFilterHyphenationDecompounder) WordList(wordList ...string) *TokenFilterHyphenationDecompounder {
	d.wordList = append(d.wordList, wordList...)
	return d
}

// WordListPath sets a path to a file containing a list of subwords. Subwords found using the
// hyphenation pattern but not in this list are excluded from the token output.
// Either this or `word_list` parameter must be specified.
func (d *TokenFilterHyphenationDecompounder) WordListPath(wordListPath string) *TokenFilterHyphenationDecompounder {
	d.wordListPath = wordListPath
	return d
}

// MaxSubwordSize sets the maximum subword character length. Longer subword tokens are excluded from the output.
// Defaults to 15.
func (d *TokenFilterHyphenationDecompounder) MaxSubwordSize(maxSubwordSize int) *TokenFilterHyphenationDecompounder {
	d.maxSubwordSize = &maxSubwordSize
	return d
}

// MinSubwordSize sets the minimum subword character length. Shorter subword tokens are excluded from the output.
// Defaults to 2.
func (d *TokenFilterHyphenationDecompounder) MinSubwordSize(minSubwordSize int) *TokenFilterHyphenationDecompounder {
	d.minSubwordSize = &minSubwordSize
	return d
}

// MinWordSize sets the minimum word character length. Shorter word tokens are excluded from the output.
// Defaults to 5.
func (d *TokenFilterHyphenationDecompounder) MinWordSize(minWordSize int) *TokenFilterHyphenationDecompounder {
	d.minWordSize = &minWordSize
	return d
}

// OnlyLongestMatch sets whether to only include the longest matching subword or not.
// Defaults to false.
func (d *TokenFilterHyphenationDecompounder) OnlyLongestMatch(onlyLongestMatch bool) *TokenFilterHyphenationDecompounder {
	d.onlyLongestMatch = &onlyLongestMatch
	return d
}

// Validate validates TokenFilterHyphenationDecompounder.
func (d *TokenFilterHyphenationDecompounder) Validate(includeName bool) error {
	var invalid []string
	if includeName && d.name == "" {
		invalid = append(invalid, "Name")
	}
	if d.hyphenationPatternsPath == "" {
		invalid = append(invalid, "HyphenationPatternsPath")
	}
	if !(len(d.wordList) > 0) && d.wordListPath == "" {
		invalid = append(invalid, "WordList || WordListPath")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (d *TokenFilterHyphenationDecompounder) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "hyphenation_decompounder",
	// 		"hyphenation_patterns_path": "analysis/hyphenation_patterns.xml",
	// 		"word_list": ["Kaffee", "zucker", "tasse"],
	// 		"word_list_path": "analysis/example_word_list.txt",
	// 		"max_subword_size": 15,
	// 		"min_subword_size": 2,
	// 		"min_word_size": 5,
	// 		"only_longest_match": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "hyphenation_decompounder"

	if d.hyphenationPatternsPath != "" {
		options["hyphenation_patterns_path"] = d.hyphenationPatternsPath
	}
	if len(d.wordList) > 0 {
		options["word_list"] = d.wordList
	}
	if d.wordListPath != "" {
		options["word_list_path"] = d.wordListPath
	}
	if d.maxSubwordSize != nil {
		options["max_subword_size"] = d.maxSubwordSize
	}
	if d.minSubwordSize != nil {
		options["min_subword_size"] = d.minSubwordSize
	}
	if d.minWordSize != nil {
		options["min_word_size"] = d.minWordSize
	}
	if d.onlyLongestMatch != nil {
		options["only_longest_match"] = d.onlyLongestMatch
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[d.name] = options
	return source, nil
}
