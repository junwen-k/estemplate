// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestSimilarityIBSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *SimilarityIB
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Distribution and Lambda.",
			s:           NewSimilarityIB("test").Distribution("ll").Lambda("df"),
			includeName: true,
			expected:    `{"test":{"distribution":"ll","lambda":"df","type":"IB"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Normalization.",
			s:           NewSimilarityIB("test").Normalization("no"),
			includeName: false,
			expected:    `{"normalization":"no","type":"IB"}`,
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
