// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestSimilarityScriptedSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *SimilarityScripted
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with WeightScript.",
			s:           NewSimilarityScripted("test").WeightScript(NewScript("double idf = Math.log((field.docCount+1.0)/(term.docFreq+1.0)) + 1.0; return query.boost * idf;")),
			includeName: true,
			expected:    `{"test":{"type":"scripted","weight_script":{"source":"double idf = Math.log((field.docCount+1.0)/(term.docFreq+1.0)) + 1.0; return query.boost * idf;"}}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Script.",
			s:           NewSimilarityScripted("test").Script(NewScript("double tf = Math.sqrt(doc.freq); double norm = 1/Math.sqrt(doc.length); return weight * tf * norm;")),
			includeName: false,
			expected:    `{"script":{"source":"double tf = Math.sqrt(doc.freq); double norm = 1/Math.sqrt(doc.length); return weight * tf * norm;"},"type":"scripted"}`,
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
