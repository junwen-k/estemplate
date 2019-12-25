// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// AnalyzerKeyword is a "noop" analyzer that accepts whatever text it is
// given and outputs the exact same text as a single term.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-keyword-analyzer.html
// for details.
type AnalyzerKeyword struct {
	Analyzer
	name string

	// fields specific to keyword analyzer
}

// NewAnalyzerKeyword initializes a new AnalyzerKeyword.
func NewAnalyzerKeyword(name string) *AnalyzerKeyword {
	return &AnalyzerKeyword{
		name: name,
	}
}

// Name returns field key for the Analyzer.
func (k *AnalyzerKeyword) Name() string {
	return k.name
}

// Validate validates AnalyzerKeyword.
func (k *AnalyzerKeyword) Validate(includeName bool) error {
	var invalid []string
	if includeName && k.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (k *AnalyzerKeyword) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "keyword"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "keyword"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[k.name] = options
	return source, nil
}
