// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterPhoneticDoubleMetaphoneSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		p           *TokenFilterPhoneticDoubleMetaphone
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Replace.",
			p:           NewTokenFilterPhoneticDoubleMetaphone("test").Replace(true),
			includeName: true,
			expected:    `{"test":{"encoder":"double_metaphone","replace":true,"type":"phonetic"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with MaxCodeLen.",
			p:           NewTokenFilterPhoneticDoubleMetaphone("test").MaxCodeLen(4),
			includeName: false,
			expected:    `{"encoder":"double_metaphone","max_code_len":4,"type":"phonetic"}`,
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
