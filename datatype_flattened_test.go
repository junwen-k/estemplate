// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeFlattenedSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		f           *DatatypeFlattened
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			f:           NewDatatypeFlattened("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"flattened"}}`,
		},
		// #1
		{
			desc:        "Include Name with Index and DepthLimit.",
			f:           NewDatatypeFlattened("test").Index(true).DepthLimit(20),
			includeName: true,
			expected:    `{"test":{"depth_limit":20,"index":true,"type":"flattened"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Index and SplitQueriesOnWhitespace.",
			f:           NewDatatypeFlattened("test").Index(true).SplitQueriesOnWhitespace(true),
			includeName: false,
			expected:    `{"index":true,"split_queries_on_whitespace":true,"type":"flattened"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.f.Source(test.includeName)
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
