// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestMetaFieldSourceSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *MetaFieldSource
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Includes.",
			s:           NewMetaFieldSource().Includes("*.count", "meta.*"),
			includeName: true,
			expected:    `{"_source":{"includes":["*.count","meta.*"]}}`,
		},
		// #1
		{
			desc:        "Include Name with Excludes.",
			includeName: true,
			s:           NewMetaFieldSource().Excludes("meta.description").Excludes("meta.other.*"),
			expected:    `{"_source":{"excludes":["meta.description","meta.other.*"]}}`,
		},
		// #2
		{
			desc:        "Exclude Name with Enabled.",
			s:           NewMetaFieldSource().Enabled(false),
			includeName: false,
			expected:    `{"enabled":false}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.s.Source(test.includeName)
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
