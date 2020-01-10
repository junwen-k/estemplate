// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterStemmerOverride token filter
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-stemmer-override-tokenfilter.html
// for details.
type TokenFilterStemmerOverride struct {
	TokenFilter
	name string

	// fields specific to stemmer override token filter
	rules     []*MappingRule
	rulesPath string
}

// NewTokenFilterStemmerOverride initializes a new TokenFilterStemmerOverride.
func NewTokenFilterStemmerOverride(name string) *TokenFilterStemmerOverride {
	return &TokenFilterStemmerOverride{
		name:  name,
		rules: make([]*MappingRule, 0),
	}
}

// Name returns field key for the Token Filter.
func (o *TokenFilterStemmerOverride) Name() string {
	return o.name
}

// Rules sets a list of mapping rules to use.
func (o *TokenFilterStemmerOverride) Rules(rules ...*MappingRule) *TokenFilterStemmerOverride {
	o.rules = append(o.rules, rules...)
	return o
}

// RulesPath sets a path to a list of mappings.
// This path must be absolute or relative to the `config` location. The file must be UTF-8 encoded.
// Rules are separated by "=>"
// Each rules in the file must be separated by a line break.
func (o *TokenFilterStemmerOverride) RulesPath(rulesPath string) *TokenFilterStemmerOverride {
	o.rulesPath = rulesPath
	return o
}

// Validate validates TokenFilterStemmerOverride.
func (o *TokenFilterStemmerOverride) Validate(includeName bool) error {
	var invalid []string
	if includeName && o.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (o *TokenFilterStemmerOverride) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "stemmer_override",
	// 		"rules" : [
	// 			"running => run",
	// 			"stemmer => stemmer"
	// 		]
	// 		"rules_path": "analysis/stemmer_override.txt"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "stemmer_override"

	if len(o.rules) > 0 {
		rules := make([]string, 0)
		for _, r := range o.rules {
			rule, err := r.Source()
			if err != nil {
				return nil, err
			}
			rules = append(rules, fmt.Sprintf("%s", rule))
		}
		options["rules"] = rules
	}
	if o.rulesPath != "" {
		options["rules_path"] = o.rulesPath
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[o.name] = options
	return source, nil
}
