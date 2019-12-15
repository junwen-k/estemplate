// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeBinarySerialization(t *testing.T) {
	tests := []struct {
		desc        string
		b           *DatatypeBinary
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with DocValues.",
			b:           NewDatatypeBinary("test").DocValues(true),
			includeName: true,
			expected:    `{"test":{"doc_values":true,"type":"binary"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with DocValues and Store.",
			b:           NewDatatypeBinary("test").DocValues(true).Store(true),
			includeName: false,
			expected:    `{"doc_values":true,"store":true,"type":"binary"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.b.Source(test.includeName)
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
