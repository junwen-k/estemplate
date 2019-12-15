// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeFloat Core Datatype for numeric value.
// A single-precision 32-bit IEEE 754 floating point number, restricted to finite values.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeFloat struct {
	Datatype
	name string

	// fields specific to float datatype
	coerce          *bool
	boost           *int
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
}

// NewDatatypeFloat initializes a new DatatypeFloat.
func NewDatatypeFloat(name string) *DatatypeFloat {
	return &DatatypeFloat{
		name: name,
	}
}

// Name is the key of the Float Property.
func (f *DatatypeFloat) Name() string {
	return f.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (f *DatatypeFloat) Coerce(coerce bool) *DatatypeFloat {
	f.coerce = &coerce
	return f
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (f *DatatypeFloat) Boost(boost int) *DatatypeFloat {
	f.boost = &boost
	return f
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (f *DatatypeFloat) DocValues(docValues bool) *DatatypeFloat {
	f.docValues = &docValues
	return f
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defatuls to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (f *DatatypeFloat) IgnoreMalformed(ignoreMalformed bool) *DatatypeFloat {
	f.ignoreMalformed = &ignoreMalformed
	return f
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (f *DatatypeFloat) Index(index bool) *DatatypeFloat {
	f.index = &index
	return f
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (f *DatatypeFloat) NullValue(nullValue int) *DatatypeFloat {
	f.nullValue = &nullValue
	return f
}

// Store sets whether the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (f *DatatypeFloat) Store(store bool) *DatatypeFloat {
	f.store = &store
	return f
}

// Validate validates DatatypeFloat.
func (f *DatatypeFloat) Validate(includeName bool) error {
	var invalid []string
	if includeName && f.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (f *DatatypeFloat) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "float",
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
	options["type"] = "float"

	if f.coerce != nil {
		options["coerce"] = f.coerce
	}
	if f.boost != nil {
		options["boost"] = f.boost
	}
	if f.docValues != nil {
		options["doc_values"] = f.docValues
	}
	if f.ignoreMalformed != nil {
		options["ignore_malformed"] = f.ignoreMalformed
	}
	if f.index != nil {
		options["index"] = f.index
	}
	if f.nullValue != nil {
		options["null_value"] = f.nullValue
	}
	if f.store != nil {
		options["store"] = f.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[f.name] = options
	return source, nil
}
