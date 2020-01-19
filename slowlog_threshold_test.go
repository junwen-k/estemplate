// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestSlowlogThresholdSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		t           *SlowlogThreshold
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Value.",
			t:           NewSlowlogThreshold("search", "query", "warn", "10s").Value("5s"),
			includeName: true,
			expected:    `{"search":{"slowlog.threshold.query.warn":"5s"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			t:           NewSlowlogThreshold("indexing", "index", "info", "5s"),
			includeName: false,
			expected:    `{"slowlog.threshold.index.info":"5s"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.t.Source(test.includeName)
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
