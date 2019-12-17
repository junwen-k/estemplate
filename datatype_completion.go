// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeCompletion Specialised Datatype for auto-complete suggestions datatype.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-suggesters.html#completion-suggester
// for details.
type DatatypeCompletion struct {
	Datatype
	name string

	// fields specific to completion datatype
	analzyer                   string
	searchAnalzyer             string
	preserveSeparators         *bool
	preservePositionIncrements *bool
	maxInputLength             *int
}

// NewDatatypeCompletion initializes a new DatatypeCompletion.
func NewDatatypeCompletion(name string) *DatatypeCompletion {
	return &DatatypeCompletion{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (c *DatatypeCompletion) Name() string {
	return c.name
}

// Analzyer sets which analzyer should be used for indexing.
// Defaults to "simple".
func (c *DatatypeCompletion) Analzyer(analzyer string) *DatatypeCompletion {
	c.analzyer = analzyer
	return c
}

// SearchAnalzyer sets the analzyer that should be used at search time on analyzed fields
// Defaults to `analzyer` setting.
func (c *DatatypeCompletion) SearchAnalzyer(searchAnalzyer string) *DatatypeCompletion {
	c.searchAnalzyer = searchAnalzyer
	return c
}

// PreserveSeparators sets whether should preserve separators in the process of indexing.
// If disabled, you could find a field starting with "Foo Fighters", if you suggest for "foof".
// Defaults to true.
func (c *DatatypeCompletion) PreserveSeparators(preserveSeparators bool) *DatatypeCompletion {
	c.preserveSeparators = &preserveSeparators
	return c
}

// PreservePositionIncrements sets whether to use position increments in the processing of indexing.
// If disabled, you could get a field starting with "The Beatles", if you suggest for "b".
// Defaults to true.
func (c *DatatypeCompletion) PreservePositionIncrements(preservePositionIncrements bool) *DatatypeCompletion {
	c.preservePositionIncrements = &preservePositionIncrements
	return c
}

// MaxInputLength sets the limit length of a single input. This limit is only used at index
// time to reduce the total number of characters per input string in order to prevent massive
// inputs from bloating the underlying datastructure.
// Defaults to 50.
func (c *DatatypeCompletion) MaxInputLength(maxInputLength int) *DatatypeCompletion {
	c.maxInputLength = &maxInputLength
	return c
}

// Validate validates DatatypeCompletion.
func (c *DatatypeCompletion) Validate(includeName bool) error {
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
func (c *DatatypeCompletion) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "completion",
	// 		"analzyer": "simple",
	// 		"search_analzyer": "standard",
	// 		"preserve_separators": true,
	// 		"preserve_position_increments": true,
	// 		"max_input_length": 50
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "completion"

	if c.analzyer != "" {
		options["analzyer"] = c.analzyer
	}
	if c.searchAnalzyer != "" {
		options["search_analzyer"] = c.searchAnalzyer
	}
	if c.preserveSeparators != nil {
		options["preserve_separators"] = c.preserveSeparators
	}
	if c.preservePositionIncrements != nil {
		options["preserve_position_increments"] = c.preservePositionIncrements
	}
	if c.maxInputLength != nil {
		options["max_input_length"] = c.maxInputLength
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
