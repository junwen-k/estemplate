// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterCJKBigramSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		b           *TokenFilterCJKBigram
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with multiple IgnoredScripts.",
			b:           NewTokenFilterCJKBigram("test").IgnoredScripts("han", "hangul").IgnoredScripts("hiragana", "katakana"),
			includeName: true,
			expected:    `{"test":{"ignored_scripts":["han","hangul","hiragana","katakana"],"type":"cjk_bigram"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with OutputUnigrams.",
			b:           NewTokenFilterCJKBigram("test").OutputUnigrams(true),
			includeName: false,
			expected:    `{"output_unigrams":true,"type":"cjk_bigram"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.b.Source(test.includeName)
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
