// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestSimilarityDFRSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *SimilarityDFR
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with BasicModel and AfterEffect.",
			s:           NewSimilarityDFR("test").BasicModel("g").AfterEffect("b"),
			includeName: true,
			expected:    `{"test":{"after_effect":"b","basic_model":"g","type":"DFR"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Normalization.",
			s:           NewSimilarityDFR("test").Normalization("no"),
			includeName: false,
			expected:    `{"normalization":"no","type":"DFR"}`,
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
