// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeDenseVectorSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		v           *DatatypeDenseVector
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Dims.",
			v:           NewDatatypeDenseVector("test").Dims(3),
			includeName: true,
			expected:    `{"test":{"dims":3,"type":"dense_vector"}}`,
		},
		// #1
		{
			desc:        "Exclude Name.",
			v:           NewDatatypeDenseVector("test"),
			includeName: false,
			expected:    `{"type":"dense_vector"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.v.Source(test.includeName)
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
