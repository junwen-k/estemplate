// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// AnalyzerStop is like the `simple` analyzer, but also supports removal of stop
// words.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-stop-analyzer.html
// for details.
type AnalyzerStop struct {
	Analyzer
	name string

	// fields specific to stop analyzer
	stopwords     []string
	stopwordsPath string
}

// NewAnalyzerStop initializes a new AnalyzerStop.
func NewAnalyzerStop(name string) *AnalyzerStop {
	return &AnalyzerStop{
		name: name,
	}
}

// Name returns field key for the Analyzer.
func (s *AnalyzerStop) Name() string {
	return s.name
}

// Stopwords sets a pre-defined stop words list like "_english_" or an array containing
// a list of stop words.
// Defaults to "_english_".
func (s *AnalyzerStop) Stopwords(stopwords ...string) *AnalyzerStop {
	s.stopwords = append(s.stopwords, stopwords...)
	return s
}

// StopwordsPath sets the path to a file containing stop words. This path is relative to
// the Elasticsearch `config` directory.
func (s *AnalyzerStop) StopwordsPath(stopwordsPath string) *AnalyzerStop {
	s.stopwordsPath = stopwordsPath
	return s
}

// Validate validates AnalyzerStop.
func (s *AnalyzerStop) Validate(includeName bool) error {
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
func (s *AnalyzerStop) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "stop",
	// 		"stopwords": ["_english_", "_russian_"],
	// 		"stopwords_path": "stopwords_english.txt"
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

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
