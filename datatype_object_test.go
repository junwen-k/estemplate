// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeObjectSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		o           *DatatypeObject
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Enabled.",
			o:           NewDatatypeObject("test").Enabled(true),
			includeName: true,
			expected:    `{"test":{"enabled":true,"type":"object"}}`,
		},
		// #1
		{
			desc:        "Include Name with Enabled and Properties.",
			o:           NewDatatypeObject("test").Enabled(true).Properties(NewDatatypeObject("inner").Dynamic(true)),
			includeName: true,
			expected:    `{"test":{"enabled":true,"properties":{"inner":{"dynamic":true,"type":"object"}},"type":"object"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Strict and Dynamic.",
			o:           NewDatatypeObject("test").Strict(true).Dynamic(true),
			includeName: false,
			expected:    `{"dynamic":"strict","type":"object"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.o.Source(test.includeName)
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
