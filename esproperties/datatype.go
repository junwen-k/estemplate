// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package esproperties

import (
	"reflect"
	"strconv"
)

// Datatype is a datatype for Elasticsearch.
// The zero Datatype is not a valid Datatype.
type Datatype int

const (
	// Invalid Datatype
	Invalid Datatype = iota
	// Text Datatype
	Text
	// Keyword Datatype
	Keyword
	// Long Datatype
	Long
	// Integer Datatype
	Integer
	// Short Datatype
	Short
	// Byte Datatype
	Byte
	// Double Datatype
	Double
	// Float Datatype
	Float
	// HalfFloat Datatype
	HalfFloat
	// ScaledFloat Datatype
	ScaledFloat
	// Date Datatype
	Date
	// DateNanoseconds Datatype
	DateNanoseconds
	// Boolean Datatype
	Boolean
	// Binary Datatype
	Binary
	// IntegerRange Datatype
	IntegerRange
	// FloatRange Datatype
	FloatRange
	// LongRange Datatype
	LongRange
	// DoubleRange Datatype
	DoubleRange
	// DateRange Datatype
	DateRange
	// Object Datatype
	Object
	// Nested Datatype
	Nested
	// GeoPoint Datatype
	GeoPoint
	// GeoShape Datatype
	GeoShape
	// IP Datatype
	IP
	// Completion Datatype
	Completion
	// TokenCount Datatype
	TokenCount
	// MapperMurmur3 Datatype
	MapperMurmur3
	// MapperAnnotatedText Datatype
	MapperAnnotatedText
	// Percolator Datatype
	Percolator
	// Join Datatype
	Join
	// RankFeature Datatype
	RankFeature
	// RankFeatures Datatype
	RankFeatures
	// DenseVector Datatype
	DenseVector
	// SparseVector Datatype
	SparseVector
	// SearchAsYouType Datatype
	SearchAsYouType
	// Alias Datatype
	Alias
	// Flattened Datatype
	Flattened
	// Shape Datatype
	Shape
)

// Decode decode datatype from string value.
func (d *Datatype) Decode(value string) error {
	switch value {
	case "text":
		*d = Text
	case "keyword":
		*d = Keyword
	case "long":
		*d = Long
	case "integer":
		*d = Integer
	case "short":
		*d = Short
	case "byte":
		*d = Byte
	case "double":
		*d = Double
	case "float":
		*d = Float
	case "half_float":
		*d = HalfFloat
	case "scaled_float":
		*d = ScaledFloat
	case "date":
		*d = Date
	case "date_nanoseconds":
		*d = DateNanoseconds
	case "boolean":
		*d = Boolean
	case "binary":
		*d = Binary
	case "integer_range":
		*d = IntegerRange
	case "float_range":
		*d = FloatRange
	case "long_range":
		*d = LongRange
	case "double_range":
		*d = DoubleRange
	case "date_range":
		*d = DateRange
	case "object":
		*d = Object
	case "nested":
		*d = Nested
	case "geo_point":
		*d = GeoPoint
	case "geo_shape":
		*d = GeoShape
	case "ip":
		*d = IP
	case "completion":
		*d = Completion
	case "token_count":
		*d = TokenCount
	case "murmur3":
		*d = MapperMurmur3
	case "annotated_text":
		*d = MapperAnnotatedText
	case "percolator":
		*d = Percolator
	case "join":
		*d = Join
	case "rank_feature":
		*d = RankFeature
	case "rank_features":
		*d = RankFeatures
	case "dense_vector":
		*d = DenseVector
	case "sparse_vector":
		*d = SparseVector
	case "search_as_you_type":
		*d = SearchAsYouType
	case "alias":
		*d = Alias
	case "flattened":
		*d = Flattened
	case "shape":
		*d = Shape
	default:
		*d = Invalid
	}
	return nil
}

// DecodeByKind decode datatype by reflect.Kind value.
func (d *Datatype) DecodeByKind(k reflect.Kind) error {
	switch k {
	case reflect.Bool:
		*d = Boolean
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		*d = Integer
	case reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.Complex64, reflect.Complex128:
		*d = Float
	case reflect.Array, reflect.Slice:
		*d = Nested
	case reflect.String:
		*d = Text
	case reflect.Map, reflect.Struct:
		*d = Object
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer, reflect.Ptr, reflect.Invalid:
	default:
		*d = Invalid
	}
	return nil
}

// String returns the name of d.
func (d *Datatype) String() string {
	if int(*d) < len(datatypeNames) {
		return datatypeNames[*d]
	}
	return "datatype" + strconv.Itoa(int(*d))
}

var datatypeNames = []string{
	Invalid:             "invalid",
	Text:                "text",
	Keyword:             "keyword",
	Long:                "long",
	Integer:             "integer",
	Short:               "short",
	Byte:                "byte",
	Double:              "double",
	Float:               "float",
	HalfFloat:           "half_float",
	ScaledFloat:         "scaled_float",
	Date:                "date",
	DateNanoseconds:     "date_nanoseconds",
	Boolean:             "boolean",
	Binary:              "binary",
	IntegerRange:        "integer_range",
	FloatRange:          "float_range",
	LongRange:           "long_range",
	DoubleRange:         "double_range",
	DateRange:           "date_range",
	Object:              "object",
	Nested:              "nested",
	GeoPoint:            "geo_point",
	GeoShape:            "geo_shape",
	IP:                  "ip",
	Completion:          "completion",
	TokenCount:          "token_count",
	MapperMurmur3:       "mapper_murmur3",
	MapperAnnotatedText: "mapper_annotated_text",
	Percolator:          "percolator",
	Join:                "join",
	RankFeature:         "rank_feature",
	RankFeatures:        "rank_features",
	DenseVector:         "dense_vector",
	SparseVector:        "sparse_vector",
	SearchAsYouType:     "search_as_you_type",
	Alias:               "alias",
	Flattened:           "flattened",
	Shape:               "shape",
}
