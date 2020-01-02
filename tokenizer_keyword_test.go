// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerKeywordSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		k           *TokenizerKeyword
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name.",
			k:           NewTokenizerKeyword("test"),
			includeName: true,
			expected:    `{"test":{"type":"keyword"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with BufferSize.",
			k:           NewTokenizerKeyword("test").BufferSize(256),
			includeName: false,
			expected:    `{"buffer_size":256,"type":"keyword"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.k.Source(test.includeName)
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
