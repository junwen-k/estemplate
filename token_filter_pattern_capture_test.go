// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterPatternCaptureSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *TokenFilterPatternCapture
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Patterns.",
			c:           NewTokenFilterPatternCapture("test").Patterns("(\\p{Ll}+|\\p{Lu}\\p{Ll}+|\\p{Lu}+)", "(\\d+)"),
			includeName: true,
			expected:    `{"test":{"patterns":["(\\p{Ll}+|\\p{Lu}\\p{Ll}+|\\p{Lu}+)","(\\d+)"],"type":"pattern_capture"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with PreserveOriginal.",
			c:           NewTokenFilterPatternCapture("test").PreserveOriginal(true),
			includeName: false,
			expected:    `{"preserve_original":true,"type":"pattern_capture"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
