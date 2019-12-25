// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// AnalyzerWhitespace breaks text into terms whenever it encounters a whitespace
// character. It does not lowercase terms.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-whitespace-analyzer.html
// for details.
type AnalyzerWhitespace struct {
	Analyzer
	name string

	// fields specific to whitespace analyzer
}

// NewAnalyzerWhitespace initializes a new AnalyzerWhitespace.
func NewAnalyzerWhitespace(name string) *AnalyzerWhitespace {
	return &AnalyzerWhitespace{
		name: name,
	}
}

// Name returns field key for the Analyzer.
func (w *AnalyzerWhitespace) Name() string {
	return w.name
}

// Validate validates AnalyzerWhitespace.
func (w *AnalyzerWhitespace) Validate(includeName bool) error {
	var invalid []string
	if includeName && w.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (w *AnalyzerWhitespace) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "whitespace"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "whitespace"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[w.name] = options
	return source, nil
}
