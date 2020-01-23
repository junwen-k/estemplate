// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeRankFeatures Specialised Datatype that can index numeric feature vectors,
// so that they can later be used to boost documents in queries with a `rank_feature`
// query.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/rank-features.html
// for details.
type DatatypeRankFeatures struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to rank features datatype
}

// NewDatatypeRankFeatures initializes a new DatatypeRankFeatures.
func NewDatatypeRankFeatures(name string) *DatatypeRankFeatures {
	return &DatatypeRankFeatures{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (f *DatatypeRankFeatures) Name() string {
	return f.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (f *DatatypeRankFeatures) CopyTo(copyTo ...string) *DatatypeRankFeatures {
	f.copyTo = append(f.copyTo, copyTo...)
	return f
}

// Validate validates DatatypeRankFeatures.
func (f *DatatypeRankFeatures) Validate(includeName bool) error {
	var invalid []string
	if includeName && f.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (f *DatatypeRankFeatures) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "rank_features",
	// 		"copy_to": ["field_1", "field_2"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "rank_features"

	if len(f.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(f.copyTo) > 1:
			copyTo = f.copyTo
			break
		case len(f.copyTo) == 1:
			copyTo = f.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[f.name] = options
	return source, nil
}
