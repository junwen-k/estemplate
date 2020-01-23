// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeCompletionSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *DatatypeCompletion
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Analyzer.",
			c:           NewDatatypeCompletion("test").Analyzer("simple"),
			includeName: true,
			expected:    `{"test":{"analyzer":"simple","type":"completion"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with PreservePositionIncrements and MaxInputLength.",
			c:           NewDatatypeCompletion("test").PreservePositionIncrements(true).MaxInputLength(100),
			includeName: false,
			expected:    `{"max_input_length":100,"preserve_position_increments":true,"type":"completion"}`,
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
