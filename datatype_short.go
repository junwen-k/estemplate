// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeShort Core Datatype for numeric value.
// A signed 16-bit integer with a minimum value of -32,768 and a maximum value of 32,767.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeShort struct {
	Datatype
	name string

	// fields specific to short datatype
	coerce          *bool
	boost           *int
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
}

// NewDatatypeShort initializes a new DatatypeShort.
func NewDatatypeShort(name string) *DatatypeShort {
	return &DatatypeShort{
		name: name,
	}
}

// Name is the key of the Short Property.
func (s *DatatypeShort) Name() string {
	return s.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (s *DatatypeShort) Coerce(coerce bool) *DatatypeShort {
	s.coerce = &coerce
	return s
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (s *DatatypeShort) Boost(boost int) *DatatypeShort {
	s.boost = &boost
	return s
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (s *DatatypeShort) DocValues(docValues bool) *DatatypeShort {
	s.docValues = &docValues
	return s
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defatuls to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (s *DatatypeShort) IgnoreMalformed(ignoreMalformed bool) *DatatypeShort {
	s.ignoreMalformed = &ignoreMalformed
	return s
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (s *DatatypeShort) Index(index bool) *DatatypeShort {
	s.index = &index
	return s
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (s *DatatypeShort) NullValue(nullValue int) *DatatypeShort {
	s.nullValue = &nullValue
	return s
}

// Store sets whether the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (s *DatatypeShort) Store(store bool) *DatatypeShort {
	s.store = &store
	return s
}

// Validate validates DatatypeShort.
func (s *DatatypeShort) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *DatatypeShort) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "short",
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
	options["type"] = "short"

	if s.coerce != nil {
		options["coerce"] = s.coerce
	}
	if s.boost != nil {
		options["boost"] = s.boost
	}
	if s.docValues != nil {
		options["doc_values"] = s.docValues
	}
	if s.ignoreMalformed != nil {
		options["ignore_malformed"] = s.ignoreMalformed
	}
	if s.index != nil {
		options["index"] = s.index
	}
	if s.nullValue != nil {
		options["null_value"] = s.nullValue
	}
	if s.store != nil {
		options["store"] = s.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
