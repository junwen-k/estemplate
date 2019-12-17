// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeIntegerRange Core Datatype for integer range.
// A range of signed 32-bit integers with a minimum value of -2³¹ and maximum of 2³¹-1.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/range.html
// for details.
type DatatypeIntegerRange struct {
	Datatype
	name string

	// fields specific to integer range datatype
	coerce *bool
	boost  *float32
	index  *bool
	store  *bool
}

// NewDatatypeIntegerRange initializes a new DatatypeIntegerRange.
func NewDatatypeIntegerRange(name string) *DatatypeIntegerRange {
	return &DatatypeIntegerRange{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (r *DatatypeIntegerRange) Name() string {
	return r.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (r *DatatypeIntegerRange) Coerce(coerce bool) *DatatypeIntegerRange {
	r.coerce = &coerce
	return r
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (r *DatatypeIntegerRange) Boost(boost float32) *DatatypeIntegerRange {
	r.boost = &boost
	return r
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (r *DatatypeIntegerRange) Index(index bool) *DatatypeIntegerRange {
	r.index = &index
	return r
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (r *DatatypeIntegerRange) Store(store bool) *DatatypeIntegerRange {
	r.store = &store
	return r
}

// Validate validates DatatypeIntegerRange.
func (r *DatatypeIntegerRange) Validate(includeName bool) error {
	var invalid []string
	if includeName && r.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (r *DatatypeIntegerRange) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "integer_range",
	// 		"coerce": true,
	// 		"boost": 2,
	// 		"index": true,
	// 		"store": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "integer_range"

	if r.coerce != nil {
		options["coerce"] = r.coerce
	}
	if r.boost != nil {
		options["boost"] = r.boost
	}
	if r.index != nil {
		options["index"] = r.index
	}
	if r.store != nil {
		options["store"] = r.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[r.name] = options
	return source, nil
}
