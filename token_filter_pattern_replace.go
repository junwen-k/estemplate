// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterPatternReplace token filter that allows to easily handle string replacements
// based on a regular expression. The regular expression is defined using the `pattern` parameter,
// and the replacement string can be provided using the `replacement` parameter.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-pattern_replace-tokenfilter.html
// for details.
type TokenFilterPatternReplace struct {
	TokenFilter
	name string

	// fields specific to pattern capture token filter
	pattern     string
	replacement string
}

// NewTokenFilterPatternReplace initializes a new TokenFilterPatternReplace.
func NewTokenFilterPatternReplace(name string) *TokenFilterPatternReplace {
	return &TokenFilterPatternReplace{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (r *TokenFilterPatternReplace) Name() string {
	return r.name
}

// Pattern sets the regular expression for matching tokens.
//
// See https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html
// for details.
func (r *TokenFilterPatternReplace) Pattern(pattern string) *TokenFilterPatternReplace {
	r.pattern = pattern
	return r
}

// Replacement sets the replacement string for matched tokens.
func (r *TokenFilterPatternReplace) Replacement(replacement string) *TokenFilterPatternReplace {
	r.replacement = replacement
	return r
}

// Validate validates TokenFilterPatternReplace.
func (r *TokenFilterPatternReplace) Validate(includeName bool) error {
	var invalid []string
	if includeName && r.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (r *TokenFilterPatternReplace) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "pattern_replace",
	// 		"pattern": "\\",
	// 		"replacement": "+"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "pattern_replace"

	if r.pattern != "" {
		options["pattern"] = r.pattern
	}
	if r.replacement != "" {
		options["replacement"] = r.replacement
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[r.name] = options
	return source, nil
}
