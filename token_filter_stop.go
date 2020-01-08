// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterStop token filter that removes stop words from token streams.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-stop-tokenfilter.html
// for details.
type TokenFilterStop struct {
	TokenFilter
	name string

	// fields specific to stop token filter
	stopwords      []string
	stopwordsPath  string
	ignoreCase     *bool
	removeTrailing *bool
}

// NewTokenFilterStop initializes a new TokenFilterStop.
func NewTokenFilterStop(name string) *TokenFilterStop {
	return &TokenFilterStop{
		name:      name,
		stopwords: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (s *TokenFilterStop) Name() string {
	return s.name
}

// Stopwords sets a pre-defined stop words list like "_english_" or an array containing
// a list of stop words.
// Defaults to "_english".
func (s *TokenFilterStop) Stopwords(stopwords ...string) *TokenFilterStop {
	s.stopwords = append(s.stopwords, stopwords...)
	return s
}

// StopwordsPath sets the path to a file containing stop words. This path must be absolute
// or relative to the `config` location. The file must be UTF-8 encoded. Each stopword in the
// file must be separated by a line break.
func (s *TokenFilterStop) StopwordsPath(stopwordsPath string) *TokenFilterStop {
	s.stopwordsPath = stopwordsPath
	return s
}

// IgnoreCase sets whether to lowercase all words first or not.
// Defaults to false.
func (s *TokenFilterStop) IgnoreCase(ignoreCase bool) *TokenFilterStop {
	s.ignoreCase = &ignoreCase
	return s
}

// RemoveTrailing sets whether to ignore the last term of a search if it is a stop word or not.
// This is very useful for the completion suggester as a query like "green a" can be extended to
// "green apple" even though you remove stop words in general.
// Defaults to true.
func (s *TokenFilterStop) RemoveTrailing(removeTrailing bool) *TokenFilterStop {
	s.removeTrailing = &removeTrailing
	return s
}

// Validate validates TokenFilterStop.
func (s *TokenFilterStop) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *TokenFilterStop) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "stop",
	// 		"stopwords": "_english_",
	// 		"stopwords_path": "analysis/example_stopwords.txt",
	// 		"ignore_case": true,
	// 		"remove_trailing": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "stop"

	if len(s.stopwords) > 0 {
		var stopwords interface{}
		switch {
		case len(s.stopwords) > 1:
			stopwords = s.stopwords
			break
		case len(s.stopwords) == 1:
			stopwords = s.stopwords[0]
			break
		default:
			stopwords = ""
		}
		options["stopwords"] = stopwords
	}
	if s.stopwordsPath != "" {
		options["stopwords_path"] = s.stopwordsPath
	}
	if s.ignoreCase != nil {
		options["ignore_case"] = s.ignoreCase
	}
	if s.removeTrailing != nil {
		options["remove_trailing"] = s.removeTrailing
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
