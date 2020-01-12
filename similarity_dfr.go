// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SimilarityDFR similarity that implements the divergence from randomness framework.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#dfr
// for details.
type SimilarityDFR struct {
	Similarity
	name string

	// field specific to dfr similarity
	basicModel    string
	afterEffect   string
	normalization string
}

// NewSimilarityDFR initializes a new SimilarityDFR.
func NewSimilarityDFR(name string) *SimilarityDFR {
	return &SimilarityDFR{
		name: name,
	}
}

// Name returns field key for the Similarity.
func (s *SimilarityDFR) Name() string {
	return s.name
}

// BasicModel sets the basic model for this similarity.
// Can be set to the following values:
// "g" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/BasicModelG.html
// "if" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/BasicModelIF.html
// "in" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/BasicModelIn.html
// "ine" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/BasicModelIne.html
func (s *SimilarityDFR) BasicModel(basicModel string) *SimilarityDFR {
	s.basicModel = basicModel
	return s
}

// AfterEffect sets the after effect for this similarity.
// Can be set to the following values:
// "b" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/AfterEffectB.html
// "l" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/AfterEffectB.html
func (s *SimilarityDFR) AfterEffect(afterEffect string) *SimilarityDFR {
	s.afterEffect = afterEffect
	return s
}

// Normalization sets the normalization for this similarity.
// Can be set to the following values:
// "no" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/Normalization.NoNormalization.html
// "h1" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationH1.html
// "h2" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationH2.html
// "h3" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationH3.html
// "z" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationZ.html
func (s *SimilarityDFR) Normalization(normalization string) *SimilarityDFR {
	s.normalization = normalization
	return s
}

// Validate validates SimilarityDFR.
func (s *SimilarityDFR) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if s.basicModel != "" {
		if _, valid := map[string]bool{
			"g":   true,
			"if":  true,
			"in":  true,
			"ine": true,
		}[s.basicModel]; !valid {
			invalid = append(invalid, "BasicModel")
		}
	}
	if s.afterEffect != "" {
		if _, valid := map[string]bool{
			"b": true,
			"l": true,
		}[s.afterEffect]; !valid {
			invalid = append(invalid, "AfterEffect")
		}
	}
	if s.normalization != "" {
		if _, valid := map[string]bool{
			"no": true,
			"h1": true,
			"h2": true,
			"h3": true,
			"z":  true,
		}[s.normalization]; !valid {
			invalid = append(invalid, "Normalization")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *SimilarityDFR) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "DFR",
	// 		"basic_model": "g",
	// 		"after_effect": "b",
	// 		"normalization": "no"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "DFR"

	if s.basicModel != "" {
		options["basic_model"] = s.basicModel
	}
	if s.afterEffect != "" {
		options["after_effect"] = s.afterEffect
	}
	if s.normalization != "" {
		options["normalization"] = s.normalization
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
