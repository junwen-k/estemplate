// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterSnowball token filter that stems words using a Snowball-generated stemmer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-snowball-tokenfilter.html
// for details.
type TokenFilterSnowball struct {
	TokenFilter
	name string

	// fields specific to snowball token filter
	language string
}

// NewTokenFilterSnowball initializes a new TokenFilterSnowball.
func NewTokenFilterSnowball(name string) *TokenFilterSnowball {
	return &TokenFilterSnowball{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (s *TokenFilterSnowball) Name() string {
	return s.name
}

// Language sets language which controls the stemmer.
// Can be set to the following values:
// "Armenian"
// "Basque"
// "Catalan"
// "Danish"
// "Dutch"
// "English"
// "Finnish"
// "French"
// "German"
// "German2"
// "Hungarian"
// "Italian"
// "Kp"
// "Lithuanian"
// "Lovins"
// "Norwegian"
// "Porter"
// "Portuguese"
// "Romanian"
// "Russian"
// "Spanish"
// "Swedish"
// "Turkish"
func (s *TokenFilterSnowball) Language(language string) *TokenFilterSnowball {
	s.language = language
	return s
}

// Validate validates TokenFilterSnowball.
func (s *TokenFilterSnowball) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *TokenFilterSnowball) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "snowball",
	// 		"language": "Armenian"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "snowball"

	if s.language != "" {
		options["language"] = s.language
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
