// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeTokenCount Specialised Datatype is really an integer field which
// accepts string values, analzyes them, then indexes the number of tokens
// in the string.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/token-count.html
// for details.
type DatatypeTokenCount struct {
	Datatype
	name string

	// fields specific to token count datatype
	analyzer                 string
	enablePositionIncrements *bool
	boost                    *float32
	docValues                *bool
	index                    *bool
	nullValue                *int
	store                    *bool
}

// NewDatatypeTokenCount initializes a new DatatypeTokenCount.
func NewDatatypeTokenCount(name string) *DatatypeTokenCount {
	return &DatatypeTokenCount{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (c *DatatypeTokenCount) Name() string {
	return c.name
}

// Analyzer sets which analyzer should be used for analyzed string fields.
// * Required field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analyzer.html
// for details.
func (c *DatatypeTokenCount) Analyzer(analyzer string) *DatatypeTokenCount {
	c.analyzer = analyzer
	return c
}

// EnablePositionIncrements sets whether if position increments should be counted.
// Set to false if you don't want to count tokens removed by analyzer filters (like `stop`).
// Defaults to true.
func (c *DatatypeTokenCount) EnablePositionIncrements(enablePositionIncrements bool) *DatatypeTokenCount {
	c.enablePositionIncrements = &enablePositionIncrements
	return c
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (c *DatatypeTokenCount) Boost(boost float32) *DatatypeTokenCount {
	c.boost = &boost
	return c
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (c *DatatypeTokenCount) DocValues(docValues bool) *DatatypeTokenCount {
	c.docValues = &docValues
	return c
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (c *DatatypeTokenCount) Index(index bool) *DatatypeTokenCount {
	c.index = &index
	return c
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (c *DatatypeTokenCount) NullValue(nullValue int) *DatatypeTokenCount {
	c.nullValue = &nullValue
	return c
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (c *DatatypeTokenCount) Store(store bool) *DatatypeTokenCount {
	c.store = &store
	return c
}

// Validate validates DatatypeTokenCount.
func (c *DatatypeTokenCount) Validate(includeName bool) error {
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
func (c *DatatypeTokenCount) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "token_count",
	// 		"analyzer": "standard",
	//    "enable_position_increments": true,
	// 		"boost": 2,
	// 		"doc_values": true,
	// 		"index": true,
	// 		"null_value": 0,
	// 		"store": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "token_count"

	if c.analyzer != "" {
		options["analyzer"] = c.analyzer
	}
	if c.enablePositionIncrements != nil {
		options["enable_position_increments"] = c.enablePositionIncrements
	}
	if c.boost != nil {
		options["boost"] = c.boost
	}
	if c.docValues != nil {
		options["doc_values"] = c.docValues
	}
	if c.index != nil {
		options["index"] = c.index
	}
	if c.nullValue != nil {
		options["null_value"] = c.nullValue
	}
	if c.store != nil {
		options["store"] = c.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
