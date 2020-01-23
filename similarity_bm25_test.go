// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestSimilarityBM25Serialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *SimilarityBM25
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with K1.",
			s:           NewSimilarityBM25("test").K1(1.2),
			includeName: true,
			expected:    `{"test":{"k1":1.2,"type":"BM25"}}`,
		},
		// #1
		{
			desc:        "Include Name with B.",
			s:           NewSimilarityBM25("test").B(0.75),
			includeName: true,
			expected:    `{"test":{"b":0.75,"type":"BM25"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with DiscountOverlaps.",
			s:           NewSimilarityBM25("test").DiscountOverlaps(true),
			includeName: false,
			expected:    `{"discount_overlaps":true,"type":"BM25"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
