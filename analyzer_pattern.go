// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"fmt"
	"strings"
)

// AnalyzerPattern uses a regular expression to split the text into terms. The regular
// expression should match the token separators not the tokens themselves. The regular
// expression defaults to `\W+` (or all non-word characters).
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-pattern-analyzer.html
// for details.
type AnalyzerPattern struct {
	Analyzer
	name string

	// fields specific to pattern analyzer
	pattern       string
	flags         []string
	lowercase     *bool
	stopwords     []string
	stopwordsPath string
}

// NewAnalyzerPattern initializes a new AnalyzerPattern.
func NewAnalyzerPattern(name string) *AnalyzerPattern {
	return &AnalyzerPattern{
		name: name,
	}
}

// Name returns field key for the Analyzer.
func (p *AnalyzerPattern) Name() string {
	return p.name
}

// Pattern sets the Java regular expression for the analyzer.
// Defaults to `\W+`.
func (p *AnalyzerPattern) Pattern(pattern string) *AnalyzerPattern {
	p.pattern = pattern
	return p
}

// Flags sets Java regular expression flags. eg "CASE_INSENSITIVE|COMMENTS".
func (p *AnalyzerPattern) Flags(flags ...string) *AnalyzerPattern {
	p.flags = append(p.flags, flags...)
	return p
}

// Lowercase sets whether if the terms be lowercased or not.
// Defaults to true.
func (p *AnalyzerPattern) Lowercase(lowercase bool) *AnalyzerPattern {
	p.lowercase = &lowercase
	return p
}

// Stopwords sets a pre-defined stop words list like "_english_" or an array containing
// a list of stop words.
// Defaults to "_none_".
func (p *AnalyzerPattern) Stopwords(stopwords ...string) *AnalyzerPattern {
	p.stopwords = append(p.stopwords, stopwords...)
	return p
}

// StopwordsPath sets the path to a file containing stop words.
func (p *AnalyzerPattern) StopwordsPath(stopwordsPath string) *AnalyzerPattern {
	p.stopwordsPath = stopwordsPath
	return p
}

// Validate validates AnalyzerPattern.
func (p *AnalyzerPattern) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *AnalyzerPattern) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "pattern",
	// 		"flags": "CASE_INSENSITIVE|COMMENTS",
	// 		"lowercase": true,
	// 		"stopwords": ["_english_", "_russian_"],
	// 		"stopwords_path": "stopwords_english.txt"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "pattern"

	if p.pattern != "" {
		options["pattern"] = p.pattern
	}
	if len(p.flags) > 0 {
		options["flags"] = strings.Join(p.flags, "|")
	}
	if p.lowercase != nil {
		options["lowercase"] = p.lowercase
	}
	if len(p.stopwords) > 0 {
		var stopwords interface{}
		switch {
		case len(p.stopwords) > 1:
			stopwords = p.stopwords
			break
		case len(p.stopwords) == 1:
			stopwords = p.stopwords[0]
			break
		default:
			stopwords = ""
		}
		options["stopwords"] = stopwords
	}
	if p.stopwordsPath != "" {
		options["stopwords_path"] = p.stopwordsPath
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
