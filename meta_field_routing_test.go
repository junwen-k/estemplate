// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestMetaFieldRoutingSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		r           *MetaFieldRouting
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name without Required.",
			r:           NewMetaFieldRouting(),
			includeName: true,
			expected:    `{"_routing":{}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Required.",
			r:           NewMetaFieldRouting().Required(true),
			includeName: false,
			expected:    `{"required":true}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.r.Source(test.includeName)
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
