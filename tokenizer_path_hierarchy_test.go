// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerPathHierarchySerialization(t *testing.T) {
	tests := []struct {
		desc        string
		h           *TokenizerPathHierarchy
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Delimiter and Replacement.",
			h:           NewTokenizerPathHierarchy("test").Delimiter("-").Replacement("/"),
			includeName: true,
			expected:    `{"test":{"delimiter":"-","replacement":"/","type":"path_hierarchy"}}`,
		},
		// #1
		{
			desc:        "Include Name with BufferSize and Reverse.",
			h:           NewTokenizerPathHierarchy("test").BufferSize(1024).Reverse(true),
			includeName: true,
			expected:    `{"test":{"buffer_size":1024,"reverse":true,"type":"path_hierarchy"}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Skip.",
			h:           NewTokenizerPathHierarchy("test").Skip(2),
			includeName: false,
			expected:    `{"skip":2,"type":"path_hierarchy"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.h.Source(test.includeName)
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
