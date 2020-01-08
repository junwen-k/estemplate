// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterPhoneticSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *TokenFilterPhonetic
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Encoder.",
			p:           NewTokenFilterPhonetic("test").Encoder("metaphone"),
			includeName: true,
			expected:    `{"test":{"encoder":"metaphone","type":"phonetic"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Replace.",
			p:           NewTokenFilterPhonetic("test").Replace(true),
			includeName: false,
			expected:    `{"replace":true,"type":"phonetic"}`,
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
