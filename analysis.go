// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// Analysis text analysis is the process of converting text, like the body of any email,
// into tokens or terms which are added to the inverted index for searching. Analysis is
// performed by an `analyzer` which can be either a built-in analyzer or a `custom` analyzer
// defined per index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis.html
// for details.
type Analysis struct {
	defaultAnalyzer Analyzer
	analyzer        []Analyzer
	normalizer      []Normalizer
	filter          []TokenFilter
	charFilter      []CharacterFilter
}

// NewAnalysis initializes a new Analysis.
func NewAnalysis() *Analysis {
	return &Analysis{}
}

// DefaultAnalyzer sets the default index-time analyzer.
func (a *Analysis) DefaultAnalyzer(defaultAnalyzer Analyzer) *Analysis {
	a.defaultAnalyzer = defaultAnalyzer
	return a
}

// Analyzer sets the analyzers for this index text analysis.
func (a *Analysis) Analyzer(analyzer ...Analyzer) *Analysis {
	a.analyzer = append(a.analyzer, analyzer...)
	return a
}

// Normalizer sets the normalizers for this index text analysis.
func (a *Analysis) Normalizer(normalizer ...Normalizer) *Analysis {
	a.normalizer = append(a.normalizer, normalizer...)
	return a
}

// Filter sets the token filters for this index text analysis.
func (a *Analysis) Filter(filter ...TokenFilter) *Analysis {
	a.filter = append(a.filter, filter...)
	return a
}

// CharFilter sets the character filters for this index text analysis.
func (a *Analysis) CharFilter(charFilter ...CharacterFilter) *Analysis {
	a.charFilter = append(a.charFilter, charFilter...)
	return a
}

// Validate validates Analysis.
func (a *Analysis) Validate() error {
	return nil
}

// Source returns the serializable JSON for the source builder.
func (a *Analysis) Source(includeName bool) (interface{}, error) {
	// {
	// 	"analysis": {
	// 		"analyzer": {
	// 			"default": {
	// 				"type": "whitespace"
	// 			},
	// 			"custom": {
	// 				"type": "custom",
	// 				"tokenizer": "standard"
	// 			}
	// 		},
	// 		"normalizer": {
	// 			"custom_normalizer": {
	// 				"type": "custom",
	// 				"char_filter": ["quote"]
	// 			}
	// 		},
	// 		"filter": {
	// 			"custom_synonym": {
	// 				"type": "synonym",
	// 				"synonyms": ["i-pod, i pod => ipod", "universe, cosmos"]
	// 			}
	// 		},
	// 		"char_filter": {
	// 			"custom_mapping": {
	// 				"type": "mapping",
	// 				"mappings": ["٠ => 0", "١ => 1", "٢ => 2"]
	// 			}
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})

	analyzers := make(map[string]interface{})
	if a.defaultAnalyzer != nil {
		defaultAnalyzer, err := a.defaultAnalyzer.Source(false)
		if err != nil {
			return nil, err
		}
		analyzers["default"] = defaultAnalyzer
	}
	if len(a.analyzer) > 0 {
		for _, _a := range a.analyzer {
			analyzer, err := _a.Source(false)
			if err != nil {
				return nil, err
			}
			analyzers[_a.Name()] = analyzer
		}
	}
	if len(analyzers) > 0 {
		options["analyzer"] = analyzers
	}
	if len(a.normalizer) > 0 {
		normalizers := make(map[string]interface{})
		for _, n := range a.normalizer {
			normalizer, err := n.Source(false)
			if err != nil {
				return nil, err
			}
			normalizers[n.Name()] = normalizer
		}
		options["normalizer"] = normalizers
	}
	if len(a.filter) > 0 {
		filters := make(map[string]interface{})
		for _, f := range a.filter {
			filter, err := f.Source(false)
			if err != nil {
				return nil, err
			}
			filters[f.Name()] = filter
		}
		options["filter"] = filters
	}
	if len(a.charFilter) > 0 {
		charFilters := make(map[string]interface{})
		for _, f := range a.charFilter {
			charFilter, err := f.Source(false)
			if err != nil {
				return nil, err
			}
			charFilters[f.Name()] = charFilter
		}
		options["char_filter"] = charFilters
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["analysis"] = options
	return source, nil
}
