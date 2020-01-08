// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterStopSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *TokenFilterStop
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Stopwords and IgnoreCase.",
			s:           NewTokenFilterStop("test").Stopwords("_none_").IgnoreCase(true),
			includeName: true,
			expected:    `{"test":{"ignore_case":true,"stopwords":"_none_","type":"stop"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			s:           NewTokenFilterStop("test").Stopwords("and", "is").StopwordsPath("analysis/example_stopwords.txt"),
			includeName: false,
			expected:    `{"stopwords":["and","is"],"stopwords_path":"analysis/example_stopwords.txt","type":"stop"}`,
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
