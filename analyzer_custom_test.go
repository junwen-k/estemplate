// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestAnalyzerCustomSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *AnalyzerCustom
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Tokenizer and CharFilter.",
			c:           NewAnalyzerCustom("test", "standard").CharFilter("html_strip"),
			includeName: true,
			expected:    `{"test":{"char_filter":["html_strip"],"tokenizer":"standard","type":"custom"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Tokenizer and Filter and PositionIncrementGap.",
			c:           NewAnalyzerCustom("test", "keyword").Filter("lowercase", "asciifolding").PositionIncrementGap(1),
			includeName: false,
			expected:    `{"filter":["lowercase","asciifolding"],"position_increment_gap":1,"tokenizer":"keyword","type":"custom"}`,
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
