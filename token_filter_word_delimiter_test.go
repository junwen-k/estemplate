// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterWordDelimiterSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		d           *TokenFilterWordDelimiter
		includeName bool
		expected    string
	}{
		// #0
		{
			desc: "Include Name with GenerateWordParts, GenerateNumberParts, CatenateWords, CatenateNumbers, CatenateAll, SplitOnCaseChange, PreserveOriginal, SplitOnNumerics and StemEnglishPossessive.",
			d: NewTokenFilterWordDelimiter("test").GenerateWordParts(true).GenerateNumberParts(true).
				CatenateWords(true).CatenateNumbers(true).CatenateAll(true).
				SplitOnCaseChange(true).PreserveOriginal(true).SplitOnNumerics(true).
				StemEnglishPossessive(true),
			includeName: true,
			expected:    `{"test":{"catenate_all":true,"catenate_numbers":true,"catenate_words":true,"generate_number_parts":true,"generate_word_parts":true,"preserve_original":true,"split_on_case_change":true,"split_on_numerics":true,"stem_english_possessive":true,"type":"word_delimiter"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with ProtectedWords and TypeTablePath.",
			d:           NewTokenFilterWordDelimiter("test").ProtectedWords("hello", "world").TypeTablePath("analysis/type_table.txt"),
			includeName: false,
			expected:    `{"protected_words":["hello","world"],"type":"word_delimiter","type_table_path":"analysis/type_table.txt"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
