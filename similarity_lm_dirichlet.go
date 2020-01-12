// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SimilarityLMDirichlet similarity that uses LM Dirichlet which assigns negative scores
// to terms that have fewer occurrences than predicted by the language model, which is illegal
// to Lucene, so such terms get a score of 0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#lm_dirichlet
// for details.
type SimilarityLMDirichlet struct {
	Similarity
	name string

	// fields specific to lm dirichlet similarity
	mu *int
}

// NewSimilarityLMDirichlet initializes a new SimilarityLMDirichlet.
func NewSimilarityLMDirichlet(name string) *SimilarityLMDirichlet {
	return &SimilarityLMDirichlet{
		name: name,
	}
}

// Name returns field key for the Similarity.
func (s *SimilarityLMDirichlet) Name() string {
	return s.name
}

// MU sets the mu for this similarity.
// Defaults to 2000.
func (s *SimilarityLMDirichlet) MU(mu int) *SimilarityLMDirichlet {
	s.mu = &mu
	return s
}

// Validate validates SimilarityLMDirichlet.
func (s *SimilarityLMDirichlet) Validate(includeName bool) error {
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
func (s *SimilarityLMDirichlet) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "LMDirichlet",
	// 		"mu": 2000
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "LMDirichlet"

	if s.mu != nil {
		options["mu"] = s.mu
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
