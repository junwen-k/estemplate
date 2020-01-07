// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterKeywordMarkerSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenFilterKeywordMarker
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Keywords and IgnoreCase.",
			w:           NewTokenFilterKeywordMarker("test").Keywords("one", "two").IgnoreCase(true),
			includeName: true,
			expected:    `{"test":{"ignore_case":true,"keywords":["one","two"],"type":"keyword_marker"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with KeywordsPath and KeywordsPattern.",
			w:           NewTokenFilterKeywordMarker("test").KeywordsPath("analysis/example_word_list.txt").KeywordsPattern(`^profit_\d+$`),
			includeName: false,
			expected:    `{"keywords_path":"analysis/example_word_list.txt","keywords_pattern":"^profit_\\d+$","type":"keyword_marker"}`,
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
