// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestDatatypeDateSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		d           *DatatypeDate
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Format.",
			d:           NewDatatypeDate("test").Format(NewDateFormat("date_optional_time").Strict(true)).Format(NewDateFormat("epoch_millis")),
			includeName: true,
			expected:    `{"test":{"format":"strict_date_optional_time||epoch_millis","type":"date"}}`,
		},
		// #1
		{
			desc:        "Exclude Name with Format, RawFormat and IgnoreMalformed.",
			d:           NewDatatypeDate("test").Format(NewDateFormat("epoch_millis")).RawFormat("strict_date_optional_time||epoch_millis").IgnoreMalformed(true),
			includeName: false,
			expected:    `{"format":"strict_date_optional_time||epoch_millis","ignore_malformed":true,"type":"date"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.d.Source(test.includeName)
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
