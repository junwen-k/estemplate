// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// SimilarityScripted similarity that allows you to use a script in order to specify how scores should
// be computed.
// ! While scripted similarities provides a lot of flexibility, there is a set of rules that they need
// to satisfy. Failing to do so could make Elasticsearch silently return wrong top hits or fail with
// internal errors at search time.
// - Returned scores must be positive.
// - All other variables remaining equal, scores must not decrease when `doc.freq` increases.
// - All other variables remaining equal, scores must not increase when `doc.length` increases.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#scripted_similarity
// for details.
type SimilarityScripted struct {
	Similarity
	name string

	// fields specific to scripted similarity
	weightScript *Script
	script       *Script
}

// NewSimilarityScripted initializes a new SimilarityScripted.
func NewSimilarityScripted(name string) *SimilarityScripted {
	return &SimilarityScripted{
		name: name,
	}
}

// Name returns field key for the Similarity.
func (s *SimilarityScripted) Name() string {
	return s.name
}

// WeightScript sets a script which will compute the document-independent part of the score
// and will be available under the `weight` variable. When no `weight_script` is provided,
// `weight` is equal to 1. The `weight_script` has access to the same variables as the `script`
// except `doc` since it is supposed to compute a document-indepent contribution to the score.
func (s *SimilarityScripted) WeightScript(weightScript *Script) *SimilarityScripted {
	s.weightScript = weightScript
	return s
}

// Script sets the script for this similarity.
func (s *SimilarityScripted) Script(script *Script) *SimilarityScripted {
	s.script = script
	return s
}

// Validate validates SimilarityScripted.
func (s *SimilarityScripted) Validate(includeName bool) error {
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
func (s *SimilarityScripted) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "scripted",
	// 		"weight_script": {
	// 			"source": "double idf = Math.log((field.docCount+1.0)/(term.docFreq+1.0)) + 1.0; return query.boost * idf;"
	// 		},
	// 		"script": {
	// 			"source": "double tf = Math.sqrt(doc.freq); double norm = 1/Math.sqrt(doc.length); return weight * tf * norm;"
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "scripted"

	if s.weightScript != nil {
		weightScript, err := s.weightScript.Source(false)
		if err != nil {
			return nil, err
		}
		options["weight_script"] = weightScript
	}
	if s.script != nil {
		script, err := s.script.Source(false)
		if err != nil {
			return nil, err
		}
		options["script"] = script
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
