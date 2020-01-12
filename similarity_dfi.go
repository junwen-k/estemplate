// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SimilarityDFI similarity that implements the divergence from independence model. It is highly
// recommended to remove stop words to get good relevance. Also beware that terms whose frequency
// is less than the expected frequency will get a score equal to 0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#dfi
// for details.
type SimilarityDFI struct {
	Similarity
	name string

	// fields specific to dfi similarity
	independenceMeasure string
}

// NewSimilarityDFI initializes a new SimilarityDFI.
func NewSimilarityDFI(name string) *SimilarityDFI {
	return &SimilarityDFI{
		name: name,
	}
}

// Name returns field key for the Similarity.
func (s *SimilarityDFI) Name() string {
	return s.name
}

// IndependenceMeasure sets the independence measure for this simlarity.
// Can be set to the following values:
// "standardized" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/IndependenceStandardized.html
// "saturated"- http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/IndependenceSaturated.html
// "chisquared" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/IndependenceChiSquared.html
func (s *SimilarityDFI) IndependenceMeasure(independenceMeasure string) *SimilarityDFI {
	s.independenceMeasure = independenceMeasure
	return s
}

// Validate validates SimilarityDFI.
func (s *SimilarityDFI) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if s.independenceMeasure != "" {
		if _, valid := map[string]bool{
			"standardized": true,
			"saturated":    true,
			"chisquared":   true,
		}[s.independenceMeasure]; !valid {
			invalid = append(invalid, "IndependenceMeasure")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *SimilarityDFI) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "DFI",
	// 		"independence_measure": "standardized"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "DFI"

	if s.independenceMeasure != "" {
		options["independence_measure"] = s.independenceMeasure
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
