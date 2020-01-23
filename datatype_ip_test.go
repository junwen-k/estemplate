// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeIPSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		ip          *DatatypeIP
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			ip:          NewDatatypeIP("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"ip"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Index and DocValues.",
			ip:          NewDatatypeIP("test").Index(true).DocValues(true),
			includeName: false,
			expected:    `{"doc_values":true,"index":true,"type":"ip"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.ip.Source(test.includeName)
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
