// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeBooleanSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		b           *DatatypeBoolean
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			b:           NewDatatypeBoolean("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"boolean"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Index and Store.",
			b:           NewDatatypeBoolean("test").Index(true).Store(true),
			includeName: false,
			expected:    `{"index":true,"store":true,"type":"boolean"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
