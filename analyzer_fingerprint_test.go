// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestAnalyzerFingerprintSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		f           *AnalyzerFingerprint
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with Stopwords.",
			f:           NewAnalyzerFingerprint("test").Stopwords("_english_"),
			includeName: true,
			expected:    `{"test":{"stopwords":"_english_","type":"fingerprint"}}`,
		},
		// #1
		{
			desc:        "Include Name with multiple Stopwords.",
			f:           NewAnalyzerFingerprint("test").Stopwords("_english_", "_russian_").Stopwords("_french_"),
			includeName: true,
			expected:    `{"test":{"stopwords":["_english_","_russian_","_french_"],"type":"fingerprint"}}`,
		},
		// #2
		{
			desc:        "Include Name with Separator and MaxOutputSize.",
			f:           NewAnalyzerFingerprint("test").Separator("test").MaxOutputSize(255),
			includeName: true,
			expected:    `{"test":{"max_output_size":255,"separator":"test","type":"fingerprint"}}`,
		},
		// #3
		{
			desc:        "Exclude Name with StopwordsPath.",
			f:           NewAnalyzerFingerprint("test").StopwordsPath("stopwords_english.txt"),
			includeName: false,
			expected:    `{"stopwords_path":"stopwords_english.txt","type":"fingerprint"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.f.Source(test.includeName)
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
