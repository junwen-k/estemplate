// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterHunspell token filter basic support for hunspell stemming.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-hunspell-tokenfilter.html
// for details.
type TokenFilterHunspell struct {
	TokenFilter
	name string

	// fields specific to hunspell token filter
	ignoreCase  *bool
	locale      string
	dictionary  string
	dedup       *bool
	longestOnly *bool
}

// NewTokenFilterHunspell initializes a new TokenFilterHunspell.
func NewTokenFilterHunspell(name string) *TokenFilterHunspell {
	return &TokenFilterHunspell{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (h *TokenFilterHunspell) Name() string {
	return h.name
}

// IgnoreCase sets whether dictionary matching should be case sensitive.
// Defaults to false.
func (h *TokenFilterHunspell) IgnoreCase(ignoreCase bool) *TokenFilterHunspell {
	h.ignoreCase = &ignoreCase
	return h
}

// Locale sets the locale for this filter. If this is unset, the `lang` or `language` are used
// instead, so one of these has to be set.
func (h *TokenFilterHunspell) Locale(locale string) *TokenFilterHunspell {
	h.locale = locale
	return h
}

// Dictionary sets the name of the dictionary. That path to your hunspell dictionaries should be
// configured via `indices.analysis.hunspell.dictionary.location` before.
func (h *TokenFilterHunspell) Dictionary(dictionary string) *TokenFilterHunspell {
	h.dictionary = dictionary
	return h
}

// Dedup sets whether if only unique terms should be returned, this needs to be set to true.
// Defaults to true.
func (h *TokenFilterHunspell) Dedup(dedup bool) *TokenFilterHunspell {
	h.dedup = &dedup
	return h
}

// LongestOnly sets whether if only longest term should be returned, set this to true.
// Defaults to false (all possible stems are returned).
func (h *TokenFilterHunspell) LongestOnly(longestOnly bool) *TokenFilterHunspell {
	h.longestOnly = &longestOnly
	return h
}

// Validate validates TokenFilterHunspell.
func (h *TokenFilterHunspell) Validate(includeName bool) error {
	var invalid []string
	if includeName && h.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (h *TokenFilterHunspell) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "hunspell",
	// 		"ignore_case": true,
	// 		"locale": "en_US",
	// 		"dictionary": "US_dictionary",
	// 		"dedup": true,
	// 		"longest_only": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "hunspell"

	if h.ignoreCase != nil {
		options["ignore_case"] = h.ignoreCase
	}
	if h.locale != "" {
		options["locale"] = h.locale
	}
	if h.dictionary != "" {
		options["dictionary"] = h.dictionary
	}
	if h.dedup != nil {
		options["dedup"] = h.dedup
	}
	if h.longestOnly != nil {
		options["longest_only"] = h.longestOnly
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[h.name] = options
	return source, nil
}
