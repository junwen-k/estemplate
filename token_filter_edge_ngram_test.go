// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterEdgeNGramSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		g           *TokenFilterEdgeNGram
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MaxGram and MinGram.",
			g:           NewTokenFilterEdgeNGram("test").MaxGram(1).MinGram(1),
			includeName: true,
			expected:    `{"test":{"max_gram":1,"min_gram":1,"type":"edge_ngram"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Side.",
			g:           NewTokenFilterEdgeNGram("test").Side("front"),
			includeName: false,
			expected:    `{"side":"front","type":"edge_ngram"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
