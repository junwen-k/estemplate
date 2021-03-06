// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeTextSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		t           *DatatypeText
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index.",
			t:           NewDatatypeText("test").Index(true),
			includeName: true,
			expected:    `{"test":{"index":true,"type":"text"}}`,
		},
		// #1
		{
			desc:        "Include Name with Index and Multi-level Fields.",
			t:           NewDatatypeText("test").Index(true).Fields(NewDatatypeText("english").Analyzer("english")),
			includeName: true,
			expected:    `{"test":{"fields":{"english":{"analyzer":"english","type":"text"}},"index":true,"type":"text"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Index and Analyzer.",
			t:           NewDatatypeText("test").Index(true).Analyzer("standard"),
			includeName: false,
			expected:    `{"analyzer":"standard","index":true,"type":"text"}`,
		},
		// #3
		{
			desc:        "Exclude Name with Index, Analyzer and FielddataFrequencyFilter.",
			t:           NewDatatypeText("test").Index(true).Analyzer("standard").FielddataFrequencyFilter(NewFielddataFrequencyFilter(0.001, 0.1)),
			includeName: false,
			expected:    `{"analyzer":"standard","fielddata_frequency_filter":{"max":0.1,"min":0.001},"index":true,"type":"text"}`,
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
