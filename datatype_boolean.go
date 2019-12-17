// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeBoolean Core Datatype for boolean which accept JSON true and false
// values, but can also accept strings which are interpreted as either true or false:
//
// False values: false, "false"
// True values: true, "true"
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/boolean.html
// for details.
type DatatypeBoolean struct {
	Datatype
	name string

	// fields specific to boolean datatype
	boost     *float32
	docValues *bool
	index     *bool
	nullValue interface{}
	store     *bool
}

// NewDatatypeBoolean initializes a new DatatypeBoolean.
func NewDatatypeBoolean(name string) *DatatypeBoolean {
	return &DatatypeBoolean{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (b *DatatypeBoolean) Name() string {
	return b.name
}

//
// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (b *DatatypeBoolean) Boost(boost float32) *DatatypeBoolean {
	b.boost = &boost
	return b
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
//
// Defaults to true.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (b *DatatypeBoolean) DocValues(docValues bool) *DatatypeBoolean {
	b.docValues = &docValues
	return b
}

//
// Index sets whether if the field should be searchable. Defaults to true.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (b *DatatypeBoolean) Index(index bool) *DatatypeBoolean {
	b.index = &index
	return b
}

// NullValue sets any of the true and false values which is substituted for
//
// any explicit null values. Defaults to null.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (b *DatatypeBoolean) NullValue(nullValue interface{}) *DatatypeBoolean {
	b.nullValue = &nullValue
	return b
}

// Store sets whether if the field value should be stored and retrievable separately
//
// from the `_source` field. Defaults to false.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (b *DatatypeBoolean) Store(store bool) *DatatypeBoolean {
	b.store = &store
	return b
}

// Validate validates DatatypeBoolean.
func (b *DatatypeBoolean) Validate(includeName bool) error {
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
func (b *DatatypeBoolean) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "boolean",
	// 		"boost": 2,
	// 		"doc_values": true,
	// 		"index": true,
	// 		"null_value": "true", // false
	// 		"store": true,
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "boolean"

	if b.boost != nil {
		options["boost"] = b.boost
	}
	if b.docValues != nil {
		options["doc_values"] = b.docValues
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
