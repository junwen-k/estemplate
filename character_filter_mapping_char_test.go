// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestCharacterFilterMappingCharSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *CharacterFilterMappingChar
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name.",
			c:           NewCharacterFilterMappingChar("test").Mappings(NewMappingRule("٠", "0"), NewMappingRule("١", "1")),
			includeName: true,
			expected:    `{"test":{"mappings":["٠ =\u003e 0","١ =\u003e 1"],"type":"mapping"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			c:           NewCharacterFilterMappingChar("test").MappingsPath("analysis/mappings.txt"),
			includeName: false,
			expected:    `{"mappings_path":"analysis/mappings.txt","type":"mapping"}`,
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
