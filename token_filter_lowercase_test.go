// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterLowercaseSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenFilterLowercase
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Language.",
			w:           NewTokenFilterLowercase("test").Language("greek"),
			includeName: true,
			expected:    `{"test":{"language":"greek","type":"lowercase"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			w:           NewTokenFilterLowercase("test"),
			includeName: false,
			expected:    `{"type":"lowercase"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.w.Source(test.includeName)
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
