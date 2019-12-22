// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestRelationSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		r           *Relation
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name without children.",
			r:           NewRelation("parent"),
			includeName: true,
			expected:    `{"relations":{"parent":""}}`,
		},
		// #1
		{
			desc:        "Exclude Name with one children.",
			r:           NewRelation("parent", "children_1"),
			includeName: false,
			expected:    `{"parent":"children_1"}`,
		},
		// #2
		{
			desc:        "Exclude Name with multiple childrens.",
			r:           NewRelation("parent", "children_1", "children_2").Children("children_3"),
			includeName: false,
			expected:    `{"parent":["children_1","children_2","children_3"]}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.r.Source(test.includeName)
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
