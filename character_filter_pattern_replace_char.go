// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"fmt"
	"strings"
)

// CharacterFilterPatternReplaceChar character filter that uses a regular expression to match characters
// which should be replaced with the specified replacement string. The replacement string can refer to capture
// groups in the regular expression.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-pattern-replace-charfilter.html
// for details.
type CharacterFilterPatternReplaceChar struct {
	CharacterFilter
	name string

	// fields specific to pattern replace char character filter
	pattern     string
	replacement string
	flags       []string
}

// NewCharacterFilterPatternReplaceChar initializes a new CharacterFilterPatternReplaceChar.
func NewCharacterFilterPatternReplaceChar(name string) *CharacterFilterPatternReplaceChar {
	return &CharacterFilterPatternReplaceChar{
		name:  name,
		flags: make([]string, 0),
	}
}

// Name returns field key for the Character Filter.
func (c *CharacterFilterPatternReplaceChar) Name() string {
	return c.name
}

// Pattern sets the Java regular expression for the character filter.
//
// See https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html
// for details.
func (c *CharacterFilterPatternReplaceChar) Pattern(pattern string) *CharacterFilterPatternReplaceChar {
	c.pattern = pattern
	return c
}

// Replacement sets the replacement string which can reference capture groups using the
// `$1 .. $9` syntax.
//
// See https://docs.oracle.com/javase/8/docs/api/java/util/regex/Matcher.html#appendReplacement-java.lang.StringBuffer-java.lang.String-
// for details.
func (c *CharacterFilterPatternReplaceChar) Replacement(replacement string) *CharacterFilterPatternReplaceChar {
	c.replacement = replacement
	return c
}

// Flags sets Java regular expression flags. eg "CASE_INSENSITIVE|COMMENTS".
//
// See https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html#field.summary
// for details.
func (c *CharacterFilterPatternReplaceChar) Flags(flags ...string) *CharacterFilterPatternReplaceChar {
	c.flags = append(c.flags, flags...)
	return c
}

// Validate validates CharacterFilterPatternReplaceChar.
func (c *CharacterFilterPatternReplaceChar) Validate(includeName bool) error {
	var invalid []string
	if includeName && c.name == "" {
		invalid = append(invalid, "Name")
	}
	if c.pattern == "" {
		invalid = append(invalid, "Pattern")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (c *CharacterFilterPatternReplaceChar) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "pattern_replace",
	// 		"pattern": "(\\d+)-(?=\\d)",
	// 		"replacement": "$1_"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "pattern_replace"

	if c.pattern != "" {
		options["pattern"] = c.pattern
	}
	if c.replacement != "" {
		options["replacement"] = c.replacement
	}
	if len(c.flags) > 0 {
		options["flags"] = strings.Join(c.flags, "|")
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
