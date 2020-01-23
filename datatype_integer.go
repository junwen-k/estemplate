// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeInteger Core Datatype for numeric value.
// A signed 32-bit integer with a minimum value of -2³¹ and a maximum value of 2³¹-1.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeInteger struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to integer datatype
	coerce          *bool
	boost           *float32
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
}

// NewDatatypeInteger initializes a new DatatypeInteger.
func NewDatatypeInteger(name string) *DatatypeInteger {
	return &DatatypeInteger{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (i *DatatypeInteger) Name() string {
	return i.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (i *DatatypeInteger) CopyTo(copyTo ...string) *DatatypeInteger {
	i.copyTo = append(i.copyTo, copyTo...)
	return i
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (i *DatatypeInteger) Coerce(coerce bool) *DatatypeInteger {
	i.coerce = &coerce
	return i
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (i *DatatypeInteger) Boost(boost float32) *DatatypeInteger {
	i.boost = &boost
	return i
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (i *DatatypeInteger) DocValues(docValues bool) *DatatypeInteger {
	i.docValues = &docValues
	return i
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (i *DatatypeInteger) IgnoreMalformed(ignoreMalformed bool) *DatatypeInteger {
	i.ignoreMalformed = &ignoreMalformed
	return i
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (i *DatatypeInteger) Index(index bool) *DatatypeInteger {
	i.index = &index
	return i
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (i *DatatypeInteger) NullValue(nullValue int) *DatatypeInteger {
	i.nullValue = &nullValue
	return i
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (i *DatatypeInteger) Store(store bool) *DatatypeInteger {
	i.store = &store
	return i
}

// Validate validates DatatypeInteger.
func (i *DatatypeInteger) Validate(includeName bool) error {
	var invalid []string
	if includeName && i.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (i *DatatypeInteger) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "integer",
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
	options["type"] = "integer"

	if len(i.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(i.copyTo) > 1:
			copyTo = i.copyTo
			break
		case len(i.copyTo) == 1:
			copyTo = i.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if i.coerce != nil {
		options["coerce"] = i.coerce
	}
	if i.boost != nil {
		options["boost"] = i.boost
	}
	if i.docValues != nil {
		options["doc_values"] = i.docValues
	}
	if i.ignoreMalformed != nil {
		options["ignore_malformed"] = i.ignoreMalformed
	}
	if i.index != nil {
		options["index"] = i.index
	}
	if i.nullValue != nil {
		options["null_value"] = i.nullValue
	}
	if i.store != nil {
		options["store"] = i.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[i.name] = options
	return source, nil
}
