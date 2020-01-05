// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterASCIIFolding token filter that converts alphabetic, numeric and symboling characters
// that are not in the Basic Latin Unicode block (first 127 ASCII characters) to their
// ASCII equivalent, if one exists. For example, the filter changes "à" to "a".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-asciifolding-tokenfilter.html
// for details.
type TokenFilterASCIIFolding struct {
	TokenFilter
	name string

	// fields specific to ascii folding token filter
	preserveOriginal *bool
}

// NewTokenFilterASCIIFolding initializes a new TokenFilterASCIIFolding.
func NewTokenFilterASCIIFolding(name string) *TokenFilterASCIIFolding {
	return &TokenFilterASCIIFolding{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (f *TokenFilterASCIIFolding) Name() string {
	return f.name
}

// PreserveOriginal sets whether to emit both original tokens and folded tokens.
// Defaults to false.
func (f *TokenFilterASCIIFolding) PreserveOriginal(preserveOriginal bool) *TokenFilterASCIIFolding {
	f.preserveOriginal = &preserveOriginal
	return f
}

// Validate validates TokenFilterASCIIFolding.
func (f *TokenFilterASCIIFolding) Validate(includeName bool) error {
	var invalid []string
	if includeName && f.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (f *TokenFilterASCIIFolding) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "asciifolding",
	// 		"preserve_original": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "asciifolding"

	if f.preserveOriginal != nil {
		options["preserve_original"] = f.preserveOriginal
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[f.name] = options
	return source, nil
}
