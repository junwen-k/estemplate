// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterPhonetic (Plugin) token filter that provides token filters which
// convert tokens to their phonetic representation using Soundex, Metaphone, and
// a variety of other algorithms.
//
// See https://www.elastic.co/guide/en/elasticsearch/plugins/7.5/analysis-phonetic.html
// for details.
type TokenFilterPhonetic struct {
	TokenFilter
	name string

	// fields specific to phonetic token filter
	encoder string
	replace *bool
}

// NewTokenFilterPhonetic initializes a new TokenFilterPhonetic.
func NewTokenFilterPhonetic(name string) *TokenFilterPhonetic {
	return &TokenFilterPhonetic{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (p *TokenFilterPhonetic) Name() string {
	return p.name
}

// Encoder sets the phonetic encoder to use.
// Can be set to the following values:
// "metaphone"
// "double_metaphone"
// "soundex"
// "refined_soundex"
// "caverphone1"
// "caverphone2"
// "cologne"
// "nysiis"
// "koelnerphonetik"
// "haasephonetik"
// "beider_morse"
// "daitch_mokotoff"
// For "double_metaphone" and "beider_morse" encoder, use NewTokenFilterPhoneticDoubleMetaphone
// and NewTokenFilterPhoneticBeiderMorse respectively.
// Defaults to "metaphone".
func (p *TokenFilterPhonetic) Encoder(encoder string) *TokenFilterPhonetic {
	p.encoder = encoder
	return p
}

// Replace sets whether or not the original token should be replaced by
// the phonetic token.
// Defaults to true.
func (p *TokenFilterPhonetic) Replace(replace bool) *TokenFilterPhonetic {
	p.replace = &replace
	return p
}

// Validate validates TokenFilterPhonetic.
func (p *TokenFilterPhonetic) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if p.encoder != "" {
		if _, valid := map[string]bool{
			"metaphone":        true,
			"double_metaphone": true,
			"soundex":          true,
			"refined_soundex":  true,
			"caverphone1":      true,
			"caverphone2":      true,
			"cologne":          true,
			"nysiis":           true,
			"koelnerphonetik":  true,
			"haasephonetik":    true,
			"beider_morse":     true,
			"daitch_mokotoff":  true,
		}[p.encoder]; !valid {
			invalid = append(invalid, "Encoder")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *TokenFilterPhonetic) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "phonetic",
	// 		"encoder": "metaphone",
	// 		"replace": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "phonetic"

	if p.encoder != "" {
		options["encoder"] = p.encoder
	}
	if p.replace != nil {
		options["replace"] = p.replace
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
