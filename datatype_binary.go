// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeBinary Core Datatype for binary which accepts a binary value as
// a Base64 encoded string. The field is not stored by default and is not searchable.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/binary.html
// for details.
type DatatypeBinary struct {
	Datatype
	name string

	// fields specific to binary datatype
	docValues *bool
	store     *bool
}

// NewDatatypeBinary initializes a new DatatypeBinary.
func NewDatatypeBinary(name string) *DatatypeBinary {
	return &DatatypeBinary{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (b *DatatypeBinary) Name() string {
	return b.name
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
//
// Defaults to true.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (b *DatatypeBinary) DocValues(docValues bool) *DatatypeBinary {
	b.docValues = &docValues
	return b
}

// Store sets whether if the field value should be stored and retrievable separately
//
// from the `_source` field. Defaults to false.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (b *DatatypeBinary) Store(store bool) *DatatypeBinary {
	b.store = &store
	return b
}

// Validate validates DatatypeBinary.
func (b *DatatypeBinary) Validate(includeName bool) error {
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
func (b *DatatypeBinary) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "binary",
	// 		"doc_values": true,
	// 		"store": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "binary"

	if b.docValues != nil {
		options["doc_values"] = b.docValues
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
