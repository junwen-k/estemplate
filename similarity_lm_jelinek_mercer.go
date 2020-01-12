// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SimilarityLMJelinekMercer similarity that uses the LM Jelinek Mercer. The algorithm attempts to
// capture important patterns in the text, while leaving out noise.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#lm_jelinek_mercer
// for details.
type SimilarityLMJelinekMercer struct {
	Similarity
	name string

	// fields specific to lm jelinek mercer similarity
	lambda *float32
}

// NewSimilarityLMJelinekMercer initializes a new SimilarityLMJelinekMercer.
func NewSimilarityLMJelinekMercer(name string) *SimilarityLMJelinekMercer {
	return &SimilarityLMJelinekMercer{
		name: name,
	}
}

// Name returns field key for the Similarity.
func (s *SimilarityLMJelinekMercer) Name() string {
	return s.name
}

// Lambda sets the optimal value depends on both the collection and query. The optimal value is
// around 0.1 for title queries and 0.7 for long queries. When value approaches 0, documents that
// match more query terms will be ranked higher than those that match fewer terms.
// Defaults to 0.1.
func (s *SimilarityLMJelinekMercer) Lambda(lambda float32) *SimilarityLMJelinekMercer {
	s.lambda = &lambda
	return s
}

// Validate validates SimilarityLMJelinekMercer.
func (s *SimilarityLMJelinekMercer) Validate(includeName bool) error {
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
func (s *SimilarityLMJelinekMercer) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "LMJelinekMercer",
	// 		"lambda": 0.1
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "LMJelinekMercer"

	if s.lambda != nil {
		options["lambda"] = s.lambda
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
