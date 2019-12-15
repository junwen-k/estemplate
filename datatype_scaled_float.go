// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeScaledFloat Core Datatype for numeric value.
// A floating point number that is backed by a long, scaled by a fixed double scaling factor.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/number.html
// for details.
type DatatypeScaledFloat struct {
	Datatype
	name string

	// fields specific to scaled float datatype
	coerce          *bool
	boost           *int
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	nullValue       *int
	store           *bool
	scalingFactor   *int
}

// NewDatatypeScaledFloat initializes a new DatatypeScaledFloat.
func NewDatatypeScaledFloat(name string) *DatatypeScaledFloat {
	return &DatatypeScaledFloat{
		name: name,
	}
}

// Name is the key of the Scaled Float Property.
func (sf *DatatypeScaledFloat) Name() string {
	return sf.name
}

// Coerce sets whether if the field should be coerced, attempting to clean up
// dirty values to fit the datatype.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/coerce.html
// for details.
func (sf *DatatypeScaledFloat) Coerce(coerce bool) *DatatypeScaledFloat {
	sf.coerce = &coerce
	return sf
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (sf *DatatypeScaledFloat) Boost(boost int) *DatatypeScaledFloat {
	sf.boost = &boost
	return sf
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (sf *DatatypeScaledFloat) DocValues(docValues bool) *DatatypeScaledFloat {
	sf.docValues = &docValues
	return sf
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defatuls to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (sf *DatatypeScaledFloat) IgnoreMalformed(ignoreMalformed bool) *DatatypeScaledFloat {
	sf.ignoreMalformed = &ignoreMalformed
	return sf
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (sf *DatatypeScaledFloat) Index(index bool) *DatatypeScaledFloat {
	sf.index = &index
	return sf
}

// NullValue sets a numeric value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (sf *DatatypeScaledFloat) NullValue(nullValue int) *DatatypeScaledFloat {
	sf.nullValue = &nullValue
	return sf
}

// Store sets whether the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (sf *DatatypeScaledFloat) Store(store bool) *DatatypeScaledFloat {
	sf.store = &store
	return sf
}

// ScalingFactor sets the scaling factor to use when encoding values.
func (sf *DatatypeScaledFloat) ScalingFactor(scalingFactor int) *DatatypeScaledFloat {
	sf.scalingFactor = &scalingFactor
	return sf
}

// Validate validates DatatypeScaledFloat.
func (sf *DatatypeScaledFloat) Validate(includeName bool) error {
	var invalid []string
	if includeName && sf.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (sf *DatatypeScaledFloat) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "scaled_float",
	// 		"coerce": true,
	// 		"boost": 2,
	// 		"doc_values": true,
	// 		"ignore_malformed": true,
	// 		"index": true,
	// 		"null_value": 0,
	// 		"store": true,
	// 		"scaling_factor": 2
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "scaled_float"

	if sf.coerce != nil {
		options["coerce"] = sf.coerce
	}
	if sf.boost != nil {
		options["boost"] = sf.boost
	}
	if sf.docValues != nil {
		options["doc_values"] = sf.docValues
	}
	if sf.ignoreMalformed != nil {
		options["ignore_malformed"] = sf.ignoreMalformed
	}
	if sf.index != nil {
		options["index"] = sf.index
	}
	if sf.nullValue != nil {
		options["null_value"] = sf.nullValue
	}
	if sf.store != nil {
		options["store"] = sf.store
	}
	if sf.scalingFactor != nil {
		options["scaling_factor"] = sf.scalingFactor
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[sf.name] = options
	return source, nil
}
