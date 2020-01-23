// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterDictionaryDecompounderSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		d           *TokenFilterDictionaryDecompounder
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with WordList and MaxSubwordSize.",
			d:           NewTokenFilterDictionaryDecompounder("test").WordList("Donau", "dampf").WordList("meer", "schiff").MaxSubwordSize(15),
			includeName: true,
			expected:    `{"test":{"max_subword_size":15,"type":"dictionary_decompounder","word_list":["Donau","dampf","meer","schiff"]}}`,
		},
		// #1
		{
			desc:        "Exclude Name with WordListPath and OnlyLongestMatch.",
			d:           NewTokenFilterDictionaryDecompounder("test").WordListPath("analysis/example_word_list.txt").OnlyLongestMatch(true),
			includeName: false,
			expected:    `{"only_longest_match":true,"type":"dictionary_decompounder","word_list_path":"analysis/example_word_list.txt"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.d.Source(test.includeName)
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
