// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestMappingRuleSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		r           *MappingRule
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:     "With Key and Value.",
			r:        NewMappingRule("key", "value"),
			expected: `"key =\u003e value"`,
		},
		// #1
		{
			desc:     "With Keys and Values.",
			r:        NewMappingRule("key", "value").Key("second_key").Value("second_value"),
			expected: `"key, second_key =\u003e value, second_value"`,
		},
		// #2
		{
			desc:     "With Key.",
			r:        NewMappingRule("key", ""),
			expected: `"key"`,
		},
		// #3
		{
			desc:     "With Values.",
			r:        NewMappingRule("", "value").Value("second_value"),
			expected: `"value, second_value"`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.r.Source()
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
