// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterPhoneticBeiderMorseSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *TokenFilterPhoneticBeiderMorse
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with RuleType.",
			p:           NewTokenFilterPhoneticBeiderMorse("test").RuleType("exact").NameType("ashkenazi"),
			includeName: true,
			expected:    `{"test":{"encoder":"beider_morse","name_type":"ashkenazi","rule_type":"exact","type":"phonetic"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Languageset.",
			p:           NewTokenFilterPhoneticBeiderMorse("test").Languageset("any", "common"),
			includeName: false,
			expected:    `{"encoder":"beider_morse","languageset":["any","common"],"type":"phonetic"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
