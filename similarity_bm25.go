// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SimilarityBM25 similarity that is TF/IDF based which has built-in tf normalization and is supposed to
// work better for short fields (like names).
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#bm25
// for details.
type SimilarityBM25 struct {
	Similarity
	name string

	// fields specific to bm25 similarity
	k1               *float32
	b                *float32
	discountOverlaps *bool
}

// NewSimilarityBM25 initializes a new SimilarityBM25.
func NewSimilarityBM25(name string) *SimilarityBM25 {
	return &SimilarityBM25{
		name: name,
	}
}

// Name returns field key for the Similarity.
func (s *SimilarityBM25) Name() string {
	return s.name
}

// K1 sets the non-linear term frequency normalization (saturation).
// Defaults to 1.2.
func (s *SimilarityBM25) K1(k1 float32) *SimilarityBM25 {
	s.k1 = &k1
	return s
}

// B sets the degree document length normalizes tf values.
// Defaults to 0.75.
func (s *SimilarityBM25) B(b float32) *SimilarityBM25 {
	s.b = &b
	return s
}

// DiscountOverlaps sets whether or not overlap tokens (Tokens with 0 position increment) are ignored
// when computing norm. If set to true, overlap tokens do not count when computing norms.
// Defaults to true.
func (s *SimilarityBM25) DiscountOverlaps(discountOverlaps bool) *SimilarityBM25 {
	s.discountOverlaps = &discountOverlaps
	return s
}

// Validate validates SimilarityBM25.
func (s *SimilarityBM25) Validate(includeName bool) error {
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
func (s *SimilarityBM25) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "BM25",
	// 		"k1": 1.2,
	// 		"b": 0.75,
	// 		"discount_overlaps": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "BM25"

	if s.k1 != nil {
		options["k1"] = s.k1
	}
	if s.b != nil {
		options["b"] = s.b
	}
	if s.discountOverlaps != nil {
		options["discount_overlaps"] = s.discountOverlaps
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
