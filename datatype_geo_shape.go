// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeGeoShape Geo Datatype facilitates the indexing of and searching with arbitrary
// geo shapes such as rectangles and polygons.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/geo-shape.html
// for details.
type DatatypeGeoShape struct {
	Datatype
	name string

	// fields specific to geo shape datatype
	tree             string
	precision        string
	treeLevels       string
	strategy         string
	distanceErrorPct *float32
	orientation      string
	pointsOnly       *bool
	ignoreMalformed  *bool
	ignoreZValue     *bool
	coerce           *bool
}

// NewDatatypeGeoShape initializes a new DatatypeGeoShape.
func NewDatatypeGeoShape(name string) *DatatypeGeoShape {
	return &DatatypeGeoShape{
		name: name,
	}
}

// TODO validations, option types and docs improvement

// Tree sets the name of the PrefixTree implementation to be used.
// Can be set to the following values:
// "geohash - GeohashPrefixTree.
// "quadtree" - QuadPrefixTree.
// * This parameter is only relevant for `term` and `recursive` strategies.
// Defaults to "quadtree".
// ! Deprecated in 6.6. PrefixTrees no longer used
func (s *DatatypeGeoShape) Tree(tree string) *DatatypeGeoShape {
	s.tree = tree
	return s
}

// Precision sets an appropriate value for the `tree_levels` parameter.
// The value specifies the desired precision and Elasticsearch will calculate the
// best tree_levels value to honor this precision.
// * This parameter is only relevant for `term` and `recursive` strategies.
// Defaults to "50m".
// ! Deprecated in 6.6. PrefixTrees no longer used
func (s *DatatypeGeoShape) Precision(precision string) *DatatypeGeoShape {
	s.precision = precision
	return s
}

// TreeLevels sets the maximum number of layers to be used by the PrefixTree.
// This can be used to control the precision of shape representations and therefore
// how many terms are indexed.
// * This parameter is only relevant for `term` and `recursive` strategies.
// Defaults to "various".
// ! Deprecated in 6.6. PrefixTrees no longer used
func (s *DatatypeGeoShape) TreeLevels(treeLevels string) *DatatypeGeoShape {
	s.treeLevels = treeLevels
	return s
}

// Strategy sets the approach for how to represent shapes at indexing and search time.
// Can be set to the following values:
// "recursive" - supports all shape types.
// "term" - supports `points_only` (the `points_only` parameter will be automatically set to true).
// ! Both strategies are deprecated and will be removed in a future version.
// Defaults to "recursive".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/geo-shape.html#prefix-trees
// for details.
func (s *DatatypeGeoShape) Strategy(strategy string) *DatatypeGeoShape {
	s.strategy = strategy
	return s
}

// DistanceErrorPct sets a hint to the PrefixTree about how precise it should be.
// * This parameter is only relevant for `term` and `recursive` strategies.
// Defaults to 0.025, or 0 (if `precision` and `tree_level` definition is explicitly defined).
// ! Deprecated in 6.6. PrefixTrees no longer used
func (s *DatatypeGeoShape) DistanceErrorPct(distanceErrorPct float32) *DatatypeGeoShape {
	s.distanceErrorPct = &distanceErrorPct
	return s
}

// Orientation sets the orientation on how to interpret vertex order for polygons / multipolygons.
// Defines one of two coordinate system rule, Right-hand or Left-hand.
// Can be set to the following values:
// "right", "ccw", "counterclockwise" - Right-hand rule.
// "left", "cw", "clockwise" - Left-hand rule.
// Defaults to "ccw".
func (s *DatatypeGeoShape) Orientation(orientation string) *DatatypeGeoShape {
	s.orientation = orientation
	return s
}

// PointsOnly sets true to optimize index and search performance for the `geohash` and `quadtree`
// when it is known that only points will be indexed, bridging the gap by improving point performance
// on a `geo_shape` field so that `geo_shape` queries are optimal on a point only field.
// Defaults to false.
func (s *DatatypeGeoShape) PointsOnly(pointsOnly bool) *DatatypeGeoShape {
	s.pointsOnly = &pointsOnly
	return s
}

// IgnoreMalformed sets whether if the field should ignore malformed GeoJSON or WKT shapes.
// Defaults to false.
func (s *DatatypeGeoShape) IgnoreMalformed(ignoreMalformed bool) *DatatypeGeoShape {
	s.ignoreMalformed = &ignoreMalformed
	return s
}

// IgnoreZValue sets whether if the field should ignore the third dimension when three
// dimension points is received. Defaults to true.
func (s *DatatypeGeoShape) IgnoreZValue(ignoreZValue bool) *DatatypeGeoShape {
	s.ignoreZValue = &ignoreZValue
	return s
}

// Coerce sets whether if unclosed linear rings in polygons will be automatically closed.
// Defaults to true.
func (s *DatatypeGeoShape) Coerce(coerce bool) *DatatypeGeoShape {
	s.coerce = &coerce
	return s
}

// Validate validates DatatypeGeoShape.
func (s *DatatypeGeoShape) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	// TODO validate precision prefixes
	// TODO validate distance error pct (0.5 max)
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *DatatypeGeoShape) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "geo_shape",
	// 		"tree": "quadtree",
	// 		"precision": "50m",
	// 		"tree_levels": "various",
	// 		"strategy": "recursive",
	// 		"distance_error_pct": 0.0,
	// 		"orientation": "ccw",
	// 		"points_only": true,
	// 		"ignore_malformed": true,
	// 		"ignore_z_value": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "geo_shape"

	if s.tree != "" {
		options["tree"] = s.tree
	}
	if s.precision != "" {
		options["precision"] = s.precision
	}
	if s.treeLevels != "" {
		options["tree_levels"] = s.treeLevels
	}
	if s.strategy != "" {
		options["strategy"] = s.strategy
	}
	if s.distanceErrorPct != nil {
		options["distance_error_pct"] = s.distanceErrorPct
	}
	if s.orientation != "" {
		options["orientation"] = s.orientation
	}
	if s.pointsOnly != nil {
		options["points_only"] = s.pointsOnly
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
