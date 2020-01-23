// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterFingerprintSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		g           *TokenFilterFingerprint
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Articles and ArticlesCase.",
			g:           NewTokenFilterFingerprint("test").MaxOutputSize(100),
			includeName: true,
			expected:    `{"test":{"max_output_size":100,"type":"fingerprint"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Separator.",
			g:           NewTokenFilterFingerprint("test").Separator("+"),
			includeName: false,
			expected:    `{"separator":"+","type":"fingerprint"}`,
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
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
