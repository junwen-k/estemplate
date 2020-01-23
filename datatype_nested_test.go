// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeNestedSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		n           *DatatypeNested
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Dynamic.",
			n:           NewDatatypeNested("test").Dynamic(true),
			includeName: true,
			expected:    `{"test":{"dynamic":true,"type":"nested"}}`,
		},
		// #1
		{
			desc:        "Include Name with Dynamic and Properties.",
			n:           NewDatatypeNested("test").Dynamic(true).Properties(NewDatatypeNested("inner").Dynamic(true)),
			includeName: true,
			expected:    `{"test":{"dynamic":true,"properties":{"inner":{"dynamic":true,"type":"nested"}},"type":"nested"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Strict and Dynamic.",
			n:           NewDatatypeNested("test").Strict(true).Dynamic(true),
			includeName: false,
			expected:    `{"dynamic":"strict","type":"nested"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.n.Source(test.includeName)
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
