// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestAnalyzerStopSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *AnalyzerStop
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Stopwords.",
			s:           NewAnalyzerStop("test").Stopwords("_english_"),
			includeName: true,
			expected:    `{"test":{"stopwords":"_english_","type":"stop"}}`,
		},
		// #1
		{
			desc:        "Include Name with multiple Stopwords.",
			s:           NewAnalyzerStop("test").Stopwords("_english_", "_russian_").Stopwords("_french_"),
			includeName: true,
			expected:    `{"test":{"stopwords":["_english_","_russian_","_french_"],"type":"stop"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with StopwordsPath.",
			s:           NewAnalyzerStop("test").StopwordsPath("stopwords_english.txt"),
			includeName: false,
			expected:    `{"stopwords_path":"stopwords_english.txt","type":"stop"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.s.Source(test.includeName)
			if err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(src)
			if err != nil {
				t.Fatalf("marshaling to JSON failed: %v", err)
			}
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
