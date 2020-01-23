// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestAnalysisSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		a           *Analysis
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with DefaultAnalyzer and Analyzers.",
			a:           NewAnalysis().DefaultAnalyzer(NewAnalyzerWhitespace("")).Analyzer(NewAnalyzerCustom("custom_keyword", "standard")),
			includeName: true,
			expected:    `{"analysis":{"analyzer":{"custom_keyword":{"tokenizer":"standard","type":"custom"},"default":{"type":"whitespace"}}}}`,
		},
		// #1
		{
			desc:        "Include Name with Normalizers.",
			a:           NewAnalysis().Normalizer(NewNormalizerCustom("sort_normalizer").Filter("lowercase", "asciifolding")),
			includeName: true,
			expected:    `{"analysis":{"normalizer":{"sort_normalizer":{"filter":["lowercase","asciifolding"],"type":"custom"}}}}`,
		},
		// #2
		{
			desc:        "Include Name with Token Filters.",
			a:           NewAnalysis().Filter(NewTokenFilterEdgeNGram("autocomplete_filter").MinGram(1).MaxGram(20)),
			includeName: true,
			expected:    `{"analysis":{"filter":{"autocomplete_filter":{"max_gram":20,"min_gram":1,"type":"edge_ngram"}}}}`,
		},
		// #3
		{
			desc:        "Include Name with Character Filters.",
			a:           NewAnalysis().CharFilter(NewCharacterFilterMappingChar("custom_mapping").Mappings(NewMappingRule("٠", "0"), NewMappingRule("١", "1"), NewMappingRule("٢", "2"))),
			includeName: true,
			expected:    `{"analysis":{"char_filter":{"custom_mapping":{"mappings":["٠ =\u003e 0","١ =\u003e 1","٢ =\u003e 2"],"type":"mapping"}}}}`,
		},
		// #4
		{
			desc:        "Exclude Name.",
			a:           NewAnalysis(),
			includeName: false,
			expected:    `{}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.a.Source(test.includeName)
			if err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(src)
			if err != nil {
				t.Fatalf("marshaling to JSON failed: %v", err)
			}
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
