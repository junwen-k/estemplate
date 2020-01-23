// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeGeoShapeSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		s           *DatatypeGeoShape
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with IgnoreZValue.",
			s:           NewDatatypeGeoShape("test").IgnoreZValue(true),
			includeName: true,
			expected:    `{"test":{"ignore_z_value":true,"type":"geo_shape"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Coerce.",
			s:           NewDatatypeGeoShape("test").Coerce(true),
			includeName: false,
			expected:    `{"coerce":true,"type":"geo_shape"}`,
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
