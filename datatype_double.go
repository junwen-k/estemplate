// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeDouble Core Datatype for numeric value.
// A double-precision 64-bit IEEE 754 floating point number, restricted to finite values.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeDouble struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to double datatype
	coerce          *bool
	boost           *float32
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
}

// NewDatatypeDouble initializes a new DatatypeDouble.
func NewDatatypeDouble(name string) *DatatypeDouble {
	return &DatatypeDouble{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (d *DatatypeDouble) Name() string {
	return d.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (d *DatatypeDouble) CopyTo(copyTo ...string) *DatatypeDouble {
	d.copyTo = append(d.copyTo, copyTo...)
	return d
}

// Coerce sets whether if the field should be coerced, attempting to clean up
//
// dirty values to fit the datatype.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (d *DatatypeDouble) Coerce(coerce bool) *DatatypeDouble {
	d.coerce = &coerce
	return d
}

//
// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (d *DatatypeDouble) Boost(boost float32) *DatatypeDouble {
	d.boost = &boost
	return d
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
//
// Defaults to true.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (d *DatatypeDouble) DocValues(docValues bool) *DatatypeDouble {
	d.docValues = &docValues
	return d
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
//
// Defaults to false.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (d *DatatypeDouble) IgnoreMalformed(ignoreMalformed bool) *DatatypeDouble {
	d.ignoreMalformed = &ignoreMalformed
	return d
}

//
// Index sets whether if the field should be searchable. Defaults to true.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (d *DatatypeDouble) Index(index bool) *DatatypeDouble {
	d.index = &index
	return d
}

// NullValue sets a numeric value which is substituted for any explicit null values.
//
// Defaults to null.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (d *DatatypeDouble) NullValue(nullValue int) *DatatypeDouble {
	d.nullValue = &nullValue
	return d
}

// Store sets whether if the field value should be stored and retrievable separately
//
// from the `_source` field. Defaults to false.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (d *DatatypeDouble) Store(store bool) *DatatypeDouble {
	d.store = &store
	return d
}

// Validate validates DatatypeDouble.
func (d *DatatypeDouble) Validate(includeName bool) error {
	var invalid []string
	if includeName && d.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (d *DatatypeDouble) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "double",
	// 		"copy_to": ["field_1", "field_2"],
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
	options["type"] = "double"

	if len(d.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(d.copyTo) > 1:
			copyTo = d.copyTo
			break
		case len(d.copyTo) == 1:
			copyTo = d.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if d.coerce != nil {
		options["coerce"] = d.coerce
	}
	if d.boost != nil {
		options["boost"] = d.boost
	}
	if d.docValues != nil {
		options["doc_values"] = d.docValues
	}
	if d.ignoreMalformed != nil {
		options["ignore_malformed"] = d.ignoreMalformed
	}
	if d.index != nil {
		options["index"] = d.index
	}
	if d.nullValue != nil {
		options["null_value"] = d.nullValue
	}
	if d.store != nil {
		options["store"] = d.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[d.name] = options
	return source, nil
}
