// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterPredicateScript token filter that takes a predicate script, and removes tokens
// that do not match the predicate.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-predicatefilter-tokenfilter.html
// for details.
type TokenFilterPredicateScript struct {
	TokenFilter
	name string

	// fields specific to predicate script token filter
	script *Script
}

// NewTokenFilterPredicateScript initializes a new TokenFilterPredicateScript.
func NewTokenFilterPredicateScript(name string) *TokenFilterPredicateScript {
	return &TokenFilterPredicateScript{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (s *TokenFilterPredicateScript) Name() string {
	return s.name
}

// Script sets the predicate script that determines whether or not the current token will
// be emitted. Note that only inline scripts are supported.
func (s *TokenFilterPredicateScript) Script(script *Script) *TokenFilterPredicateScript {
	s.script = script
	return s
}

// Validate validates TokenFilterPredicateScript.
func (s *TokenFilterPredicateScript) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *TokenFilterPredicateScript) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "predicate_token_filter",
	// 		"script": {
	// 			"source": "token.getTerm().length() > 5"
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "predicate_token_filter"

	if s.script != nil {
		script, err := s.script.Source(false)
		if err != nil {
			return nil, err
		}
		options["script"] = script
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
