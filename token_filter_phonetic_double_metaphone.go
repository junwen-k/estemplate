// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterPhoneticDoubleMetaphone (Plugin) token filter that provides token filters which
// convert tokens to their phonetic representation using Soundex, Metaphone, and
// a variety of other algorithms.
//
// See https://www.elastic.co/guide/en/elasticsearch/plugins/7.5/analysis-phonetic-token-filter.html
// for details.
type TokenFilterPhoneticDoubleMetaphone struct {
	TokenFilter
	name string

	// fields specific to phonetic token beider morse filter
	replace    *bool
	maxCodeLen *int
}

// NewTokenFilterPhoneticDoubleMetaphone initializes a new TokenFilterPhoneticDoubleMetaphone.
func NewTokenFilterPhoneticDoubleMetaphone(name string) *TokenFilterPhoneticDoubleMetaphone {
	return &TokenFilterPhoneticDoubleMetaphone{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (p *TokenFilterPhoneticDoubleMetaphone) Name() string {
	return p.name
}

// Replace sets whether or not the original token should be replaced by
// the phonetic token.
// Defaults to true.
func (p *TokenFilterPhoneticDoubleMetaphone) Replace(replace bool) *TokenFilterPhoneticDoubleMetaphone {
	p.replace = &replace
	return p
}

// MaxCodeLen sets the maximum length of the emitted metaphone token.
// Defaults to 4.
func (p *TokenFilterPhoneticDoubleMetaphone) MaxCodeLen(maxCodeLen int) *TokenFilterPhoneticDoubleMetaphone {
	p.maxCodeLen = &maxCodeLen
	return p
}

// Validate validates TokenFilterPhoneticDoubleMetaphone.
func (p *TokenFilterPhoneticDoubleMetaphone) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *TokenFilterPhoneticDoubleMetaphone) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "phonetic",
	// 		"encoder": "double_metaphone",
	// 		"replace": true,
	// 		"max_code_len": 4
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "phonetic"
	options["encoder"] = "double_metaphone"

	if p.replace != nil {
		options["replace"] = p.replace
	}
	if p.maxCodeLen != nil {
		options["max_code_len"] = p.maxCodeLen
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
