// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterPredicateScriptSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *TokenFilterPredicateScript
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Script.",
			p:           NewTokenFilterPredicateScript("test").Script(NewScript("token.getTerm().length() > 5")),
			includeName: true,
			expected:    `{"test":{"script":{"source":"token.getTerm().length() \u003e 5"},"type":"predicate_token_filter"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			p:           NewTokenFilterPredicateScript("test"),
			includeName: false,
			expected:    `{"type":"predicate_token_filter"}`,
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
