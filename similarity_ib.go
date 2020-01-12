// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SimilarityIB similarity that uses information based model. The algorithm is based on the
// concept that the information content in any symbolic distribution sequence is primarily
// determined by the repetitive usage of its basic elements. For written texts this
// challenge would correspond to comparing the writing styles of different authors.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#ib
// for details.
type SimilarityIB struct {
	Similarity
	name string

	// fields specific to ib similarity
	distribution  string
	lambda        string
	normalization string
}

// NewSimilarityIB initializes a new SimilarityIB.
func NewSimilarityIB(name string) *SimilarityIB {
	return &SimilarityIB{
		name: name,
	}
}

// Name returns field key for the Similarity.
func (s *SimilarityIB) Name() string {
	return s.name
}

// Distribution sets the distribution for this similarity.
// Can be set to the following values:
// "ll" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/DistributionLL.html
// "spl" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/DistributionSPL.html
func (s *SimilarityIB) Distribution(distribution string) *SimilarityIB {
	s.distribution = distribution
	return s
}

// Lambda sets the lambda for this similarity.
// Can be set to the following values:
// "df" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/LambdaDF.html
// "ttf" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/LambdaTTF.html
func (s *SimilarityIB) Lambda(lambda string) *SimilarityIB {
	s.lambda = lambda
	return s
}

// Normalization sets the normalization for this similarity.
// Can be set to the following values:
// "no" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/Normalization.NoNormalization.html
// "h1" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationH1.html
// "h2" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationH2.html
// "h3" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationH3.html
// "z" - http://lucene.apache.org/core/8_3_0/core/org/apache/lucene/search/similarities/NormalizationZ.html
func (s *SimilarityIB) Normalization(normalization string) *SimilarityIB {
	s.normalization = normalization
	return s
}

// Validate validates SimilarityIB.
func (s *SimilarityIB) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if s.distribution != "" {
		if _, valid := map[string]bool{
			"ll":  true,
			"spl": true,
		}[s.distribution]; !valid {
			invalid = append(invalid, "Distribution")
		}
	}
	if s.lambda != "" {
		if _, valid := map[string]bool{
			"df":  true,
			"ttf": true,
		}[s.lambda]; !valid {
			invalid = append(invalid, "Lambda")
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
func (s *SimilarityIB) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "IB",
	// 		"distribution": "ll",
	// 		"lambda": "df",
	// 		"normalization": "no"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "IB"

	if s.distribution != "" {
		options["distribution"] = s.distribution
	}
	if s.lambda != "" {
		options["lambda"] = s.lambda
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
