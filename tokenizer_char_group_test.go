// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerCharGroupSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		g           *TokenizerCharGroup
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with multiple TokenizeOnChars.",
			g:           NewTokenizerCharGroup("test").TokenizeOnChars("whitespace", "-").TokenizeOnChars("\n"),
			includeName: true,
			expected:    `{"test":{"tokenize_on_chars":["whitespace","-","\n"],"type":"char_group"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			g:           NewTokenizerCharGroup("test"),
			includeName: false,
			expected:    `{"type":"char_group"}`,
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
