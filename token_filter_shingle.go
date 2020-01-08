// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterShingle token filter that constructs shingles (token n-grams) from a
// token stream. In other words, it creates combinations of tokens as a single token.
// For example, the sentence "please divide this sentence into shingles" might be
// tokenized into shingles "please divide", "divide this", "this sentence", "sentence into",
// and "into shingles".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-shingle-tokenfilter.html
// for details.
type TokenFilterShingle struct {
	TokenFilter
	name string

	// fields specific to shingle token filter
	maxShingleSize             *int
	minShingleSize             *int
	outputUnigrams             *bool
	outputUnigramsIfNoShingles *bool
	tokenSeparator             string
	filterToken                string
}

// NewTokenFilterShingle initializes a new TokenFilterShingle.
func NewTokenFilterShingle(name string) *TokenFilterShingle {
	return &TokenFilterShingle{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (s *TokenFilterShingle) Name() string {
	return s.name
}

// MaxShingleSize sets the maximum shingle size.
// Defaults to 2.
func (s *TokenFilterShingle) MaxShingleSize(maxShingleSize int) *TokenFilterShingle {
	s.maxShingleSize = &maxShingleSize
	return s
}

// MinShingleSize sets the minimum shingle size.
// Defaults to 2.
func (s *TokenFilterShingle) MinShingleSize(minShingleSize int) *TokenFilterShingle {
	s.minShingleSize = &minShingleSize
	return s
}

// OutputUnigrams sets whether output will contain the input tokens (unigrams) as well
// as the shingles or not.
// Defaults to true.
func (s *TokenFilterShingle) OutputUnigrams(outputUnigrams bool) *TokenFilterShingle {
	s.outputUnigrams = &outputUnigrams
	return s
}

// OutputUnigramsIfNoShingles sets whether the output will contain the input tokens (unigrams)
// if no shingles are available. Note if `output_unigrams` is set to true, this setting has no
// effect.
// Defaults to false.
func (s *TokenFilterShingle) OutputUnigramsIfNoShingles(outputUnigramsIfNoShingles bool) *TokenFilterShingle {
	s.outputUnigramsIfNoShingles = &outputUnigramsIfNoShingles
	return s
}

// TokenSeparator sets the string to use when joining adjacent tokens to form a shingle.
// Defaults to " " (space).
func (s *TokenFilterShingle) TokenSeparator(tokenSeparator string) *TokenFilterShingle {
	s.tokenSeparator = tokenSeparator
	return s
}

// FilterToken sets the string to use as a replacement for each position at which there is
// no actual token in the stream. For instance this string is used if the position increment is
// greater than one when a `stop` filter is used together with the `shingle` filter.
// Defaults to "_".
func (s *TokenFilterShingle) FilterToken(filterToken string) *TokenFilterShingle {
	s.filterToken = filterToken
	return s
}

// Validate validates TokenFilterShingle.
func (s *TokenFilterShingle) Validate(includeName bool) error {
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
func (s *TokenFilterShingle) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "shingle",
	// 		"max_shingle_size": 2,
	// 		"min_shingle_size": 2,
	// 		"output_unigrams": true,
	// 		"output_unigrams_if_no_shingles": false,
	// 		"token_separator": "/",
	// 		"filter_token": "_"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "shingle"

	if s.maxShingleSize != nil {
		options["max_shingle_size"] = s.maxShingleSize
	}
	if s.minShingleSize != nil {
		options["min_shingle_size"] = s.minShingleSize
	}
	if s.outputUnigrams != nil {
		options["output_unigrams"] = s.outputUnigrams
	}
	if s.outputUnigramsIfNoShingles != nil {
		options["output_unigrams_if_no_shingles"] = s.outputUnigramsIfNoShingles
	}
	if s.tokenSeparator != "" {
		options["token_separator"] = s.tokenSeparator
	}
	if s.filterToken != "" {
		options["filter_token"] = s.filterToken
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
