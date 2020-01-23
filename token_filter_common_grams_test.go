// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterCommonGramsSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		g           *TokenFilterCommonGrams
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with CommonWords and IgnoreCase.",
			g:           NewTokenFilterCommonGrams("test").CommonWords("a", "is").CommonWords("the").IgnoreCase(true),
			includeName: true,
			expected:    `{"test":{"common_words":["a","is","the"],"ignore_case":true,"type":"common_grams"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with CommonWordsPath and QueryMode.",
			g:           NewTokenFilterCommonGrams("test").CommonWordsPath("common_words.txt").QueryMode(true),
			includeName: false,
			expected:    `{"common_words_path":"common_words.txt","query_mode":true,"type":"common_grams"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.g.Source(test.includeName)
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
