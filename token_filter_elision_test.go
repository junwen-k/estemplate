// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterElisionSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		g           *TokenFilterElision
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Articles and ArticlesCase.",
			g:           NewTokenFilterElision("test").Articles("l", "m").Articles("t").ArticlesCase(true),
			includeName: true,
			expected:    `{"test":{"articles":["l","m","t"],"articles_case":true,"type":"elision"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with ArticlesPath.",
			g:           NewTokenFilterElision("test").ArticlesPath("token_filter/example_elision_list.txt"),
			includeName: false,
			expected:    `{"articles_path":"token_filter/example_elision_list.txt","type":"elision"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.g.Source(test.includeName)
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
