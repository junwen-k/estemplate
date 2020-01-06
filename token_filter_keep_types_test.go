// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterKeepTypesSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		d           *TokenFilterKeepTypes
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with multiple Types.",
			d:           NewTokenFilterKeepTypes("test").Types("<NUM>", "<ALPHANUM>"),
			includeName: true,
			expected:    `{"test":{"type":"keep_types","types":["\u003cNUM\u003e","\u003cALPHANUM\u003e"]}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Mode.",
			d:           NewTokenFilterKeepTypes("test").Mode("exclude"),
			includeName: false,
			expected:    `{"mode":"exclude","type":"keep_types"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.d.Source(test.includeName)
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
