// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterStemmerMappingRule mapping rule for Stemmer Token Filter.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-stemmer-override-tokenfilter.html
// for details.
type TokenFilterStemmerMappingRule struct {
	from string
	to   string
}

// NewTokenFilterStemmerMappingRule initializes a new TokenFilterStemmerMappingRule.
func NewTokenFilterStemmerMappingRule(from, to string) *TokenFilterStemmerMappingRule {
	return &TokenFilterStemmerMappingRule{
		from: from,
		to:   to,
	}
}

// From sets from key for custom mapping.
func (r *TokenFilterStemmerMappingRule) From(from string) *TokenFilterStemmerMappingRule {
	r.from = from
	return r
}

// To sets to key for custom mapping.
func (r *TokenFilterStemmerMappingRule) To(to string) *TokenFilterStemmerMappingRule {
	r.to = to
	return r
}

// Source returns the serializable JSON for the source builder.
func (r *TokenFilterStemmerMappingRule) Source() (interface{}, error) {
	// "from => to"
	source := fmt.Sprintf("%s => %s", r.from, r.to)

	return source, nil
}
