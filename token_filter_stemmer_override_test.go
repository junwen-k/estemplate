// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterStemmerOverrideSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *TokenFilterStemmerOverride
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Rules.",
			s:           NewTokenFilterStemmerOverride("test").Rules(NewTokenFilterStemmerMappingRule("running", "run"), NewTokenFilterStemmerMappingRule("stemmer", "stemmer")),
			includeName: true,
			expected:    `{"test":{"rules":["running =\u003e run","stemmer =\u003e stemmer"],"type":"stemmer_override"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with RulesPath.",
			s:           NewTokenFilterStemmerOverride("test").RulesPath("analysis/stemmer_override.txt"),
			includeName: false,
			expected:    `{"rules_path":"analysis/stemmer_override.txt","type":"stemmer_override"}`,
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