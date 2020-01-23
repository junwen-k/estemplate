// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeShape Specialised Datatype for arbitrary cartesian geometries. The
// Datatype facilitates the indexing of and searching with arbitrary `x`, `y`
// cartesian shapes such as rectangles and polygons.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/shape.html
// for details.
type DatatypeShape struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to shape datatype
	orientation     string
	ignoreMalformed *bool
	ignoreZValue    *bool
	coerce          *bool
}

// NewDatatypeShape initializes a new DatatypeShape.
func NewDatatypeShape(name string) *DatatypeShape {
	return &DatatypeShape{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (s *DatatypeShape) Name() string {
	return s.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (s *DatatypeShape) CopyTo(copyTo ...string) *DatatypeShape {
	s.copyTo = append(s.copyTo, copyTo...)
	return s
}

// Orientation sets the orientation on how to interpret vertex order for polygons / multipolygons.
// Defines one of two coordinate system rule, Right-hand or Left-hand.
// Can be set to the following values:
// "right", "ccw", "counterclockwise" - Right-hand rule.
// "left", "cw", "clockwise" - Left-hand rule.
// Defaults to "ccw".
func (s *DatatypeShape) Orientation(orientation string) *DatatypeShape {
	s.orientation = orientation
	return s
}

// IgnoreMalformed sets whether if the field should ignore malformed GeoJSON or WKT shapes.
// Defaults to false.
func (s *DatatypeShape) IgnoreMalformed(ignoreMalformed bool) *DatatypeShape {
	s.ignoreMalformed = &ignoreMalformed
	return s
}

// IgnoreZValue sets whether if the field should ignore the third dimension when three
// dimension points is received. Defaults to true.
func (s *DatatypeShape) IgnoreZValue(ignoreZValue bool) *DatatypeShape {
	s.ignoreZValue = &ignoreZValue
	return s
}

// Coerce sets whether if unclosed linear rings in polygons will be automatically closed.
// Defaults to false.
func (s *DatatypeShape) Coerce(coerce bool) *DatatypeShape {
	s.coerce = &coerce
	return s
}

// Validate validates DatatypeShape.
func (s *DatatypeShape) Validate(includeName bool) error {
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
func (s *DatatypeShape) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "shape",
	// 		"copy_to": ["field_1", "field_2"],
	// 		"orientation": "ccw",
	// 		"ignore_malformed": true,
	// 		"ignore_z_value": true,
	// 		"coerce": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "shape"

	if len(s.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(s.copyTo) > 1:
			copyTo = s.copyTo
			break
		case len(s.copyTo) == 1:
			copyTo = s.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if s.orientation != "" {
		options["orientation"] = s.orientation
	}
	if s.ignoreMalformed != nil {
		options["ignore_malformed"] = s.ignoreMalformed
	}
	if s.ignoreZValue != nil {
		options["ignore_z_value"] = s.ignoreZValue
	}
	if s.coerce != nil {
		options["coerce"] = s.coerce
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
