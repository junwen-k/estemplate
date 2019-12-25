// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// AnalyzerFingerprint is a specialist analyzer which creates a fingerprint
// which can be used for duplicate detection.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-fingerprint-analyzer.html
// for details.
type AnalyzerFingerprint struct {
	Analyzer
	name string

	// fields specific to fingerprint analyzer
	separator     string
	maxOutputSize *int
	stopwords     []string
	stopwordsPath string
}

// NewAnalyzerFingerprint initializes a new AnalyzerFingerprint.
func NewAnalyzerFingerprint(name string) *AnalyzerFingerprint {
	return &AnalyzerFingerprint{
		name: name,
	}
}

// Name returns field key for the Analyzer.
func (f *AnalyzerFingerprint) Name() string {
	return f.name
}

// Separator sets the character to use to concatenate the terms.
// Defaults to " " (space).
func (f *AnalyzerFingerprint) Separator(separator string) *AnalyzerFingerprint {
	f.separator = separator
	return f
}

// MaxOutputSize sets the maximum token size to emit. Tokens larger than this size will
// be discarded.
// Defaults to 255.
func (f *AnalyzerFingerprint) MaxOutputSize(maxOutputSize int) *AnalyzerFingerprint {
	f.maxOutputSize = &maxOutputSize
	return f
}

// Stopwords sets a pre-defined stop words list like "_english_" or an array containing
// a list of stop words.
// Defaults to "_none_".
func (f *AnalyzerFingerprint) Stopwords(stopwords ...string) *AnalyzerFingerprint {
	f.stopwords = append(f.stopwords, stopwords...)
	return f
}

// StopwordsPath sets the path to a file containing stop words.
func (f *AnalyzerFingerprint) StopwordsPath(stopwordsPath string) *AnalyzerFingerprint {
	f.stopwordsPath = stopwordsPath
	return f
}

// Validate validates AnalyzerFingerprint.
func (f *AnalyzerFingerprint) Validate(includeName bool) error {
	var invalid []string
	if includeName && f.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (f *AnalyzerFingerprint) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "fingerprint",
	// 		"separator": "text",
	// 		"max_output_size": 255,
	// 		"stopwords": ["_english_", "_russian_"],
	// 		"stopwords_path": "stopwords_english.txt"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "fingerprint"

	if f.separator != "" {
		options["separator"] = f.separator
	}
	if f.maxOutputSize != nil {
		options["max_output_size"] = f.maxOutputSize
	}
	if len(f.stopwords) > 0 {
		var stopwords interface{}
		switch {
		case len(f.stopwords) > 1:
			stopwords = f.stopwords
			break
		case len(f.stopwords) == 1:
			stopwords = f.stopwords[0]
			break
		default:
			stopwords = ""
		}
		options["stopwords"] = stopwords
	}
	if f.stopwordsPath != "" {
		options["stopwords_path"] = f.stopwordsPath
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[f.name] = options
	return source, nil
}
