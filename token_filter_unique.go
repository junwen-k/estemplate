// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterUnique token filter that can be used to only index unique tokens during
// analysis. By default it is applied on all the token stream.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-unique-tokenfilter.html
// for details.
type TokenFilterUnique struct {
	TokenFilter
	name string

	// fields specific to unique token filter
	onlyOnSamePosition *bool
}

// NewTokenFilterUnique initializes a new TokenFilterUnique.
func NewTokenFilterUnique(name string) *TokenFilterUnique {
	return &TokenFilterUnique{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (u *TokenFilterUnique) Name() string {
	return u.name
}

// OnlyOnSamePosition sets whether to only remove duplicate tokens on the same position.
func (u *TokenFilterUnique) OnlyOnSamePosition(onlyOnSamePosition bool) *TokenFilterUnique {
	u.onlyOnSamePosition = &onlyOnSamePosition
	return u
}

// Validate validates TokenFilterUnique.
func (u *TokenFilterUnique) Validate(includeName bool) error {
	var invalid []string
	if includeName && u.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (u *TokenFilterUnique) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "unique",
	// 		"only_on_same_position": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "unique"

	if u.onlyOnSamePosition != nil {
		options["only_on_same_position"] = u.onlyOnSamePosition
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[u.name] = options
	return source, nil
}
