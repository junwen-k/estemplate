// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterSynonym token filter that allows to easily handle synonyms during
// the analysis process.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-synonym-tokenfilter.html
// for details.
type TokenFilterSynonym struct {
	TokenFilter
	name string

	// fields specific to synonym token filter
	synonyms     []*MappingRule
	rawSynonyms  []string
	synonymsPath string
	expand       *bool
	lenient      *bool
	format       string
	tokenizer    string
	ignoreCase   *bool
}

// NewTokenFilterSynonym initializes a new TokenFilterSynonym.
func NewTokenFilterSynonym(name string) *TokenFilterSynonym {
	return &TokenFilterSynonym{
		name:        name,
		synonyms:    make([]*MappingRule, 0),
		rawSynonyms: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (s *TokenFilterSynonym) Name() string {
	return s.name
}

// Synonyms sets a list of synonyms to be used for the filter.
func (s *TokenFilterSynonym) Synonyms(synonyms ...*MappingRule) *TokenFilterSynonym {
	s.synonyms = append(s.synonyms, synonyms...)
	return s
}

// RawSynonyms sets a list of synonyms to be used for the filter. Use this if you prefer to use
// strings instead of using MappingRule type.
func (s *TokenFilterSynonym) RawSynonyms(rawSynonyms ...string) *TokenFilterSynonym {
	s.rawSynonyms = append(s.rawSynonyms, rawSynonyms...)
	return s
}

// SynonymsPath sets the path to a file containing synonyms. This path must be absolute
// or relative to the `config` location. The file must be UTF-8 encoded. Each synonym in the
// file must be separated by a line break.
func (s *TokenFilterSynonym) SynonymsPath(synonymsPath string) *TokenFilterSynonym {
	s.synonymsPath = synonymsPath
	return s
}

// Expand sets whether to expand explicit mapping or not. Only affects synonyms with no
// explicit mappings.
// Defaults to true.
func (s *TokenFilterSynonym) Expand(expand bool) *TokenFilterSynonym {
	s.expand = &expand
	return s
}

// Lenient sets whether to ignore exceptions while parsing the synonym configuration, ignoring
// synonym rules which cannot get parsed.
// Defaults to false.
func (s *TokenFilterSynonym) Lenient(lenient bool) *TokenFilterSynonym {
	s.lenient = &lenient
	return s
}

// Format sets synonym format to be used.
// Can be set to the following values:
// "solr"
// "wordnet"
func (s *TokenFilterSynonym) Format(format string) *TokenFilterSynonym {
	s.format = format
	return s
}

// Tokenizer sets the tokenizer to be used to tokenize the synonym.
// ! Deprecated, this parameter is for backwards compatibility for indices
// created before 6.0.
func (s *TokenFilterSynonym) Tokenizer(tokenizer string) *TokenFilterSynonym {
	s.tokenizer = tokenizer
	return s
}

// IgnoreCase sets whether matches for tokenizer matching should be case-insensitive or not.
// ! Deprecated, this parameter only works with `tokenizer` parameter only.
func (s *TokenFilterSynonym) IgnoreCase(ignoreCase bool) *TokenFilterSynonym {
	s.ignoreCase = &ignoreCase
	return s
}

// Validate validates TokenFilterSynonym.
func (s *TokenFilterSynonym) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if s.format != "" {
		if _, valid := map[string]bool{
			"solr":    true,
			"wordnet": true,
		}[s.format]; !valid {
			invalid = append(invalid, "Format")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *TokenFilterSynonym) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "synonym",
	// 		"synonyms": ["i-pod, i pod => ipod", "universe, cosmos"],
	// 		"synonyms_path": "analysis/synonym.txt",
	// 		"expand": true,
	// 		"lenient": true,
	// 		"format": "solr",
	// 		"tokenizer": "standard",
	// 		"ignore_case": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "synonym"

	if len(s.synonyms) > 0 {
		var synonyms interface{}
		var ss []interface{}
		for _, syn := range s.synonyms {
			synonym, err := syn.Source()
			if err != nil {
				return nil, err
			}
			ss = append(ss, synonym)
		}
		switch {
		case len(ss) > 1:
			synonyms = ss
			break
		case len(ss) == 1:
			synonyms = ss[0]
			break
		default:
			synonyms = ""
		}
		options["synonyms"] = synonyms
	}
	if len(s.rawSynonyms) > 0 {
		var synonyms interface{}
		switch {
		case len(s.rawSynonyms) > 1:
			synonyms = s.rawSynonyms
			break
		case len(s.rawSynonyms) == 1:
			synonyms = s.rawSynonyms[0]
			break
		default:
			synonyms = ""
		}
		options["synonyms"] = synonyms
	}
	if s.synonymsPath != "" {
		options["synonyms_path"] = s.synonymsPath
	}
	if s.expand != nil {
		options["expand"] = s.expand
	}
	if s.lenient != nil {
		options["lenient"] = s.lenient
	}
	if s.format != "" {
		options["format"] = s.format
	}
	if s.tokenizer != "" {
		options["tokenizer"] = s.tokenizer
	}
	if s.ignoreCase != nil {
		options["ignore_case"] = s.ignoreCase
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
