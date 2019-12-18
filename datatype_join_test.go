// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeJoinSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		j           *DatatypeJoin
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with EagerGlobalOrdinals.",
			j:           NewDatatypeJoin("test").EagerGlobalOrdinals(true),
			includeName: true,
			expected:    `{"test":{"eager_global_ordinals":true,"type":"join"}}`,
		},
		// #1
		{
			desc:        "Include Name with Relations.",
			j:           NewDatatypeJoin("test").Relations(NewRelation("parent", "children")),
			includeName: true,
			expected:    `{"test":{"relations":{"parent":"children"},"type":"join"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Multiple Relations.",
			j:           NewDatatypeJoin("test").Relations(NewRelation("parent_1", "children_1", "children_2"), NewRelation("children_1").Children("children_3")),
			includeName: false,
			expected:    `{"relations":{"children_1":"children_3","parent_1":["children_1","children_2"]},"type":"join"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.j.Source(test.includeName)
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
