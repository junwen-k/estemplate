// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestFielddataFrequencyFilterSerialization(t *testing.T) {
	tests := []struct {
		desc     string
		f        *FielddataFrequencyFilter
		expected string
	}{
		// #0
		{
			desc:     "Without MinSegmentSize.",
			f:        NewFielddataFrequencyFilter(0.001, 0.1),
			expected: `{"max":0.1,"min":0.001}`,
		},
		// #1
		{
			desc:     "With MinSegmentSize.",
			f:        NewFielddataFrequencyFilter(0.001, 0.1).MinSegmentSize(1),
			expected: `{"max":0.1,"min":0.001,"min_segment_size":1}`,
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
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
