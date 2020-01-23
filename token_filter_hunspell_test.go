// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterHunspellSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		h           *TokenFilterHunspell
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Locale and Dictionary.",
			h:           NewTokenFilterHunspell("test").Locale("en_US").Dictionary("US_dictionary"),
			includeName: true,
			expected:    `{"test":{"dictionary":"US_dictionary","locale":"en_US","type":"hunspell"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Dedup and LongestOnly.",
			h:           NewTokenFilterHunspell("test").Dedup(true).LongestOnly(true),
			includeName: false,
			expected:    `{"dedup":true,"longest_only":true,"type":"hunspell"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.h.Source(test.includeName)
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
