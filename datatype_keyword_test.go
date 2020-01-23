// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeKeywordSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		k           *DatatypeKeyword
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			k:           NewDatatypeKeyword("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"keyword"}}`,
		},
		// #1
		{
			desc:        "Include Name with Index and Multi-level Fields.",
			k:           NewDatatypeKeyword("test").Index(true).Fields(NewDatatypeKeyword("keyword")),
			includeName: true,
			expected:    `{"test":{"fields":{"keyword":{"type":"keyword"}},"index":true,"type":"keyword"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Index and Normalizer.",
			k:           NewDatatypeKeyword("test").Index(true).Normalizer("my_normalizer"),
			includeName: false,
			expected:    `{"index":true,"normalizer":"my_normalizer","type":"keyword"}`,
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
