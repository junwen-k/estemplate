// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterPatternReplaceSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *TokenFilterPatternReplace
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Patterns.",
			c:           NewTokenFilterPatternReplace("test").Pattern(`\`),
			includeName: true,
			expected:    `{"test":{"pattern":"\\","type":"pattern_replace"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with PreserveOriginal.",
			c:           NewTokenFilterPatternReplace("test").Replacement("+"),
			includeName: false,
			expected:    `{"replacement":"+","type":"pattern_replace"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.c.Source(test.includeName)
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
