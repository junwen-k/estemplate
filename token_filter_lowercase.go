// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterLowercase token filter that changes token text to lowercase. For example,
// you can use the `lowercase` filter to change "THE Lazy DoG" to "the lazy dog".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-lowercase-tokenfilter.html
// for details.
type TokenFilterLowercase struct {
	TokenFilter
	name string

	// fields specific to lowercase token filter
	language string
}

// NewTokenFilterLowercase initializes a new TokenFilterLowercase.
func NewTokenFilterLowercase(name string) *TokenFilterLowercase {
	return &TokenFilterLowercase{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (l *TokenFilterLowercase) Name() string {
	return l.name
}

// Language sets a language-specific lowercase token filter to use.
// Can be set to the following values:
// "greek" - Lucene's GreekLowerCaseFilter
// "irish" - Lucene's IrishLowerCaseFilter
// "turkish" - Lucene's TurkishLowerCaseFilter
// Defaults to Lucene's LowerCaseFilter
func (l *TokenFilterLowercase) Language(language string) *TokenFilterLowercase {
	l.language = language
	return l
}

// Validate validates TokenFilterLowercase.
func (l *TokenFilterLowercase) Validate(includeName bool) error {
	var invalid []string
	if includeName && l.name == "" {
		invalid = append(invalid, "Name")
	}
	if l.language != "" {
		if _, valid := map[string]bool{
			"greek":   true,
			"irish":   true,
			"turkish": true,
		}[l.language]; !valid {
			invalid = append(invalid, "Language")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (l *TokenFilterLowercase) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "lowercase",
	// 		"language": "greek"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "lowercase"

	if l.language != "" {
		options["language"] = l.language
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[l.name] = options
	return source, nil
}
