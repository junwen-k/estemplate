// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterSynonymGraphSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *TokenFilterSynonymGraph
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Expand, Lenient and Format.",
			s:           NewTokenFilterSynonymGraph("test").Expand(true).Lenient(true).Format("wordnet"),
			includeName: true,
			expected:    `{"test":{"expand":true,"format":"wordnet","lenient":true,"type":"synonym_graph"}}`,
		},
		// #1
		{
			desc:        "Include Name with Synonyms and SynonymsPath.",
			s:           NewTokenFilterSynonymGraph("test").Synonyms("lol, laughing out loud", "universe, cosmos").SynonymsPath("analysis/synonym.txt"),
			includeName: true,
			expected:    `{"test":{"synonyms":["lol, laughing out loud","universe, cosmos"],"synonyms_path":"analysis/synonym.txt","type":"synonym_graph"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Tokenizer and IgnoreCase.",
			s:           NewTokenFilterSynonymGraph("test").Tokenizer("standard").IgnoreCase(true),
			includeName: false,
			expected:    `{"ignore_case":true,"tokenizer":"standard","type":"synonym_graph"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
