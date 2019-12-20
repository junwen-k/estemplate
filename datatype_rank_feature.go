// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeRankFeature Specialised Datatype that index numbers so that they can later
// be used to boost documents in queries with a `rank_feature` query.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/rank-feature.html
// for details.
type DatatypeRankFeature struct {
	Datatype
	name string

	// fields specific to rank feature datatype
	positiveScoreImpact *bool
}

// NewDatatypeRankFeature initializes a new DatatypeRankFeature.
func NewDatatypeRankFeature(name string) *DatatypeRankFeature {
	return &DatatypeRankFeature{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (f *DatatypeRankFeature) Name() string {
	return f.name
}

// PositiveScoreImpact sets and allows `rank_feature` query to modify the scoring formula in such a way
// that the score decreases with the value of the feature instead of increasing.
// Defaults to true.
func (f *DatatypeRankFeature) PositiveScoreImpact(positiveScoreImpact bool) *DatatypeRankFeature {
	f.positiveScoreImpact = &positiveScoreImpact
	return f
}

// Validate validates DatatypeRankFeature.
func (f *DatatypeRankFeature) Validate(includeName bool) error {
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
func (f *DatatypeRankFeature) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "rank_feature",
	// 		"positive_score_impact": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "rank_feature"

	if f.positiveScoreImpact != nil {
		options["positive_score_impact"] = f.positiveScoreImpact
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[f.name] = options
	return source, nil
}
