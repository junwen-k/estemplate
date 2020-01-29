// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package esproperties

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/junwen-k/estemplate"
)

type nestedTest struct {
	InnerNested []nestedTest `es:"inner_nested,nested"`
}

func TestESPropertiesNewSerialization(t *testing.T) {
	tests := []struct {
		builder     Builder
		desc        string
		expected    string
		nestedLimit int
		origin      interface{}
	}{
		// #0
		{
			builder:     DefaultBuilder,
			desc:        "Text Datatype test",
			expected:    `{"string":{"type":"text"},"string_slice":{"type":"text"}}`,
			nestedLimit: 1,
			origin: struct {
				String      string   `es:"string,text"`
				StringSlice []string `es:"string_slice,text"`
			}{},
		},
		// #1
		{
			builder:     DefaultBuilder,
			desc:        "Keyword Datatype test",
			expected:    `{"string":{"type":"keyword"},"string_slice":{"type":"keyword"}}`,
			nestedLimit: 1,
			origin: struct {
				String      string   `es:"string,keyword"`
				StringSlice []string `es:"string_slice,keyword"`
			}{},
		},
		// #2
		{
			builder:     DefaultBuilder,
			desc:        "Integer Datatype test",
			expected:    `{"integer":{"type":"integer"},"integer_slice":{"type":"integer"}}`,
			nestedLimit: 1,
			origin: struct {
				Integer      int   `es:"integer,integer"`
				IntegerSlice []int `es:"integer_slice,integer"`
			}{},
		},
		// #3
		{
			builder:     DefaultBuilder,
			desc:        "Long Datatype test",
			expected:    `{"integer":{"type":"long"},"integer_slice":{"type":"long"}}`,
			nestedLimit: 1,
			origin: struct {
				Integer      int   `es:"integer,long"`
				IntegerSlice []int `es:"integer_slice,long"`
			}{},
		},
		// #4
		{
			builder:     DefaultBuilder,
			desc:        "Integer Datatype test",
			expected:    `{"integer":{"type":"integer"},"integer_slice":{"type":"integer"}}`,
			nestedLimit: 1,
			origin: struct {
				Integer      int   `es:"integer,integer"`
				IntegerSlice []int `es:"integer_slice,integer"`
			}{},
		},
		// #5
		{
			builder:     DefaultBuilder,
			desc:        "Short Datatype test",
			expected:    `{"integer":{"type":"short"},"integer_slice":{"type":"short"}}`,
			nestedLimit: 1,
			origin: struct {
				Integer      int   `es:"integer,short"`
				IntegerSlice []int `es:"integer_slice,short"`
			}{},
		},
		// #6
		{
			builder:     DefaultBuilder,
			desc:        "Byte Datatype test",
			expected:    `{"integer":{"type":"byte"},"integer_slice":{"type":"byte"}}`,
			nestedLimit: 1,
			origin: struct {
				Integer      int   `es:"integer,byte"`
				IntegerSlice []int `es:"integer_slice,byte"`
			}{},
		},
		// #7
		{
			builder:     DefaultBuilder,
			desc:        "Double Datatype test",
			expected:    `{"integer":{"type":"double"},"integer_slice":{"type":"double"}}`,
			nestedLimit: 1,
			origin: struct {
				Integer      int   `es:"integer,double"`
				IntegerSlice []int `es:"integer_slice,double"`
			}{},
		},
		// #8
		{
			builder:     DefaultBuilder,
			desc:        "Float Datatype test",
			expected:    `{"float":{"type":"float"},"float_slice":{"type":"float"}}`,
			nestedLimit: 1,
			origin: struct {
				Float      float32   `es:"float,float"`
				FloatSlice []float32 `es:"float_slice,float"`
			}{},
		},
		// #9
		{
			builder:     DefaultBuilder,
			desc:        "HalfFloat Datatype test",
			expected:    `{"float":{"type":"half_float"},"float_slice":{"type":"half_float"}}`,
			nestedLimit: 1,
			origin: struct {
				Float      float32   `es:"float,half_float"`
				FloatSlice []float32 `es:"float_slice,half_float"`
			}{},
		},
		// #10
		{
			builder:     DefaultBuilder,
			desc:        "ScaledFloat Datatype test",
			expected:    `{"float":{"type":"scaled_float"},"float_slice":{"type":"scaled_float"}}`,
			nestedLimit: 1,
			origin: struct {
				Float      float32   `es:"float,scaled_float"`
				FloatSlice []float32 `es:"float_slice,scaled_float"`
			}{},
		},
		// #11
		{
			builder:     DefaultBuilder,
			desc:        "Date Datatype test",
			expected:    `{"date":{"type":"date"}}`,
			nestedLimit: 1,
			origin: struct {
				Date *time.Time `es:"date,date"`
			}{},
		},
		// #12
		{
			builder:     DefaultBuilder,
			desc:        "DateNanoseconds Datatype test",
			expected:    `{"date":{"type":"date_nanos"}}`,
			nestedLimit: 1,
			origin: struct {
				Date *time.Time `es:"date,date_nanoseconds"`
			}{},
		},
		// #13
		{
			builder:     DefaultBuilder,
			desc:        "Boolean Datatype test",
			expected:    `{"boolean":{"type":"boolean"},"boolean_slice":{"type":"boolean"}}`,
			nestedLimit: 1,
			origin: struct {
				Boolean      bool   `es:"boolean,boolean"`
				BooleanSlice []bool `es:"boolean_slice,boolean"`
			}{},
		},
		// #14
		{
			builder:     DefaultBuilder,
			desc:        "Binary Datatype test",
			expected:    `{"binary":{"type":"binary"},"binary_slice":{"type":"binary"}}`,
			nestedLimit: 1,
			origin: struct {
				Binary      string   `es:"binary,binary"`
				BinarySlice []string `es:"binary_slice,binary"`
			}{},
		},
		// #15
		{
			builder:     DefaultBuilder,
			desc:        "IntegerRange Datatype test",
			expected:    `{"integer_range":{"type":"integer_range"}}`,
			nestedLimit: 1,
			origin: struct {
				IntegerRange struct {
					Gte string
					Lte string
				} `es:"integer_range,integer_range"`
			}{},
		},
		// #16
		{
			builder:     DefaultBuilder,
			desc:        "FloatRange Datatype test",
			expected:    `{"float_range":{"type":"float_range"}}`,
			nestedLimit: 1,
			origin: struct {
				FloatRange struct {
					Gte string
					Lte string
				} `es:"float_range,float_range"`
			}{},
		},
		// #17
		{
			builder:     DefaultBuilder,
			desc:        "LongRange Datatype test",
			expected:    `{"long_range":{"type":"long_range"}}`,
			nestedLimit: 1,
			origin: struct {
				LongRange struct {
					Gte string
					Lte string
				} `es:"long_range,long_range"`
			}{},
		},
		// #18
		{
			builder:     DefaultBuilder,
			desc:        "DoubleRange Datatype test",
			expected:    `{"double_range":{"type":"double_range"}}`,
			nestedLimit: 1,
			origin: struct {
				DoubleRange struct {
					Gte string
					Lte string
				} `es:"double_range,double_range"`
			}{},
		},
		// #19
		{
			builder:     DefaultBuilder,
			desc:        "DateRange Datatype test",
			expected:    `{"date_range":{"type":"date_range"}}`,
			nestedLimit: 1,
			origin: struct {
				DateRange struct {
					Gte *time.Time
					Lte *time.Time
				} `es:"date_range,date_range"`
			}{},
		},
		// #20
		{
			builder:     DefaultBuilder,
			desc:        "Object Datatype test",
			expected:    `{"object":{"type":"object"}}`,
			nestedLimit: 1,
			origin: struct {
				Object struct {
					String  string
					Integer int
				} `es:"object,object"`
			}{},
		},
		// #21
		{
			builder:     DefaultBuilder,
			desc:        "Nested Datatype test",
			expected:    `{"nested":{"type":"nested"}}`,
			nestedLimit: 1,
			origin: struct {
				Nested []struct {
					String  string
					Integer int
				} `es:"nested,nested"`
			}{},
		},
		// #22
		{
			builder:     DefaultBuilder,
			desc:        "GeoPoint Datatype test",
			expected:    `{"location":{"type":"geo_point"}}`,
			nestedLimit: 1,
			origin: struct {
				Location string `es:"location,geo_point"`
			}{},
		},
		// #23
		{
			builder:     DefaultBuilder,
			desc:        "GeoShape Datatype test",
			expected:    `{"shape":{"type":"geo_shape"}}`,
			nestedLimit: 1,
			origin: struct {
				Shape string `es:"shape,geo_shape"`
			}{},
		},
		// #24
		{
			builder:     DefaultBuilder,
			desc:        "IP Datatype test",
			expected:    `{"ip":{"type":"ip"}}`,
			nestedLimit: 1,
			origin: struct {
				IP string `es:"ip,ip"`
			}{},
		},
		// #25
		{
			builder:     DefaultBuilder,
			desc:        "Completion Datatype test",
			expected:    `{"completion":{"type":"completion"}}`,
			nestedLimit: 1,
			origin: struct {
				Completion string `es:"completion,completion"`
			}{},
		},
		// #26
		{
			builder:     DefaultBuilder,
			desc:        "TokenCount Datatype test",
			expected:    `{"token_count":{"type":"token_count"}}`,
			nestedLimit: 1,
			origin: struct {
				TokenCount string `es:"token_count,token_count"`
			}{},
		},
		// #27
		{
			builder:     DefaultBuilder,
			desc:        "MapperMurMur3 Datatype test",
			expected:    `{"murmur3":{"type":"murmur3"}}`,
			nestedLimit: 1,
			origin: struct {
				MurMur3 string `es:"murmur3,murmur3"`
			}{},
		},
		// #28
		{
			builder:     DefaultBuilder,
			desc:        "MapperAnnotatedText Datatype test",
			expected:    `{"annotated_text":{"type":"annotated_text"}}`,
			nestedLimit: 1,
			origin: struct {
				AnnotatedText string `es:"annotated_text,annotated_text"`
			}{},
		},
		// #29
		{
			builder:     DefaultBuilder,
			desc:        "Percolator Datatype test",
			expected:    `{"percolator":{"type":"percolator"}}`,
			nestedLimit: 1,
			origin: struct {
				Percolator string `es:"percolator,percolator"`
			}{},
		},
		// #30
		{
			builder:     DefaultBuilder,
			desc:        "Join Datatype test",
			expected:    `{"join":{"type":"join"}}`,
			nestedLimit: 1,
			origin: struct {
				Join string `es:"join,join"`
			}{},
		},
		// #31
		{
			builder:     DefaultBuilder,
			desc:        "RankFeature Datatype test",
			expected:    `{"rank_feature":{"type":"rank_feature"}}`,
			nestedLimit: 1,
			origin: struct {
				RankFeature string `es:"rank_feature,rank_feature"`
			}{},
		},
		// #32
		{
			builder:     DefaultBuilder,
			desc:        "RankFeatures Datatype test",
			expected:    `{"rank_features":{"type":"rank_features"}}`,
			nestedLimit: 1,
			origin: struct {
				RankFeatures []string `es:"rank_features,rank_features"`
			}{},
		},
		// #33
		{
			builder:     DefaultBuilder,
			desc:        "DenseVector Datatype test",
			expected:    `{"dense_vector":{"type":"dense_vector"}}`,
			nestedLimit: 1,
			origin: struct {
				DenseVector string `es:"dense_vector,dense_vector"`
			}{},
		},
		// #34
		{
			builder:     DefaultBuilder,
			desc:        "SparseVector Datatype test",
			expected:    `{"sparse_vector":{"type":"sparse_vector"}}`,
			nestedLimit: 1,
			origin: struct {
				SparseVector string `es:"sparse_vector,sparse_vector"`
			}{},
		},
		// #35
		{
			builder:     DefaultBuilder,
			desc:        "SearchAsYouType Datatype test",
			expected:    `{"match":{"type":"search_as_you_type"}}`,
			nestedLimit: 1,
			origin: struct {
				Match string `es:"match,search_as_you_type"`
			}{},
		},
		// #36
		{
			builder:     DefaultBuilder,
			desc:        "Alias Datatype test",
			expected:    `{"alias":{"type":"alias"}}`,
			nestedLimit: 1,
			origin: struct {
				Alias string `es:"alias,alias"`
			}{},
		},
		// #37
		{
			builder:     DefaultBuilder,
			desc:        "Flattened Datatype test",
			expected:    `{"flattened":{"type":"flattened"}}`,
			nestedLimit: 1,
			origin: struct {
				Flattened string `es:"flattened,flattened"`
			}{},
		},
		// #38
		{
			builder:     DefaultBuilder,
			desc:        "Shape Datatype test",
			expected:    `{"shape":{"type":"shape"}}`,
			nestedLimit: 1,
			origin: struct {
				Shape string `es:"shape,shape"`
			}{},
		},
		// #39
		{
			builder:     DefaultBuilder,
			desc:        "Dynamic Datatype test",
			expected:    `{"array":{"type":"nested"},"boolean":{"type":"boolean"},"complex_128":{"type":"float"},"complex_64":{"type":"float"},"float_32":{"type":"float"},"float_64":{"type":"float"},"integer":{"type":"integer"},"integer_16":{"type":"integer"},"integer_32":{"type":"integer"},"integer_64":{"type":"integer"},"integer_8":{"type":"integer"},"slice":{"type":"nested"},"string":{"type":"text"},"struct":{"type":"object"},"uinteger":{"type":"integer"},"uinteger_16":{"type":"integer"},"uinteger_32":{"type":"integer"},"uinteger_64":{"type":"integer"},"uinteger_8":{"type":"integer"},"uinteger_pointer":{"type":"integer"}}`,
			nestedLimit: 1,
			origin: struct {
				Boolean         bool                   `es:"boolean"`
				Integer         int                    `es:"integer"`
				Integer8        int8                   `es:"integer_8"`
				Integer16       int16                  `es:"integer_16"`
				Integer32       int32                  `es:"integer_32"`
				Integer64       int64                  `es:"integer_64"`
				UInteger        uint                   `es:"uinteger"`
				UInteger8       uint8                  `es:"uinteger_8"`
				UInteger16      uint16                 `es:"uinteger_16"`
				UInteger32      uint32                 `es:"uinteger_32"`
				UInteger64      uint64                 `es:"uinteger_64"`
				UIntegerPointer *uint                  `es:"uinteger_pointer"`
				Float32         float32                `es:"float_32"`
				Float64         float64                `es:"float_64"`
				Complex64       complex64              `es:"complex_64"`
				Complex128      complex128             `es:"complex_128"`
				Array           []struct{}             `es:"array"`
				Slice           []struct{}             `es:"slice"`
				String          string                 `es:"string"`
				Map             map[string]interface{} `es:"map"`
				Struct          struct{}               `es:"struct"`
			}{},
		},
		// #40
		{
			builder:     DefaultBuilder,
			desc:        "Nested Limit Datatype test",
			expected:    `{"nested":{"properties":{"inner_nested":{"properties":{"inner_nested":{"type":"nested"}},"type":"nested"}},"type":"nested"}}`,
			nestedLimit: 3,
			origin: struct {
				Nested nestedTest `es:"nested,nested"`
			}{},
		},
		// #41
		{
			builder: func(name string, nestedCount int, dt Datatype, datatype estemplate.Datatype) estemplate.Datatype {
				switch dt {
				case Text:
					return datatype.(*estemplate.DatatypeText).Fields(estemplate.NewDatatypeKeyword("keyword"))
				}
				return datatype
			},
			desc:        "Custom builder Datatype test",
			expected:    `{"string":{"fields":{"keyword":{"type":"keyword"}},"type":"text"}}`,
			nestedLimit: 1,
			origin: struct {
				String string `es:"string,text"`
			}{},
		},
		// #42
		{
			builder: func(name string, nestedCount int, dt Datatype, datatype estemplate.Datatype) estemplate.Datatype {
				if nestedCount == 1 && name == "inner_nested" {
					return datatype.(*estemplate.DatatypeNested).Dynamic(true)
				}
				return datatype
			},
			desc:        "Custom builder Datatype test",
			expected:    `{"nested":{"properties":{"inner_nested":{"dynamic":true,"type":"nested"}},"type":"nested"}}`,
			nestedLimit: 2,
			origin: struct {
				Nested nestedTest `es:"nested,nested"`
			}{},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			datatypes, err := New(test.origin, test.nestedLimit, test.builder)
			if err != nil {
				t.Fatal(err)
			}
			src, err := ToProperties(datatypes)
			if err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(src)
			if err != nil {
				t.Fatalf("marshaling to JSON failed: %v", err)
			}
			if got, expected := string(data), test.expected; got != expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
