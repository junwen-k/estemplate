// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterKeepWords token filter that keeps only tokens contained in a specified
// word list.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-keep-words-tokenfilter.html
// for details.
type TokenFilterKeepWords struct {
	TokenFilter
	name string

	// fields specific to keep words token filter
	keepWords     []string
	keepWordsPath string
	keepWordsCase *bool
}

// NewTokenFilterKeepWords initializes a new TokenFilterKeepWords.
func NewTokenFilterKeepWords(name string) *TokenFilterKeepWords {
	return &TokenFilterKeepWords{
		name:      name,
		keepWords: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (w *TokenFilterKeepWords) Name() string {
	return w.name
}

// KeepWords sets a list of words to keep. Only tokens that match words in this list
// are included in the output.
// Either this or `keep_words_path` parameter must be specified.
func (w *TokenFilterKeepWords) KeepWords(keepWords ...string) *TokenFilterKeepWords {
	w.keepWords = append(w.keepWords, keepWords...)
	return w
}

// KeepWordsPath sets a path to a file containing a list of words to keep. Only tokens
// that match words in this list are included in the output. This path must be absolute
// or relative to the `config` location, and the file must be UTF-8 encoded. Each word
// in the file must be separated by a line break.
// Either this or `keep_words` parameter must be specified.
func (w *TokenFilterKeepWords) KeepWordsPath(keepWordsPath string) *TokenFilterKeepWords {
	w.keepWordsPath = keepWordsPath
	return w
}

// KeepWordsCase sets whether to lowercase all keep words.
// Defaults to false.
func (w *TokenFilterKeepWords) KeepWordsCase(keepWordsCase bool) *TokenFilterKeepWords {
	w.keepWordsCase = &keepWordsCase
	return w
}

// Validate validates TokenFilterKeepWords.
func (w *TokenFilterKeepWords) Validate(includeName bool) error {
	var invalid []string
	if includeName && w.name == "" {
		invalid = append(invalid, "Name")
	}
	if !(len(w.keepWords) > 0) && w.keepWordsPath == "" {
		invalid = append(invalid, "KeepWords || KeepWordsPath")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (w *TokenFilterKeepWords) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "keep",
	// 		"keep_words": ["one", "two", "three"],
	// 		"keep_words_path": "analysis/example_word_list.txt",
	// 		"keep_words_case": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "keep"

	if len(w.keepWords) > 0 {
		options["keep_words"] = w.keepWords
	}
	if w.keepWordsPath != "" {
		options["keep_words_path"] = w.keepWordsPath
	}
	if w.keepWordsCase != nil {
		options["keep_words_case"] = w.keepWordsCase
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[w.name] = options
	return source, nil
}
