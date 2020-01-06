// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterDictionaryDecompounder token filter that uses a specified list of words
// and a brute force approach to find subwords in compound words. If found, these subwords
// are included in the token output.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-dict-decomp-tokenfilter.html
// for details.
type TokenFilterDictionaryDecompounder struct {
	TokenFilter
	name string

	// fields specific to dictionary decompounder token filter
	wordList         []string
	wordListPath     string
	maxSubwordSize   *int
	minSubwordSize   *int
	minWordSize      *int
	onlyLongestMatch *bool
}

// NewTokenFilterDictionaryDecompounder initializes a new TokenFilterDictionaryDecompounder.
func NewTokenFilterDictionaryDecompounder(name string) *TokenFilterDictionaryDecompounder {
	return &TokenFilterDictionaryDecompounder{
		name:     name,
		wordList: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (d *TokenFilterDictionaryDecompounder) Name() string {
	return d.name
}

// WordList sets a list of subwords to look for in the token stream. If found, the subword is included
// in the token output.
// Either this or `word_list_path` parameter must be specified.
func (d *TokenFilterDictionaryDecompounder) WordList(wordList ...string) *TokenFilterDictionaryDecompounder {
	d.wordList = append(d.wordList, wordList...)
	return d
}

// WordListPath sets a path to a file containing a list of subwords to find in the token stream. If found,
// the subword is included in the token output. This path must be absolute or relative to the `config`
// location. The file must be UTF-8 encoded. Each token in the file must be separated by a line break.
// Either this or `word_list` parameter must be specified.
func (d *TokenFilterDictionaryDecompounder) WordListPath(wordListPath string) *TokenFilterDictionaryDecompounder {
	d.wordListPath = wordListPath
	return d
}

// MaxSubwordSize sets the maximum subword character length. Longer subword tokens are excluded from the output.
// Defaults to 15.
func (d *TokenFilterDictionaryDecompounder) MaxSubwordSize(maxSubwordSize int) *TokenFilterDictionaryDecompounder {
	d.maxSubwordSize = &maxSubwordSize
	return d
}

// MinSubwordSize sets the minimum subword character length. Shorter subword tokens are excluded from the output.
// Defaults to 2.
func (d *TokenFilterDictionaryDecompounder) MinSubwordSize(minSubwordSize int) *TokenFilterDictionaryDecompounder {
	d.minSubwordSize = &minSubwordSize
	return d
}

// MinWordSize sets the minimum word character length. Shorter word tokens are excluded from the output.
// Defaults to 5.
func (d *TokenFilterDictionaryDecompounder) MinWordSize(minWordSize int) *TokenFilterDictionaryDecompounder {
	d.minWordSize = &minWordSize
	return d
}

// OnlyLongestMatch sets whether to only include the longest matching subword or not.
// Defaults to false.
func (d *TokenFilterDictionaryDecompounder) OnlyLongestMatch(onlyLongestMatch bool) *TokenFilterDictionaryDecompounder {
	d.onlyLongestMatch = &onlyLongestMatch
	return d
}

// Validate validates TokenFilterDictionaryDecompounder.
func (d *TokenFilterDictionaryDecompounder) Validate(includeName bool) error {
	var invalid []string
	if includeName && d.name == "" {
		invalid = append(invalid, "Name")
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
func (d *TokenFilterDictionaryDecompounder) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "dictionary_decompounder",
	// 		"word_list": ["Donau", "dampf", "meer", "schiff"],
	// 		"word_list_path": "analysis/example_word_list.txt",
	// 		"max_subword_size": 15,
	// 		"min_subword_size": 2,
	// 		"min_word_size": 5,
	// 		"only_longest_match": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "dictionary_decompounder"

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
