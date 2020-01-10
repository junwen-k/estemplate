// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// CharacterFilterHTMLStrip character filter that strips HTML elements from the text
// and replaces HTML entities with their decoded value (e.g. replacing &amp; with &).
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-htmlstrip-charfilter.html
// for details.
type CharacterFilterHTMLStrip struct {
	CharacterFilter
	name string

	// fields specific to html strip character filter
	escapedTags []string
}

// NewCharacterFilterHTMLStrip initializes a new CharacterFilterHTMLStrip.
func NewCharacterFilterHTMLStrip(name string) *CharacterFilterHTMLStrip {
	return &CharacterFilterHTMLStrip{
		name:        name,
		escapedTags: make([]string, 0),
	}
}

// Name returns field key for the Character Filter.
func (s *CharacterFilterHTMLStrip) Name() string {
	return s.name
}

// EscapedTags sets an array of HTML tags which should not be stripped from the
// original text.
func (s *CharacterFilterHTMLStrip) EscapedTags(escapedTags ...string) *CharacterFilterHTMLStrip {
	s.escapedTags = append(s.escapedTags, escapedTags...)
	return s
}

// Validate validates CharacterFilterHTMLStrip.
func (s *CharacterFilterHTMLStrip) Validate(includeName bool) error {
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
func (s *CharacterFilterHTMLStrip) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "html_strip",
	// 		"escaped_tags": ["b"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "html_strip"

	if len(s.escapedTags) > 0 {
		options["escaped_tags"] = s.escapedTags
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
