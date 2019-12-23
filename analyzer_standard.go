// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// AnalyzerStandard is the default analyzer which is used if none is specified. It provides
// grammer based tokenization (based on Unicode Text Segmentation algorithm, as specified
// in Unicode Standard Annex #29) and works well for most languages.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-standard-analyzer.html
// for details.
type AnalyzerStandard struct {
	name string

	// fields specific to standard analyzer
	maxTokenLength *int
	stopwords      []string
	stopwordsPath  string
}

// NewAnalyzerStandard initializes a new AnalyzerStandard.
func NewAnalyzerStandard(name string) *AnalyzerStandard {
	return &AnalyzerStandard{
		name: name,
	}
}

// Name returns field key for the Analyzer.
func (s *AnalyzerStandard) Name() string {
	return s.name
}

// MaxTokenLength sets the maximum token length and if a token is seen that exceeds this
// length then it is split at `max_token_length` intervals.
// Defaults to 255.
func (s *AnalyzerStandard) MaxTokenLength(maxTokenLength int) *AnalyzerStandard {
	s.maxTokenLength = &maxTokenLength
	return s
}

// Stopwords sets a pre-defined stop words list like "_english_" or an array containing
// a list of stop words.
// Defaults to "_none_".
func (s *AnalyzerStandard) Stopwords(stopwords ...string) *AnalyzerStandard {
	s.stopwords = append(s.stopwords, stopwords...)
	return s
}

// StopwordsPath sets the path to a file containing stop words.
func (s *AnalyzerStandard) StopwordsPath(stopwordsPath string) *AnalyzerStandard {
	s.stopwordsPath = stopwordsPath
	return s
}

// Validate validates AnalyzerStandard.
func (s *AnalyzerStandard) Validate(includeName bool) error {
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
func (s *AnalyzerStandard) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "standard",
	// 		"max_token_length": 255,
	// 		"stopwords": ["_english_", "_russian_"],
	// 		"stopwords_path": "stopwords_english.txt"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "standard"

	if s.maxTokenLength != nil {
		options["max_token_length"] = s.maxTokenLength
	}
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

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
