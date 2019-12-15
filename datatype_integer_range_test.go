// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeIntegerRangeSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		r           *DatatypeIntegerRange
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			r:           NewDatatypeIntegerRange("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"integer_range"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Index and Coerce.",
			r:           NewDatatypeIntegerRange("test").Index(true).Coerce(true),
			includeName: false,
			expected:    `{"coerce":true,"index":true,"type":"integer_range"}`,
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
