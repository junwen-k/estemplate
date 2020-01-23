// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterConditionalSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *TokenFilterConditional
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with multiple Filter.",
			c:           NewTokenFilterConditional("test").Filter("lowercase", "asciifolding"),
			includeName: true,
			expected:    `{"test":{"filter":["lowercase","asciifolding"],"type":"condition"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Script.",
			c:           NewTokenFilterConditional("test").Script(NewScript(`token.getTerm().length() < threshold`).Params("threshold", 5)),
			includeName: false,
			expected:    `{"script":{"params":{"threshold":5},"source":"token.getTerm().length() \u003c threshold"},"type":"condition"}`,
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
