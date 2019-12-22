// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestIndexPrefixesSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *IndexPrefixes
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MaxChars and MinChars.",
			p:           NewIndexPrefixes(2, 5),
			includeName: true,
			expected:    `{"index_prefixes":{"max_chars":5,"min_chars":2}}`,
		},
		// #1
		{
			desc:        "Exclude Name with modified MinChars.",
			p:           NewIndexPrefixes(2, 5).MinChars(1),
			includeName: false,
			expected:    `{"max_chars":5,"min_chars":1}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.p.Source(test.includeName)
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
