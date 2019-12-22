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
		desc     string
		s        *MetaFieldSource
		expected string
	}{
		// #0
		{
			desc:     "With Includes.",
			s:        NewMetaFieldSource().Includes("*.count", "meta.*"),
			expected: `{"_source":{"includes":["*.count","meta.*"]}}`,
		},
		// #1
		{
			desc:     "With Excludes.",
			s:        NewMetaFieldSource().Excludes("meta.description").Excludes("meta.other.*"),
			expected: `{"_source":{"excludes":["meta.description","meta.other.*"]}}`,
		},
		// #2
		{
			desc:     "With Enabled.",
			s:        NewMetaFieldSource().Enabled(false),
			expected: `{"_source":{"enabled":false}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.s.Source()
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
