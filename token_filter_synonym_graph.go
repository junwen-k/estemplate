// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterSynonymGraph token filter that allows to easily handle synonyms, including
// multi-word synonyms correctly the analysis process. This token filter is designed to be
// used as part of a search analyzer only. If you want to apply synonyms during indexing
// please use the standard `synonym` token filter.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-synonym-graph-tokenfilter.html
// for details.
type TokenFilterSynonymGraph struct {
	TokenFilter
	name string

	// fields specific to synonym token graph filter
	synonyms     []string
	synonymsPath string
	expand       *bool
	lenient      *bool
	format       string
	tokenizer    string
	ignoreCase   *bool
}

// NewTokenFilterSynonymGraph initializes a new TokenFilterSynonymGraph.
func NewTokenFilterSynonymGraph(name string) *TokenFilterSynonymGraph {
	return &TokenFilterSynonymGraph{
		name:     name,
		synonyms: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (g *TokenFilterSynonymGraph) Name() string {
	return g.name
}

// Synonyms sets a list of synonyms to be used for the filter.
func (g *TokenFilterSynonymGraph) Synonyms(synonyms ...string) *TokenFilterSynonymGraph {
	g.synonyms = append(g.synonyms, synonyms...)
	return g
}

// SynonymsPath sets the path to a file containing synonyms. This path must be absolute
// or relative to the `config` location. The file must be UTF-8 encoded. Each synonym in the
// file must be separated by a line break.
func (g *TokenFilterSynonymGraph) SynonymsPath(synonymsPath string) *TokenFilterSynonymGraph {
	g.synonymsPath = synonymsPath
	return g
}

// Expand sets whether to expand explicit mapping or not. Only affects synonyms with no
// explicit mappings.
// Defaults to true.
func (g *TokenFilterSynonymGraph) Expand(expand bool) *TokenFilterSynonymGraph {
	g.expand = &expand
	return g
}

// Lenient sets whether to ignore exceptions while parsing the synonym configuration, ignoring
// synonym rules which cannot get parsed.
// Defaults to false.
func (g *TokenFilterSynonymGraph) Lenient(lenient bool) *TokenFilterSynonymGraph {
	g.lenient = &lenient
	return g
}

// Format sets synonym format to be used.
// Can be set to the following values:
// "solr"
// "wordnet"
func (g *TokenFilterSynonymGraph) Format(format string) *TokenFilterSynonymGraph {
	g.format = format
	return g
}

// Tokenizer sets the tokenizer to be used to tokenize the synonym.
// ! Deprecated, this parameter is for backwards compatibility for indices
// created before 6.0.
func (g *TokenFilterSynonymGraph) Tokenizer(tokenizer string) *TokenFilterSynonymGraph {
	g.tokenizer = tokenizer
	return g
}

// IgnoreCase sets whether matches for tokenizer matching should be case-insensitive or not.
// ! Deprecated, this parameter only works with `tokenizer` parameter only.
func (g *TokenFilterSynonymGraph) IgnoreCase(ignoreCase bool) *TokenFilterSynonymGraph {
	g.ignoreCase = &ignoreCase
	return g
}

// Validate validates TokenFilterSynonymGraph.
func (g *TokenFilterSynonymGraph) Validate(includeName bool) error {
	var invalid []string
	if includeName && g.name == "" {
		invalid = append(invalid, "Name")
	}
	if g.format != "" {
		if _, valid := map[string]bool{
			"solr":    true,
			"wordnet": true,
		}[g.format]; !valid {
			invalid = append(invalid, "Format")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (g *TokenFilterSynonymGraph) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "synonym_graph",
	// 		"synonyms": ["lol, laughing out loud", "universe, cosmos"],
	// 		"synonyms_path": "analysis/synonym.txt",
	// 		"expand": true,
	// 		"lenient": true,
	// 		"format": "solr",
	// 		"tokenizer": "standard",
	// 		"ignore_case": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "synonym_graph"

	if len(g.synonyms) > 0 {
		var synonyms interface{}
		switch {
		case len(g.synonyms) > 1:
			synonyms = g.synonyms
			break
		case len(g.synonyms) == 1:
			synonyms = g.synonyms[0]
			break
		default:
			synonyms = ""
		}
		options["synonyms"] = synonyms
	}
	if g.synonymsPath != "" {
		options["synonyms_path"] = g.synonymsPath
	}
	if g.expand != nil {
		options["expand"] = g.expand
	}
	if g.lenient != nil {
		options["lenient"] = g.lenient
	}
	if g.format != "" {
		options["format"] = g.format
	}
	if g.tokenizer != "" {
		options["tokenizer"] = g.tokenizer
	}
	if g.ignoreCase != nil {
		options["ignore_case"] = g.ignoreCase
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[g.name] = options
	return source, nil
}
