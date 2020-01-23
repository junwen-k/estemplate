// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeFloatRange Core Datatype for float range.
// A range of single-precision 32-bit IEEE 754 floating point values.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/range.html
// for details.
type DatatypeFloatRange struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to float range datatype
	coerce *bool
	boost  *float32
	index  *bool
	store  *bool
}

// NewDatatypeFloatRange initializes a new DatatypeFloatRange.
func NewDatatypeFloatRange(name string) *DatatypeFloatRange {
	return &DatatypeFloatRange{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (r *DatatypeFloatRange) Name() string {
	return r.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (r *DatatypeFloatRange) CopyTo(copyTo ...string) *DatatypeFloatRange {
	r.copyTo = append(r.copyTo, copyTo...)
	return r
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (r *DatatypeFloatRange) Coerce(coerce bool) *DatatypeFloatRange {
	r.coerce = &coerce
	return r
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (r *DatatypeFloatRange) Boost(boost float32) *DatatypeFloatRange {
	r.boost = &boost
	return r
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (r *DatatypeFloatRange) Index(index bool) *DatatypeFloatRange {
	r.index = &index
	return r
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (r *DatatypeFloatRange) Store(store bool) *DatatypeFloatRange {
	r.store = &store
	return r
}

// Validate validates DatatypeFloatRange.
func (r *DatatypeFloatRange) Validate(includeName bool) error {
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
func (r *DatatypeFloatRange) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "float_range",
	// 		"copy_to": ["field_1", "field_2"],
	// 		"coerce": true,
	// 		"boost": 2,
	// 		"index": true,
	// 		"store": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "float_range"

	if len(r.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(r.copyTo) > 1:
			copyTo = r.copyTo
			break
		case len(r.copyTo) == 1:
			copyTo = r.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
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
