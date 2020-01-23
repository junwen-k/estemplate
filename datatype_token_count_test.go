// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeTokenCountSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *DatatypeTokenCount
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Index and Analyzer.",
			c:           NewDatatypeTokenCount("test").Index(true).Analyzer("standard"),
			includeName: true,
			expected:    `{"test":{"analyzer":"standard","index":true,"type":"token_count"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Index.",
			c:           NewDatatypeTokenCount("test").Index(true).EnablePositionIncrements(true),
			includeName: false,
			expected:    `{"enable_position_increments":true,"index":true,"type":"token_count"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.c.Source(test.includeName)
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
