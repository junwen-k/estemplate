// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterStemmerMappingRuleSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		r           *TokenFilterStemmerMappingRule
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:     "Without From and To.",
			r:        NewTokenFilterStemmerMappingRule("from", "to"),
			expected: `"from =\u003e to"`,
		},
		// #1
		{
			desc:     "With From and To.",
			r:        NewTokenFilterStemmerMappingRule("from", "to").From("updated_from").To("updated_to"),
			expected: `"updated_from =\u003e updated_to"`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.r.Source()
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
