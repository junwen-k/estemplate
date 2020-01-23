// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDateFormatSerialization(t *testing.T) {
	tests := []struct {
		desc     string
		f        *DateFormat
		expected string
	}{
		// #0
		{
			desc:     "Without Strict.",
			f:        NewDateFormat("format"),
			expected: `"format"`,
		},
		// #1
		{
			desc:     "With Strict.",
			f:        NewDateFormat("format").Strict(true),
			expected: `"strict_format"`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.f.Source()
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
