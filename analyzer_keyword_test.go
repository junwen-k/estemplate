// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestAnalyzerKeywordSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		k           *AnalyzerKeyword
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name.",
			k:           NewAnalyzerKeyword("test"),
			includeName: true,
			expected:    `{"test":{"type":"keyword"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			k:           NewAnalyzerKeyword("test"),
			includeName: false,
			expected:    `{"type":"keyword"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
