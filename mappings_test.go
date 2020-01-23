// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestMappingsSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		m           *Mappings
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with DynamicTemplates.",
			m:           NewMappings().DynamicTemplates(NewDynamicTemplate("template_1").MatchMappingType("string").Mapping(NewDatatypeText("").Analyzer("{name}"))),
			includeName: true,
			expected:    `{"mappings":{"dynamic_templates":[{"template_1":{"mapping":{"analyzer":"{name}","type":"text"},"match_mapping_type":"string"}}]}}`,
		},
		// #1
		{
			desc:        "Include Name with DateDetection and DynamicDateFormats.",
			m:           NewMappings().DateDetection(true).DynamicDateFormats(NewDateFormat("date_optional_time").Strict(true)),
			includeName: true,
			expected:    `{"mappings":{"date_detection":true,"dynamic_date_formats":["strict_date_optional_time"]}}`,
		},
		// #2
		{
			desc:        "Include Name with MetaSource.",
			m:           NewMappings().MetaSource(NewMetaFieldSource().Enabled(false)),
			includeName: true,
			expected:    `{"mappings":{"_source":{"enabled":false}}}`,
		},
		// #3
		{
			desc:        "Exclude Name with Properties.",
			m:           NewMappings().Properties(NewDatatypeText("field_1").Analyzer("standard"), NewDatatypeKeyword("field_2").Store(false)),
			includeName: false,
			expected:    `{"properties":{"field_1":{"analyzer":"standard","type":"text"},"field_2":{"store":false,"type":"keyword"}}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.m.Source(test.includeName)
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
