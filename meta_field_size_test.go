// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestMetaFieldSizeSerialization(t *testing.T) {
	tests := []struct {
		desc     string
		s        *MetaFieldSize
		expected string
	}{
		// #0
		{
			desc:     "Without Enabled.",
			s:        NewMetaFieldSize(),
			expected: `{"_size":{}}`,
		},
		// #1
		{
			desc:     "With Enabled.",
			s:        NewMetaFieldSize().Enabled(true),
			expected: `{"_size":{"enabled":true}}`,
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
