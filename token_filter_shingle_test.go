// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterShingleSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *TokenFilterShingle
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MaxShingleSize and MinShingleSize.",
			p:           NewTokenFilterShingle("test").MaxShingleSize(2).MinShingleSize(2).TokenSeparator("/"),
			includeName: true,
			expected:    `{"test":{"max_shingle_size":2,"min_shingle_size":2,"token_separator":"/","type":"shingle"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			p:           NewTokenFilterShingle("test").OutputUnigrams(true).OutputUnigramsIfNoShingles(true).FilterToken("_"),
			includeName: false,
			expected:    `{"filter_token":"_","output_unigrams":true,"output_unigrams_if_no_shingles":true,"type":"shingle"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.p.Source(test.includeName)
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
