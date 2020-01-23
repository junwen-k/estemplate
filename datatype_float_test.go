// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeFloatSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		f           *DatatypeFloat
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			f:           NewDatatypeFloat("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"float"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Index and Coerce.",
			f:           NewDatatypeFloat("test").Index(true).Coerce(true),
			includeName: false,
			expected:    `{"coerce":true,"index":true,"type":"float"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
