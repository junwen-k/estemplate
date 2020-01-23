// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeGeoPoint Geo Datatype for latitude-longitude pairs, which can be used in:
// - to find geo-points within a bounding box, within a certain distance of a central point, or within a polygon.
// - to aggregate documents geographically or by distance from a central point.
// - to integrate distance into a document’s relevance score.
// - to sort documents by distance.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/geo-point.html
// for details.
type DatatypeGeoPoint struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to geo point datatype
	ignoreMalformed *bool
	ignoreZValue    *bool
	nullValue       interface{}
}

// NewDatatypeGeoPoint initializes a new DatatypeGeoPoint.
func NewDatatypeGeoPoint(name string) *DatatypeGeoPoint {
	return &DatatypeGeoPoint{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (p *DatatypeGeoPoint) Name() string {
	return p.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (p *DatatypeGeoPoint) CopyTo(copyTo ...string) *DatatypeGeoPoint {
	p.copyTo = append(p.copyTo, copyTo...)
	return p
}

// IgnoreMalformed sets whether if the field should ignore malformed geo-points.
// Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (p *DatatypeGeoPoint) IgnoreMalformed(ignoreMalformed bool) *DatatypeGeoPoint {
	p.ignoreMalformed = &ignoreMalformed
	return p
}

// IgnoreZValue sets whether if the field should ignore the third dimension when three
// dimension points is received. Defaults to true.
func (p *DatatypeGeoPoint) IgnoreZValue(ignoreZValue bool) *DatatypeGeoPoint {
	p.ignoreZValue = &ignoreZValue
	return p
}

// NullValue sets a geopoint value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (p *DatatypeGeoPoint) NullValue(nullValue interface{}) *DatatypeGeoPoint {
	p.nullValue = nullValue
	return p
}

// Validate validates DatatypeGeoPoint.
func (p *DatatypeGeoPoint) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *DatatypeGeoPoint) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "geo_point",
	// 		"copy_to": ["field_1", "field_2"],
	// 		"ignore_malformed": true,
	// 		"ignore_z_value": true,
	// 		"null_value": [ 0, 0 ]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "geo_point"

	if len(p.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(p.copyTo) > 1:
			copyTo = p.copyTo
			break
		case len(p.copyTo) == 1:
			copyTo = p.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if p.ignoreMalformed != nil {
		options["ignore_malformed"] = p.ignoreMalformed
	}
	if p.ignoreZValue != nil {
		options["ignore_z_value"] = p.ignoreZValue
	}
	if p.nullValue != nil {
		options["null_value"] = p.nullValue
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
