// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// CharacterFilterMappingChar character filter that accepts a map of keys and values.
// Whenever it encounters a string of characters that is the same as a key, it replaces
// them with the value associated with that key.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-mapping-charfilter.html
// for details.
type CharacterFilterMappingChar struct {
	CharacterFilter
	name string

	// fields specific to mapping char character filter
	mappings     []*MappingRule
	rawMappings  []string
	mappingsPath string
}

// NewCharacterFilterMappingChar initializes a new CharacterFilterMappingChar.
func NewCharacterFilterMappingChar(name string) *CharacterFilterMappingChar {
	return &CharacterFilterMappingChar{
		name:        name,
		mappings:    make([]*MappingRule, 0),
		rawMappings: make([]string, 0),
	}
}

// Name returns field key for the Character Filter.
func (c *CharacterFilterMappingChar) Name() string {
	return c.name
}

// Mappings sets an array of mappings, with each element having the form `key => value`.
func (c *CharacterFilterMappingChar) Mappings(mappings ...*MappingRule) *CharacterFilterMappingChar {
	c.mappings = append(c.mappings, mappings...)
	return c
}

// RawMappings sets an array of mappings, with each element having the form `key => value`. Use this if you prefer
// to use strings instead of MappingRule type.
func (c *CharacterFilterMappingChar) RawMappings(rawMappings ...string) *CharacterFilterMappingChar {
	c.rawMappings = append(c.rawMappings, rawMappings...)
	return c
}

// MappingsPath sets a path to a file containing a `key => value` mappings. This path must be
// absolute or relative to the `config` location. The file must be UTF-8 encoded.
// Each mapping in the file must be separated by a line break.
func (c *CharacterFilterMappingChar) MappingsPath(mappingsPath string) *CharacterFilterMappingChar {
	c.mappingsPath = mappingsPath
	return c
}

// Validate validates CharacterFilterMappingChar.
func (c *CharacterFilterMappingChar) Validate(includeName bool) error {
	var invalid []string
	if includeName && c.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (c *CharacterFilterMappingChar) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "mapping",
	// 		"mappings": ["٠ => 0", "١ => 1", "٢ => 2"],
	// 		"mappings_path": "analysis/mappings.txt"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "mapping"

	if len(c.mappings) > 0 {
		mappings := make([]string, 0)
		for _, m := range c.mappings {
			mapping, err := m.Source()
			if err != nil {
				return nil, err
			}
			mappings = append(mappings, fmt.Sprintf("%s", mapping))
		}
		options["mappings"] = mappings
	}
	if len(c.rawMappings) > 0 {
		var mappings interface{}
		switch {
		case len(c.rawMappings) > 1:
			mappings = c.rawMappings
			break
		case len(c.rawMappings) == 1:
			mappings = c.rawMappings[0]
			break
		default:
			mappings = ""
		}
		options["mappings"] = mappings
	}
	if c.mappingsPath != "" {
		options["mappings_path"] = c.mappingsPath
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
