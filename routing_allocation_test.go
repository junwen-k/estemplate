// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestRoutingAllocationSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		a           *RoutingAllocation
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with single value.",
			a:           NewRoutingAllocation("include", "_name", "node_name"),
			includeName: true,
			expected:    `{"routing":{"allocation.include._name":"node_name"}}`,
		},
		// #1
		{
			desc:        "Include Name with multiple values.",
			a:           NewRoutingAllocation("require", "_id", "id_1", "id_2"),
			includeName: true,
			expected:    `{"routing":{"allocation.require._id":"id_1,id_2"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with exclude.",
			a:           NewRoutingAllocation("exclude", "custom_attribute", "value_1", "value_2", "value_3"),
			includeName: false,
			expected:    `{"allocation.exclude.custom_attribute":"value_1,value_2,value_3"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.a.Source(test.includeName)
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
