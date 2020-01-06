// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterKeepTypes token filter that keeps or removes tokens of a specific type.
// For example, you can use this filter to change "3 quick foxes" to "quick foxes" by
// keeping only `<ALPHANUM>` (alphanumeric) tokens.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-keep-types-tokenfilter.html
// for details.
type TokenFilterKeepTypes struct {
	TokenFilter
	name string

	// fields specific to keep types token filter
	types []string
	mode  string
}

// NewTokenFilterKeepTypes initializes a new TokenFilterKeepTypes.
func NewTokenFilterKeepTypes(name string) *TokenFilterKeepTypes {
	return &TokenFilterKeepTypes{
		name:  name,
		types: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (t *TokenFilterKeepTypes) Name() string {
	return t.name
}

// Types sets a list of token types to keep or remove.
func (t *TokenFilterKeepTypes) Types(types ...string) *TokenFilterKeepTypes {
	t.types = append(t.types, types...)
	return t
}

// Mode sets the mode for the filter.
// Can be set to the following values:
// "include" - Keep only the specified token types.
// "exclude" - Remove the specified token types.
// Defaults to "include".
func (t *TokenFilterKeepTypes) Mode(mode string) *TokenFilterKeepTypes {
	t.mode = mode
	return t
}

// Include convenience function that sets the mode for the filter to be "include".
func (t *TokenFilterKeepTypes) Include() *TokenFilterKeepTypes {
	return t.Mode("include")
}

// Exclude convenience function that sets the mode for the filter to be "exclude".
func (t *TokenFilterKeepTypes) Exclude() *TokenFilterKeepTypes {
	return t.Mode("exclude")
}

// Validate validates TokenFilterKeepTypes.
func (t *TokenFilterKeepTypes) Validate(includeName bool) error {
	var invalid []string
	if includeName && t.name == "" {
		invalid = append(invalid, "Name")
	}
	if !(len(t.types) > 0) {
		invalid = append(invalid, "Types")
	}
	if t.mode != "" {
		if _, valid := map[string]bool{
			"include": true,
			"exclude": true,
		}[t.mode]; !valid {
			invalid = append(invalid, "Mode")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (t *TokenFilterKeepTypes) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "keep_types",
	// 		"word_list": ["<ALPHANUM>"],
	// 		"mode": "exclude"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "keep_types"

	if len(t.types) > 0 {
		options["types"] = t.types
	}
	if t.mode != "" {
		options["mode"] = t.mode
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[t.name] = options
	return source, nil
}
