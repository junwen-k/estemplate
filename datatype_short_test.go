// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeShortSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *DatatypeShort
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			s:           NewDatatypeShort("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"short"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Index and Coerce.",
			s:           NewDatatypeShort("test").Index(true).Coerce(true),
			includeName: false,
			expected:    `{"coerce":true,"index":true,"type":"short"}`,
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