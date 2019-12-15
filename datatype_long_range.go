// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeLongRange Core Datatype for long range.
// A range of signed 64-bit integers with a minimum value of -2⁶³ and maximum of 2⁶³-1.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/range.html
// for details.
type DatatypeLongRange struct {
	Datatype
	name string

	// fields specific to long range datatype
	coerce *bool
	boost  *int
	index  *bool
	store  *bool
}

// NewDatatypeLongRange initializes a new DatatypeLongRange.
func NewDatatypeLongRange(name string) *DatatypeLongRange {
	return &DatatypeLongRange{
		name: name,
	}
}

// Name is the key of the Byte Property.
func (r *DatatypeLongRange) Name() string {
	return r.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (r *DatatypeLongRange) Coerce(coerce bool) *DatatypeLongRange {
	r.coerce = &coerce
	return r
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (r *DatatypeLongRange) Boost(boost int) *DatatypeLongRange {
	r.boost = &boost
	return r
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (r *DatatypeLongRange) Index(index bool) *DatatypeLongRange {
	r.index = &index
	return r
}

// Store sets whether the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (r *DatatypeLongRange) Store(store bool) *DatatypeLongRange {
	r.store = &store
	return r
}

// Validate validates DatatypeLongRange.
func (r *DatatypeLongRange) Validate(includeName bool) error {
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
func (r *DatatypeLongRange) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "long_range",
	// 		"coerce": true,
	// 		"boost": 2,
	// 		"index": true,
	// 		"store": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "long_range"

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
