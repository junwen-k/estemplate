// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestSimilarityLMDirichletSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *SimilarityLMDirichlet
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MU.",
			s:           NewSimilarityLMDirichlet("test").MU(1000),
			includeName: true,
			expected:    `{"test":{"mu":1000,"type":"LMDirichlet"}}`,
		},
		// #2
		{
			desc:        "Exclude Name.",
			s:           NewSimilarityLMDirichlet("test"),
			includeName: false,
			expected:    `{"type":"LMDirichlet"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
