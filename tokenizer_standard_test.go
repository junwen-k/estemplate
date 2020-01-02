// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerStandardSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *TokenizerStandard
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MaxTokenLength.",
			s:           NewTokenizerStandard("test").MaxTokenLength(255),
			includeName: true,
			expected:    `{"test":{"max_token_length":255,"type":"standard"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			s:           NewTokenizerStandard("test"),
			includeName: false,
			expected:    `{"type":"standard"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
