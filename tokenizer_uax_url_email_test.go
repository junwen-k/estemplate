// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenizerUAXURLEmailSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenizerUAXURLEmail
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with MaxTokenLength.",
			w:           NewTokenizerUAXURLEmail("test").MaxTokenLength(255),
			includeName: true,
			expected:    `{"test":{"max_token_length":255,"type":"uax_url_email"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			w:           NewTokenizerUAXURLEmail("test"),
			includeName: false,
			expected:    `{"type":"uax_url_email"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
