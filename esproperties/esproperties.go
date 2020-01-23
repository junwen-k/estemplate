// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package esproperties

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/junwen_k/estemplate"
)

const (
	// TagName Elasticsearch properties tag name constant.
	TagName = "es"
)

// Builder custom builder function for each datatype based on name, current nested count (starts from 0),
// datatype enum and default drafted datatype.
type Builder func(name string, nestedCount int, dt Datatype, datatype estemplate.Datatype) estemplate.Datatype

// DefaultBuilder returns default datatype without any customization.
var DefaultBuilder = func(name string, nestedCount int, dt Datatype, datatype estemplate.Datatype) estemplate.Datatype {
	return datatype
}

// New generates Elasticsearch mapping properties from origin's struct tags.
func New(origin interface{}, nestedLimit int, builder Builder) ([]estemplate.Datatype, error) {
	t, err := getStructFromOrigin(origin)
	if err != nil {
		return nil, err
	}

	return generateDatatypes(t, 0, nestedLimit, builder)
}

// ToProperties convenience function that converts datatypes to properties.
// Example:
//z
// src, _ := esproperties.ToProperties(estemplate.NewDatatypeNested("array"))
// data, _ := json.Marshal(src)
// fmt.Println(data) // {"array":{"type":"nested"}}
func ToProperties(datatypes []estemplate.Datatype) (interface{}, error) {
	properties := make(map[string]interface{})
	if len(datatypes) > 0 {
		for _, datatype := range datatypes {
			property, err := datatype.Source(false)
			if err != nil {
				return nil, err
			}
			properties[datatype.Name()] = property
		}
	}

	return properties, nil
}

// getTypeElem returns element's type.
func getTypeElem(t reflect.Type, isNested bool) (reflect.Type, bool) {
	kind := t.Kind()
	if kind == reflect.Slice {
		switch t.Elem().Kind() {
		case reflect.Struct:
			isNested = true
		}
	}
	if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Map || kind == reflect.Ptr || kind == reflect.Slice {
		return getTypeElem(t.Elem(), isNested)
	}
	return t, isNested
}

// getValueElem returns value of interface contains / pointer points to.
func getValueElem(v reflect.Value) reflect.Value {
	kind := v.Kind()
	if kind == reflect.Interface || kind == reflect.Ptr {
		return getValueElem(v.Elem())
	}
	return v
}

// getStructFromOrigin returns struct type from origin.
func getStructFromOrigin(origin interface{}) (reflect.Type, error) {
	value := getValueElem(reflect.ValueOf(origin))
	kind := value.Kind()
	switch kind {
	case reflect.Struct, reflect.Slice:
		t, _ := getTypeElem(value.Type(), false)
		return t, nil
	default:
		return reflect.TypeOf(origin), fmt.Errorf("requires struct or slice; got %s", kind)
	}
}

// splitESTag splits "es" tag value into name and datatype.
func splitESTag(rawTags string) (string, string) {
	tags := strings.Split(rawTags, ",")
	if len(tags) > 1 {
		return strings.TrimSpace(tags[0]), strings.TrimSpace(tags[1])
	}
	return strings.TrimSpace(tags[0]), ""
}

// generateDatatypes generate Datatypes for properties.
func generateDatatypes(t reflect.Type, nestedCount, nestedLimit int, builder Builder) ([]estemplate.Datatype, error) {
	kind := t.Kind()
	if kind != reflect.Struct {
		return nil, fmt.Errorf("requires struct; got %s", kind)
	}

	datatypes := make([]estemplate.Datatype, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		esTag := field.Tag.Get(TagName)
		if esTag == "" || esTag == "-" {
			continue
		}

		name, _dt := splitESTag(esTag)
		if name == "" {
			name = field.Name
		}

		var (
			dt  Datatype
			err error
		)
		if _dt == "" {
			t, isNested := getTypeElem(field.Type, false)
			if isNested {
				err = dt.DecodeByKind(reflect.Slice)
			} else {
				err = dt.DecodeByKind(t.Kind())
			}
		} else {
			err = dt.Decode(_dt)
		}
		if err != nil {
			return nil, err
		}

		var datatype estemplate.Datatype
		switch dt {
		case Text:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeText(name))
		case Keyword:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeKeyword(name))
		case Long:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeLong(name))
		case Integer:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeInteger(name))
		case Short:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeShort(name))
		case Byte:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeByte(name))
		case Double:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeDouble(name))
		case Float:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeFloat(name))
		case HalfFloat:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeHalfFloat(name))
		case ScaledFloat:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeScaledFloat(name))
		case Date:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeDate(name))
		case DateNanoseconds:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeDateNanoseconds(name))
		case Boolean:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeBoolean(name))
		case Binary:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeBinary(name))
		case IntegerRange:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeIntegerRange(name))
		case FloatRange:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeFloatRange(name))
		case LongRange:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeLongRange(name))
		case DoubleRange:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeDoubleRange(name))
		case DateRange:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeDateRange(name))
		case Object:
			if nestedCount < nestedLimit {
				t, _ := getTypeElem(t.Field(i).Type, false)
				nestedDatatypes, err := generateDatatypes(t, nestedCount+1, nestedLimit, builder)
				if err != nil {
					return nil, err
				}
				datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeObject(name).Properties(nestedDatatypes...))
			}
		case Nested:
			if nestedCount < nestedLimit {
				t, _ := getTypeElem(t.Field(i).Type, false)
				nestedDatatypes, err := generateDatatypes(t, nestedCount+1, nestedLimit, builder)
				if err != nil {
					return nil, err
				}
				datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeNested(name).Properties(nestedDatatypes...))
			}
		case GeoPoint:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeGeoPoint(name))
		case GeoShape:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeGeoShape(name))
		case IP:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeIP(name))
		case Completion:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeCompletion(name))
		case TokenCount:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeTokenCount(name))
		case MapperMurmur3:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeMapperMurmur3(name))
		case MapperAnnotatedText:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeMapperAnnotatedText(name))
		case Percolator:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypePercolator(name))
		case Join:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeJoin(name))
		case RankFeature:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeRankFeature(name))
		case RankFeatures:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeRankFeatures(name))
		case DenseVector:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeDenseVector(name))
		case SparseVector:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeSparseVector(name))
		case SearchAsYouType:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeSearchAsYouType(name))
		case Alias:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeAlias(name))
		case Flattened:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeFlattened(name))
		case Shape:
			datatype = builder(name, nestedCount, dt, estemplate.NewDatatypeShape(name))
		case Invalid:
		default:
			return nil, fmt.Errorf("Undefined Datatype '%s' for field '%s'", t, field.Name)
		}

		if datatype != nil {
			datatypes = append(datatypes, datatype)
		}
	}

	return datatypes, nil
}
