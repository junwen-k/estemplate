// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeByte Core Datatype for numeric value.
// A signed 8-bit integer with a minimum value of -128 and a maximum value of 127.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeByte struct {
	Datatype
	name string

	// fields specific to byte datatype
	coerce          *bool
	boost           *int
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
}

// NewDatatypeByte initializes a new DatatypeByte.
func NewDatatypeByte(name string) *DatatypeByte {
	return &DatatypeByte{
		name: name,
	}
}

// Name is the key of the Byte Property.
func (b *DatatypeByte) Name() string {
	return b.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (b *DatatypeByte) Coerce(coerce bool) *DatatypeByte {
	b.coerce = &coerce
	return b
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (b *DatatypeByte) Boost(boost int) *DatatypeByte {
	b.boost = &boost
	return b
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (b *DatatypeByte) DocValues(docValues bool) *DatatypeByte {
	b.docValues = &docValues
	return b
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defatuls to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (b *DatatypeByte) IgnoreMalformed(ignoreMalformed bool) *DatatypeByte {
	b.ignoreMalformed = &ignoreMalformed
	return b
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (b *DatatypeByte) Index(index bool) *DatatypeByte {
	b.index = &index
	return b
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (b *DatatypeByte) NullValue(nullValue int) *DatatypeByte {
	b.nullValue = &nullValue
	return b
}

// Store sets whether the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (b *DatatypeByte) Store(store bool) *DatatypeByte {
	b.store = &store
	return b
}

// Validate validates DatatypeByte.
func (b *DatatypeByte) Validate(includeName bool) error {
	var invalid []string
	if includeName && b.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (b *DatatypeByte) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "byte",
	// 		"coerce": true,
	// 		"boost": 2,
	// 		"doc_values": true,
	// 		"ignore_malformed": true,
	// 		"index": true,
	// 		"null_value": 0,
	// 		"store": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "byte"

	if b.coerce != nil {
		options["coerce"] = b.coerce
	}
	if b.boost != nil {
		options["boost"] = b.boost
	}
	if b.docValues != nil {
		options["doc_values"] = b.docValues
	}
	if b.ignoreMalformed != nil {
		options["ignore_malformed"] = b.ignoreMalformed
	}
	if b.index != nil {
		options["index"] = b.index
	}
	if b.nullValue != nil {
		options["null_value"] = b.nullValue
	}
	if b.store != nil {
		options["store"] = b.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[b.name] = options
	return source, nil
}