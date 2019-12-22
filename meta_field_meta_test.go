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
		desc     string
		m        *MetaFieldMeta
		expected string
	}{
		// #0
		{
			desc: "With structure Value.",
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
			expected: `{"_meta":{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}}`,
		},
		// #1
		{
			desc: "With map Value.",
			m: NewMetaFieldMeta().Value(map[string]interface{}{
				"class": "MyApp::User",
				"version": map[string]interface{}{
					"max": "1.3",
					"min": "1.0",
				},
			}),
			expected: `{"_meta":{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}}`,
		},
		// #2
		{
			desc: "With RawJSON.",
			m: NewMetaFieldMeta().RawJSON(`
				{
					"class": "MyApp::User",
					"version": {
						"min": "1.0",
						"max": "1.3"
					}
				}
			`),
			expected: `{"_meta":{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}}`,
		},
		// #3
		{
			desc: "With Value and RawJSON, RawJSON should overwrite value.",
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
			expected: `{"_meta":{"class":"MyApp::User","version":{"max":"1.3","min":"1.0"}}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.m.Source()
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
