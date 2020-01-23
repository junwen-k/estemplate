// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestScriptSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *Script
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Params.",
			s:           NewScript("doc['my_field'] * multiplier").Params("multiplier", 5),
			includeName: true,
			expected:    `{"script":{"params":{"multiplier":5},"source":"doc['my_field'] * multiplier"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with ID and RawParams.",
			s:           NewScript("").ID("calculate-score").RawParams(map[string]interface{}{"my_modifier": 2}),
			includeName: false,
			expected:    `{"id":"calculate-score","params":{"my_modifier":2}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.s.Source(test.includeName)
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
