// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterUniqueSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		u           *TokenFilterUnique
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with OnlyOnSamePosition.",
			u:           NewTokenFilterUnique("test").OnlyOnSamePosition(true),
			includeName: true,
			expected:    `{"test":{"only_on_same_position":true,"type":"unique"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			u:           NewTokenFilterUnique("test"),
			includeName: false,
			expected:    `{"type":"unique"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.u.Source(test.includeName)
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
