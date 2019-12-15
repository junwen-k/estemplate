// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeHalfFloat Core Datatype for numeric value.
// A half-precision 16-bit IEEE 754 floating point number, restricted to finite values.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeHalfFloat struct {
	Datatype
	name string

	// fields specific to half float datatype
	coerce          *bool
	boost           *int
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
}

// NewDatatypeHalfFloat initializes a new DatatypeHalfFloat.
func NewDatatypeHalfFloat(name string) *DatatypeHalfFloat {
	return &DatatypeHalfFloat{
		name: name,
	}
}

// Name is the key of the Half Float Property.
func (hf *DatatypeHalfFloat) Name() string {
	return hf.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (hf *DatatypeHalfFloat) Coerce(coerce bool) *DatatypeHalfFloat {
	hf.coerce = &coerce
	return hf
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (hf *DatatypeHalfFloat) Boost(boost int) *DatatypeHalfFloat {
	hf.boost = &boost
	return hf
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (hf *DatatypeHalfFloat) DocValues(docValues bool) *DatatypeHalfFloat {
	hf.docValues = &docValues
	return hf
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defatuls to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (hf *DatatypeHalfFloat) IgnoreMalformed(ignoreMalformed bool) *DatatypeHalfFloat {
	hf.ignoreMalformed = &ignoreMalformed
	return hf
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (hf *DatatypeHalfFloat) Index(index bool) *DatatypeHalfFloat {
	hf.index = &index
	return hf
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (hf *DatatypeHalfFloat) NullValue(nullValue int) *DatatypeHalfFloat {
	hf.nullValue = &nullValue
	return hf
}

// Store sets whether the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (hf *DatatypeHalfFloat) Store(store bool) *DatatypeHalfFloat {
	hf.store = &store
	return hf
}

// Validate validates DatatypeHalfFloat.
func (hf *DatatypeHalfFloat) Validate(includeName bool) error {
	var invalid []string
	if includeName && hf.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (hf *DatatypeHalfFloat) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "half_float",
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
	options["type"] = "half_float"

	if hf.coerce != nil {
		options["coerce"] = hf.coerce
	}
	if hf.boost != nil {
		options["boost"] = hf.boost
	}
	if hf.docValues != nil {
		options["doc_values"] = hf.docValues
	}
	if hf.ignoreMalformed != nil {
		options["ignore_malformed"] = hf.ignoreMalformed
	}
	if hf.index != nil {
		options["index"] = hf.index
	}
	if hf.nullValue != nil {
		options["null_value"] = hf.nullValue
	}
	if hf.store != nil {
		options["store"] = hf.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[hf.name] = options
	return source, nil
}
