// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestMetaFieldMetaSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		m           *MetaFieldMeta
		includeName bool
		expected    string
	}{
		// #0
		{
			desc: "Include Name with structure Value.",
			m: NewMetaFieldMeta().Value(&struct {
				Class   string `json:"class"`
				Version struct {
					Min string `json:"min"`
					Max string `json:"max"`
				} `json:"version"`
			}{
				Class: "MyApp::User",
				Version: struct {
					Min string `json:"min"`
					Max string `json:"max"`
				}{
					Min: "1.0",
					Max: "1.3",
				},
			}),
			includeName: true,
			expected:    `{"_meta":{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}}`,
		},
		// #1
		{
			desc: "Include Name with map Value.",
			m: NewMetaFieldMeta().Value(map[string]interface{}{
				"class": "MyApp::User",
				"version": map[string]interface{}{
					"max": "1.3",
					"min": "1.0",
				},
			}),
			includeName: true,
			expected:    `{"_meta":{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}}`,
		},
		// #2
		{
			desc: "Include Name with RawJSON.",
			m: NewMetaFieldMeta().RawJSON(`
				{
					"class": "MyApp::User",
					"version": {
						"min": "1.0",
						"max": "1.3"
					}
				}
			`),
			includeName: true,
			expected:    `{"_meta":{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}}`,
		},
		// #3
		{
			desc: "Exclude Name with Value and RawJSON, RawJSON should overwrite value.",
			m: NewMetaFieldMeta().Value(map[string]interface{}{
				"class": "MyApp::User",
				"version": map[string]interface{}{
					"max": "1.3",
					"min": "1.0",
				},
				"type": "value",
			}).RawJSON(`
				{
					"class": "MyApp::User",
					"version": {
						"min": "1.0",
						"max": "1.3"
					}
				}
			`),
			includeName: false,
			expected:    `{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}`,
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
