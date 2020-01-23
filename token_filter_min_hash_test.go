// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestTokenFilterMinHashSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		h           *TokenFilterMinHash
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with HashCount and BucketCount.",
			h:           NewTokenFilterMinHash("test").HashCount(1).BucketCount(512),
			includeName: true,
			expected:    `{"test":{"bucket_count":512,"hash_count":1,"type":"min_hash"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with HashSetSize and WithRotation.",
			h:           NewTokenFilterMinHash("test").HashSetSize(1).WithRotation(true),
			includeName: false,
			expected:    `{"hash_set_size":1,"type":"min_hash","with_rotation":true}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.h.Source(test.includeName)
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
