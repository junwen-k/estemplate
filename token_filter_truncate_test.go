// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterTruncateSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		t           *TokenFilterTruncate
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Limit.",
			t:           NewTokenFilterTruncate("test").Limit(10),
			includeName: true,
			expected:    `{"test":{"limit":10,"type":"truncate"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			t:           NewTokenFilterTruncate("test"),
			includeName: false,
			expected:    `{"type":"truncate"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.t.Source(test.includeName)
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
