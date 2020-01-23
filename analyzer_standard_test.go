// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestAnalyzerStandardSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *AnalyzerStandard
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Stopwords.",
			s:           NewAnalyzerStandard("test").Stopwords("_english_"),
			includeName: true,
			expected:    `{"test":{"stopwords":"_english_","type":"standard"}}`,
		},
		// #1
		{
			desc:        "Include Name with multiple Stopwords.",
			s:           NewAnalyzerStandard("test").Stopwords("_english_", "_russian_").Stopwords("_french_"),
			includeName: true,
			expected:    `{"test":{"stopwords":["_english_","_russian_","_french_"],"type":"standard"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with MaxTokenLength and StopwordsPath.",
			s:           NewAnalyzerStandard("test").MaxTokenLength(255).StopwordsPath("stopwords_english.txt"),
			includeName: false,
			expected:    `{"max_token_length":255,"stopwords_path":"stopwords_english.txt","type":"standard"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
