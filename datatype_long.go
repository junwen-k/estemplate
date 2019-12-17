// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeLong Core Datatype for numeric value.
// A signed 64-bit integer with a minimum value of -2⁶³ and a maximum value of 2⁶³-1.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeLong struct {
	Datatype
	name string

	// fields specific to long datatype
	coerce          *bool
	boost           *float32
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
}

// NewDatatypeLong initializes a new DatatypeLong.
func NewDatatypeLong(name string) *DatatypeLong {
	return &DatatypeLong{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (l *DatatypeLong) Name() string {
	return l.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (l *DatatypeLong) Coerce(coerce bool) *DatatypeLong {
	l.coerce = &coerce
	return l
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (l *DatatypeLong) Boost(boost float32) *DatatypeLong {
	l.boost = &boost
	return l
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (l *DatatypeLong) DocValues(docValues bool) *DatatypeLong {
	l.docValues = &docValues
	return l
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (l *DatatypeLong) IgnoreMalformed(ignoreMalformed bool) *DatatypeLong {
	l.ignoreMalformed = &ignoreMalformed
	return l
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (l *DatatypeLong) Index(index bool) *DatatypeLong {
	l.index = &index
	return l
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (l *DatatypeLong) NullValue(nullValue int) *DatatypeLong {
	l.nullValue = &nullValue
	return l
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (l *DatatypeLong) Store(store bool) *DatatypeLong {
	l.store = &store
	return l
}

// Validate validates DatatypeLong.
func (l *DatatypeLong) Validate(includeName bool) error {
	var invalid []string
	if includeName && l.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (l *DatatypeLong) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "long",
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
	options["type"] = "long"

	if l.coerce != nil {
		options["coerce"] = l.coerce
	}
	if l.boost != nil {
		options["boost"] = l.boost
	}
	if l.docValues != nil {
		options["doc_values"] = l.docValues
	}
	if l.ignoreMalformed != nil {
		options["ignore_malformed"] = l.ignoreMalformed
	}
	if l.index != nil {
		options["index"] = l.index
	}
	if l.nullValue != nil {
		options["null_value"] = l.nullValue
	}
	if l.store != nil {
		options["store"] = l.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[l.name] = options
	return source, nil
}
