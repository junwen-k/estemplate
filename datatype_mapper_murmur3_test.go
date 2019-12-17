// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeMapperMurmur3Serialization(t *testing.T) {
	tests := []struct {
		desc        string
		m3          *DatatypeMapperMurmur3
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name.",
			m3:          NewDatatypeMapperMurmur3("test"),
			includeName: true,
			expected:    `{"test":{"type":"murmur3"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			m3:          NewDatatypeMapperMurmur3("test"),
			includeName: false,
			expected:    `{"type":"murmur3"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.m3.Source(test.includeName)
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
