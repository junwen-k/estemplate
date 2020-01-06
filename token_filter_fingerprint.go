// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterFingerprint token filter that sorts and removes duplicate tokens from a
// token stream, then concatenates the stream into a single output token.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-fingerprint-tokenfilter.html
// for details.
type TokenFilterFingerprint struct {
	TokenFilter
	name string

	// fields specific to fingerprint token filter
	maxOutputSize *int
	separator     string
}

// NewTokenFilterFingerprint initializes a new TokenFilterFingerprint.
func NewTokenFilterFingerprint(name string) *TokenFilterFingerprint {
	return &TokenFilterFingerprint{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (f *TokenFilterFingerprint) Name() string {
	return f.name
}

// MaxOutputSize sets the maximum character length, including whitespace, of the output token.
// Concatenated tokens longer than 255 will result in no token output.
// Defaults to 255.
func (f *TokenFilterFingerprint) MaxOutputSize(maxOutputSize int) *TokenFilterFingerprint {
	f.maxOutputSize = &maxOutputSize
	return f
}

// Separator sets a character to use to concatenate the token stream input.
// Defaults to " " (space).
func (f *TokenFilterFingerprint) Separator(separator string) *TokenFilterFingerprint {
	f.separator = separator
	return f
}

// Validate validates TokenFilterFingerprint.
func (f *TokenFilterFingerprint) Validate(includeName bool) error {
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
func (f *TokenFilterFingerprint) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "fingerprint",
	// 		"max_output_size": 100,
	// 		"separator": "+"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "fingerprint"

	if f.maxOutputSize != nil {
		options["max_output_size"] = f.maxOutputSize
	}
	if f.separator != "" {
		options["separator"] = f.separator
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[f.name] = options
	return source, nil
}
