// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerNGramSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		n           *TokenizerNGram
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MinGram and MaxGram.",
			n:           NewTokenizerNGram("test").MinGram(3).MaxGram(3),
			includeName: true,
			expected:    `{"test":{"max_gram":3,"min_gram":3,"type":"ngram"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with TokenChars.",
			n:           NewTokenizerNGram("test").TokenChars("letter").TokenChars("digit", "whitespace"),
			includeName: false,
			expected:    `{"token_chars":["letter","digit","whitespace"],"type":"ngram"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.n.Source(test.includeName)
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
