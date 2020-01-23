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
	name   string
	copyTo []string

	// fields specific to completion datatype
	analyzer                   string
	searchAnalyzer             string
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

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (c *DatatypeCompletion) CopyTo(copyTo ...string) *DatatypeCompletion {
	c.copyTo = append(c.copyTo, copyTo...)
	return c
}

// Analyzer sets which analyzer should be used for indexing.
// Defaults to "simple".
func (c *DatatypeCompletion) Analyzer(analyzer string) *DatatypeCompletion {
	c.analyzer = analyzer
	return c
}

// SearchAnalyzer sets the analyzer that should be used at search time on analyzed fields
// Defaults to `analyzer` setting.
func (c *DatatypeCompletion) SearchAnalyzer(searchAnalyzer string) *DatatypeCompletion {
	c.searchAnalyzer = searchAnalyzer
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
	// 		"copy_to": ["field_1", "field_2"],
	// 		"analyzer": "simple",
	// 		"search_analyzer": "standard",
	// 		"preserve_separators": true,
	// 		"preserve_position_increments": true,
	// 		"max_input_length": 50
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "completion"

	if len(c.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(c.copyTo) > 1:
			copyTo = c.copyTo
			break
		case len(c.copyTo) == 1:
			copyTo = c.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if c.analyzer != "" {
		options["analyzer"] = c.analyzer
	}
	if c.searchAnalyzer != "" {
		options["search_analyzer"] = c.searchAnalyzer
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
