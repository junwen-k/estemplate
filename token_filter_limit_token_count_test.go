// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterLimitTokenCountSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenFilterLimitTokenCount
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MaxTokenCount.",
			w:           NewTokenFilterLimitTokenCount("test").MaxTokenCount(5),
			includeName: true,
			expected:    `{"test":{"max_token_count":5,"type":"limit"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with ConsumeAllTokens.",
			w:           NewTokenFilterLimitTokenCount("test").ConsumeAllTokens(true),
			includeName: false,
			expected:    `{"consume_all_tokens":true,"type":"limit"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
