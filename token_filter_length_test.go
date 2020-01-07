// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterLengthSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenFilterLength
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Min.",
			w:           NewTokenFilterLength("test").Min(2),
			includeName: true,
			expected:    `{"test":{"min":2,"type":"length"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Max.",
			w:           NewTokenFilterLength("test").Max(10),
			includeName: false,
			expected:    `{"max":10,"type":"length"}`,
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
