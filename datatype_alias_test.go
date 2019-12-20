// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeAliasSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		a           *DatatypeAlias
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Path.",
			a:           NewDatatypeAlias("test").Path("field1"),
			includeName: true,
			expected:    `{"test":{"path":"field1","type":"alias"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Path.",
			a:           NewDatatypeAlias("test").Path("object1.field1"),
			includeName: false,
			expected:    `{"path":"object1.field1","type":"alias"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
