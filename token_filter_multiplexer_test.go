// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterMultiplexerSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		w           *TokenFilterMultiplexer
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Filters.",
			w:           NewTokenFilterMultiplexer("test").Filters("lowercase", "lowercase, porter_stem"),
			includeName: true,
			expected:    `{"test":{"filters":["lowercase","lowercase, porter_stem"],"type":"multiplexer"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with PreserveOriginal.",
			w:           NewTokenFilterMultiplexer("test").PreserveOriginal(true),
			includeName: false,
			expected:    `{"preserve_original":true,"type":"multiplexer"}`,
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
