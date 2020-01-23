// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerSimplePatternSplitSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *TokenizerSimplePatternSplit
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Pattern.",
			s:           NewTokenizerSimplePatternSplit("test").Pattern(`_`),
			includeName: true,
			expected:    `{"test":{"pattern":"_","type":"simple_pattern_split"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			s:           NewTokenizerSimplePatternSplit("test"),
			includeName: false,
			expected:    `{"type":"simple_pattern_split"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.s.Source(test.includeName)
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
