// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerPatternSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *TokenizerPattern
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Pattern.",
			p:           NewTokenizerPattern("test").Pattern(`/W+`),
			includeName: true,
			expected:    `{"test":{"pattern":"/W+","type":"pattern"}}`,
		},
		// #1
		{
			desc:        "Include Name with multiple Flags.",
			p:           NewTokenizerPattern("test").Flags("CASE_INSENSITIVE", "COMMENTS"),
			includeName: true,
			expected:    `{"test":{"flags":"CASE_INSENSITIVE|COMMENTS","type":"pattern"}}`,
		},
		// #2
		{
			desc:        "Include Name with Group.",
			p:           NewTokenizerPattern("test").Group(1),
			includeName: true,
			expected:    `{"test":{"group":1,"type":"pattern"}}`,
		},
		// #3
		{
			desc:        "Exclude Name.",
			p:           NewTokenizerPattern("test"),
			includeName: false,
			expected:    `{"type":"pattern"}`,
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
