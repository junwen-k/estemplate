// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// AnalyzerSimple divdes text into terms whenever it encounters any whitespace
// character. It lowercases all terms.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-simple-analyzer.html
// for details.
type AnalyzerSimple struct {
	Analyzer
	name string

	// fields specific to simple analyzer
}

// NewAnalyzerSimple initializes a new AnalyzerSimple.
func NewAnalyzerSimple(name string) *AnalyzerSimple {
	return &AnalyzerSimple{
		name: name,
	}
}

// Name returns field key for the Analyzer.
func (s *AnalyzerSimple) Name() string {
	return s.name
}

// Validate validates AnalyzerSimple.
func (s *AnalyzerSimple) Validate(includeName bool) error {
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
func (s *AnalyzerSimple) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "simple"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "simple"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
