// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerSimplePatternSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *TokenizerSimplePattern
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Pattern.",
			p:           NewTokenizerSimplePattern("test").Pattern(`[0123456789]{3}`),
			includeName: true,
			expected:    `{"test":{"pattern":"[0123456789]{3}","type":"simple_pattern"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			p:           NewTokenizerSimplePattern("test"),
			includeName: false,
			expected:    `{"type":"simple_pattern"}`,
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
