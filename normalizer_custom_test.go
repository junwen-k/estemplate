// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestNormalizerCustomSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		c           *NormalizerCustom
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with CharFilter.",
			c:           NewNormalizerCustom("test").CharFilter("html_strip"),
			includeName: true,
			expected:    `{"test":{"char_filter":["html_strip"],"type":"custom"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Filter.",
			includeName: false,
			c:           NewNormalizerCustom("test").Filter("lowercase", "asciifolding"),
			expected:    `{"filter":["lowercase","asciifolding"],"type":"custom"}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
