// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterNGramSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenFilterNGram
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MaxGram.",
			w:           NewTokenFilterNGram("test").MaxGram(2),
			includeName: true,
			expected:    `{"test":{"max_gram":2,"type":"ngram"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with MinGram.",
			w:           NewTokenFilterNGram("test").MinGram(1),
			includeName: false,
			expected:    `{"min_gram":1,"type":"ngram"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.w.Source(test.includeName)
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
