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
		desc        string
		s           *MetaFieldSize
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name without Enabled.",
			s:           NewMetaFieldSize(),
			includeName: true,
			expected:    `{"_size":{}}`,
		},
		// #1
		{
			desc:        "Exclude Name With Enabled.",
			s:           NewMetaFieldSize().Enabled(true),
			includeName: false,
			expected:    `{"enabled":true}`,
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
