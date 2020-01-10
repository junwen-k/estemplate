// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestCharacterFilterPatternReplaceCharSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *CharacterFilterPatternReplaceChar
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name.",
			c:           NewCharacterFilterPatternReplaceChar("test").Pattern(`(\\d+)-(?=\\d)`).Replacement("$1_"),
			includeName: true,
			expected:    `{"test":{"pattern":"(\\\\d+)-(?=\\\\d)","replacement":"$1_","type":"pattern_replace"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			c:           NewCharacterFilterPatternReplaceChar("test").Flags("CASE_INSENSITIVE", "COMMENTS"),
			includeName: false,
			expected:    `{"flags":"CASE_INSENSITIVE|COMMENTS","type":"pattern_replace"}`,
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
