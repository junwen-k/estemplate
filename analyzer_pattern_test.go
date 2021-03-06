// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestAnalyzerPatternSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *AnalyzerPattern
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Stopwords.",
			p:           NewAnalyzerPattern("test").Stopwords("_english_"),
			includeName: true,
			expected:    `{"test":{"stopwords":"_english_","type":"pattern"}}`,
		},
		// #1
		{
			desc:        "Include Name with multiple Stopwords.",
			p:           NewAnalyzerPattern("test").Stopwords("_english_", "_russian_").Stopwords("_french_"),
			includeName: true,
			expected:    `{"test":{"stopwords":["_english_","_russian_","_french_"],"type":"pattern"}}`,
		},
		// #2
		{
			desc:        "Include Name with Pattern, Lowercase and multiple Flags.",
			p:           NewAnalyzerPattern("test").Pattern(`/W+`).Lowercase(true).Flags("CASE_INSENSITIVE", "COMMENTS"),
			includeName: true,
			expected:    `{"test":{"flags":"CASE_INSENSITIVE|COMMENTS","lowercase":true,"pattern":"/W+","type":"pattern"}}`,
		},
		// #3
		{
			desc:        "Exclude Name with StopwordsPath.",
			p:           NewAnalyzerPattern("test").StopwordsPath("stopwords_english.txt"),
			includeName: false,
			expected:    `{"stopwords_path":"stopwords_english.txt","type":"pattern"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.p.Source(test.includeName)
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
