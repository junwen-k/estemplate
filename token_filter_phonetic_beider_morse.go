// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterPhoneticBeiderMorse (Plugin) token filter that provides token filters which
// convert tokens to their phonetic representation using Soundex, Metaphone, and
// a variety of other algorithms.
//
// See https://www.elastic.co/guide/en/elasticsearch/plugins/7.5/analysis-phonetic-token-filter.html
// for details.
type TokenFilterPhoneticBeiderMorse struct {
	TokenFilter
	name string

	// fields specific to phonetic token beider morse filter
	ruleType    string
	nameType    string
	languageset []string
}

// NewTokenFilterPhoneticBeiderMorse initializes a new TokenFilterPhoneticBeiderMorse.
func NewTokenFilterPhoneticBeiderMorse(name string) *TokenFilterPhoneticBeiderMorse {
	return &TokenFilterPhoneticBeiderMorse{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (p *TokenFilterPhoneticBeiderMorse) Name() string {
	return p.name
}

// RuleType sets whether matching should be "exact" or "approx".
// Can be set to the following values:
// "exact"
// "approx"
// Defaults to "approx".
func (p *TokenFilterPhoneticBeiderMorse) RuleType(ruleType string) *TokenFilterPhoneticBeiderMorse {
	p.ruleType = ruleType
	return p
}

// NameType sets whether names are "ashkenazi", "sephardic", or "generic".
// Can be set to the following values:
// "ashkenazi"
// "sephardic"
// "generic"
// Defaults to "generic".
func (p *TokenFilterPhoneticBeiderMorse) NameType(nameType string) *TokenFilterPhoneticBeiderMorse {
	p.nameType = nameType
	return p
}

// Languageset sets the languages to check. If not specified, then the language will be guessed.
func (p *TokenFilterPhoneticBeiderMorse) Languageset(languageset ...string) *TokenFilterPhoneticBeiderMorse {
	p.languageset = append(p.languageset, languageset...)
	return p
}

// Validate validates TokenFilterPhoneticBeiderMorse.
func (p *TokenFilterPhoneticBeiderMorse) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if p.ruleType != "" {
		if _, valid := map[string]bool{
			"exact":  true,
			"approx": true,
		}[p.ruleType]; !valid {
			invalid = append(invalid, "RuleType")
		}
	}
	if p.nameType != "" {
		if _, valid := map[string]bool{
			"ashkenazi": true,
			"sephardic": true,
			"generic":   true,
		}[p.nameType]; !valid {
			invalid = append(invalid, "NameType")
		}
	}
	if len(p.languageset) > 0 {
		for _, language := range p.languageset {
			if _, valid := map[string]bool{
				"any":       true,
				"common":    true,
				"cyrillic":  true,
				"english":   true,
				"french":    true,
				"german":    true,
				"hebrew":    true,
				"hungarian": true,
				"polish":    true,
				"romanian":  true,
				"russian":   true,
				"spanish":   true,
			}[language]; !valid {
				invalid = append(invalid, "LanguageSet")
			}
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *TokenFilterPhoneticBeiderMorse) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "phonetic",
	// 		"encoder": "beider_morse",
	// 		"rule_type": "exact",
	// 		"name_type": "ashkenazi",
	// 		"languageset": ["any", "common"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "phonetic"
	options["encoder"] = "beider_morse"

	if p.ruleType != "" {
		options["rule_type"] = p.ruleType
	}
	if p.nameType != "" {
		options["name_type"] = p.nameType
	}
	if len(p.languageset) > 0 {
		options["languageset"] = p.languageset
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
