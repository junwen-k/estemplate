// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeSearchAsYouTypeSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		t           *DatatypeSearchAsYouType
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			t:           NewDatatypeSearchAsYouType("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"search_as_you_type"}}`,
		},
		// #1
		{
			desc:        "Include Name with Index and MaxShingleSize.",
			t:           NewDatatypeSearchAsYouType("test").Index(true).MaxShingleSize(3),
			includeName: true,
			expected:    `{"test":{"index":true,"max_shingle_size":3,"type":"search_as_you_type"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Index and Analyzer.",
			t:           NewDatatypeSearchAsYouType("test").Index(true).Analyzer("standard"),
			includeName: false,
			expected:    `{"analyzer":"standard","index":true,"type":"search_as_you_type"}`,
		},
		// #3
		{
			desc:        "Exclude Name with Index, Analyzer andn MaxShingleSize.",
			t:           NewDatatypeSearchAsYouType("test").Index(true).Analyzer("standard").MaxShingleSize(3),
			includeName: false,
			expected:    `{"analyzer":"standard","index":true,"max_shingle_size":3,"type":"search_as_you_type"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.t.Source(test.includeName)
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
