// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterKeepWordsSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenFilterKeepWords
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with KeepWords and KeepWordsCase.",
			w:           NewTokenFilterKeepWords("test").KeepWords("one", "two").KeepWordsCase(true),
			includeName: true,
			expected:    `{"test":{"keep_words":["one","two"],"keep_words_case":true,"type":"keep"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with KeepWordsPath.",
			w:           NewTokenFilterKeepWords("test").KeepWordsPath("analysis/example_word_list.txt"),
			includeName: false,
			expected:    `{"keep_words_path":"analysis/example_word_list.txt","type":"keep"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.w.Source(test.includeName)
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
