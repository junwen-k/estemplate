// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterWordDelimiterGraph token filter that splits words into subwords and performs optional
// transformations on subword groups. Words are split into subwords with the following rules:
// - split on intra-word delimiters (by default, all non alpha-numeric characters): "Wi-Fi" → "Wi", "Fi"
// - split on case transitions: "PowerShot" → "Power", "Shot"
// - split on letter-number transitions: "SD500" → "SD", "500"
// - leading and trailing intra-word delimiters on each subword are ignored: "//hello---there, dude" → "hello", "there", "dude"
// - trailing "'s" are removed for each subword: "O’Neil’s" → "O", "Neil"
//
// Unlike the `word_delimiter`, this token filter correctly handles positions for multi terms expansion at search-time when any
// of the following options are set to true:
// - "preserve_original"
// - "catenate_numbers"
// - "catenate_words"
// - "catenate_all"
// ! Experimental in Lucene.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-word-delimiter-graph-tokenfilter.html
// for details.
type TokenFilterWordDelimiterGraph struct {
	TokenFilter
	name string

	// fields specific to word delimiter graph token filter
	generateWordParts     *bool
	generateNumberParts   *bool
	catenateWords         *bool
	catenateNumbers       *bool
	catenateAll           *bool
	splitOnCaseChange     *bool
	preserveOriginal      *bool
	splitOnNumerics       *bool
	stemEnglishPossessive *bool
	protectedWords        []string
	protectedWordsPath    string
	typeTable             []string
	typeTablePath         string
}

// NewTokenFilterWordDelimiterGraph initializes a new TokenFilterWordDelimiterGraph.
func NewTokenFilterWordDelimiterGraph(name string) *TokenFilterWordDelimiterGraph {
	return &TokenFilterWordDelimiterGraph{
		name:           name,
		protectedWords: make([]string, 0),
		typeTable:      make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (g *TokenFilterWordDelimiterGraph) Name() string {
	return g.name
}

// GenerateWordParts sets whether parts of words to be generated. For example:
// "Power-Shot", "(Power,Shot)" => "Power" "Shot".
// Defaults to true.
func (g *TokenFilterWordDelimiterGraph) GenerateWordParts(generateWordParts bool) *TokenFilterWordDelimiterGraph {
	g.generateWordParts = &generateWordParts
	return g
}

// GenerateNumberParts sets whether number subwords to be generated. For example:
// "500-42" => "500" "42".
// Defaults to true.
func (g *TokenFilterWordDelimiterGraph) GenerateNumberParts(generateNumberParts bool) *TokenFilterWordDelimiterGraph {
	g.generateNumberParts = &generateNumberParts
	return g
}

// CatenateWords sets whether maximum runs of word parts to be catenated. For example:
// "wi-fi" => "wifi".
// Defaults to false.
func (g *TokenFilterWordDelimiterGraph) CatenateWords(catenateWords bool) *TokenFilterWordDelimiterGraph {
	g.catenateWords = &catenateWords
	return g
}

// CatenateNumbers sets whether maximum runs of number parts to be catenated. For example:
// "500-42" => "50042".
// Defaults to false.
func (g *TokenFilterWordDelimiterGraph) CatenateNumbers(catenateNumbers bool) *TokenFilterWordDelimiterGraph {
	g.catenateNumbers = &catenateNumbers
	return g
}

// CatenateAll sets whether all subword parts to be catenated. For example:
// "wi-fi-4000" => "wifi4000".
// Defaults to false.
func (g *TokenFilterWordDelimiterGraph) CatenateAll(catenateAll bool) *TokenFilterWordDelimiterGraph {
	g.catenateAll = &catenateAll
	return g
}

// SplitOnCaseChange sets whether to split word on case change.
// Defaults to true.
func (g *TokenFilterWordDelimiterGraph) SplitOnCaseChange(splitOnCaseChange bool) *TokenFilterWordDelimiterGraph {
	g.splitOnCaseChange = &splitOnCaseChange
	return g
}

// PreserveOriginal sets whether to include original words in subwords. For example:
// "500-42" => "500-42" "500" "42".
// Defaults to false.
func (g *TokenFilterWordDelimiterGraph) PreserveOriginal(preserveOriginal bool) *TokenFilterWordDelimiterGraph {
	g.preserveOriginal = &preserveOriginal
	return g
}

// SplitOnNumerics sets whether to split on numerics part of a word into tokens.
// Defaults to true.
func (g *TokenFilterWordDelimiterGraph) SplitOnNumerics(splitOnNumerics bool) *TokenFilterWordDelimiterGraph {
	g.splitOnNumerics = &splitOnNumerics
	return g
}

// StemEnglishPossessive sets whether to remove trailing word for each subword.
// Defaults to true.
func (g *TokenFilterWordDelimiterGraph) StemEnglishPossessive(stemEnglishPossessive bool) *TokenFilterWordDelimiterGraph {
	g.stemEnglishPossessive = &stemEnglishPossessive
	return g
}

// ProtectedWords sets a list of protected words from being delimiter.
func (g *TokenFilterWordDelimiterGraph) ProtectedWords(protectedWords ...string) *TokenFilterWordDelimiterGraph {
	g.protectedWords = append(g.protectedWords, protectedWords...)
	return g
}

// ProtectedWordsPath sets a path to a file containing protected words.
// This path must be absolute or relative to the `config` location.
// The file must be UTF-8 encoded. Each word in the file must be separated
// by a line break.
func (g *TokenFilterWordDelimiterGraph) ProtectedWordsPath(protectedWordsPath string) *TokenFilterWordDelimiterGraph {
	g.protectedWordsPath = protectedWordsPath
	return g
}

// TypeTable sets a custom type mapping table. For example:
// $ => DIGIT
// \\u200D => ALPHANUM
func (g *TokenFilterWordDelimiterGraph) TypeTable(typeTable ...string) *TokenFilterWordDelimiterGraph {
	g.typeTable = append(g.typeTable, typeTable...)
	return g
}

// TypeTablePath sets a path to a file containing type table.
// This path must be absolute or relative to the `config` location.
// The file must be UTF-8 encoded. Each type map in the file must be separated
// by a line break.
func (g *TokenFilterWordDelimiterGraph) TypeTablePath(typeTablePath string) *TokenFilterWordDelimiterGraph {
	g.typeTablePath = typeTablePath
	return g
}

// Validate validates TokenFilterWordDelimiterGraph.
func (g *TokenFilterWordDelimiterGraph) Validate(includeName bool) error {
	var invalid []string
	if includeName && g.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (g *TokenFilterWordDelimiterGraph) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "word_delimiter_graph",
	// 		"generate_word_parts": true,
	// 		"generate_number_parts": true,
	// 		"catenate_words": true,
	// 		"catenate_numbers": true,
	// 		"catenate_all": true,
	// 		"split_on_case_change": true,
	// 		"preserve_original": true,
	// 		"split_on_numerics": true,
	// 		"stem_english_possessive": true,
	// 		"protected_words": ["hello", "world"],
	// 		"protected_words_path": "analysis/protected_words.txt",
	// 		"type_table": ["$ => DIGIT", "\\u200D => ALPHANUM"],
	// 		"type_table_path": "analysis/type_table.txt"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "word_delimiter_graph"

	if g.generateWordParts != nil {
		options["generate_word_parts"] = g.generateWordParts
	}
	if g.generateNumberParts != nil {
		options["generate_number_parts"] = g.generateNumberParts
	}
	if g.catenateWords != nil {
		options["catenate_words"] = g.catenateWords
	}
	if g.catenateNumbers != nil {
		options["catenate_numbers"] = g.catenateNumbers
	}
	if g.catenateAll != nil {
		options["catenate_all"] = g.catenateAll
	}
	if g.splitOnCaseChange != nil {
		options["split_on_case_change"] = g.splitOnCaseChange
	}
	if g.preserveOriginal != nil {
		options["preserve_original"] = g.preserveOriginal
	}
	if g.splitOnNumerics != nil {
		options["split_on_numerics"] = g.splitOnNumerics
	}
	if g.stemEnglishPossessive != nil {
		options["stem_english_possessive"] = g.stemEnglishPossessive
	}
	if len(g.protectedWords) > 0 {
		var protectedWords interface{}
		switch {
		case len(g.protectedWords) > 1:
			protectedWords = g.protectedWords
			break
		case len(g.protectedWords) == 1:
			protectedWords = g.protectedWords[0]
			break
		default:
			protectedWords = ""
		}
		options["protected_words"] = protectedWords
	}
	if g.protectedWordsPath != "" {
		options["protected_words_path"] = g.protectedWordsPath
	}
	if len(g.typeTable) > 0 {
		var typeTable interface{}
		switch {
		case len(g.typeTable) > 1:
			typeTable = g.typeTable
			break
		case len(g.typeTable) == 1:
			typeTable = g.typeTable[0]
			break
		default:
			typeTable = ""
		}
		options["type_table"] = typeTable
	}
	if g.typeTablePath != "" {
		options["type_table_path"] = g.typeTablePath
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[g.name] = options
	return source, nil
}
