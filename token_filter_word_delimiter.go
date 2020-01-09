// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterWordDelimiter token filter that splits words into subwords and performs optional
// transformations on subword groups. Words are split into subwords with the following rules:
// - split on intra-word delimiters (by default, all non alpha-numeric characters): "Wi-Fi" → "Wi", "Fi"
// - split on case transitions: "PowerShot" → "Power", "Shot"
// - split on letter-number transitions: "SD500" → "SD", "500"
// - leading and trailing intra-word delimiters on each subword are ignored: "//hello---there, dude" → "hello", "there", "dude"
// - trailing "'s" are removed for each subword: "O’Neil’s" → "O", "Neil"
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-word-delimiter-tokenfilter.html
// for details.
type TokenFilterWordDelimiter struct {
	TokenFilter
	name string

	// fields specific to word delimiter token filter
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

// NewTokenFilterWordDelimiter initializes a new TokenFilterWordDelimiter.
func NewTokenFilterWordDelimiter(name string) *TokenFilterWordDelimiter {
	return &TokenFilterWordDelimiter{
		name:           name,
		protectedWords: make([]string, 0),
		typeTable:      make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (d *TokenFilterWordDelimiter) Name() string {
	return d.name
}

// GenerateWordParts sets whether parts of words to be generated. For example:
// "Power-Shot", "(Power,Shot)" => "Power" "Shot".
// Defaults to true.
func (d *TokenFilterWordDelimiter) GenerateWordParts(generateWordParts bool) *TokenFilterWordDelimiter {
	d.generateWordParts = &generateWordParts
	return d
}

// GenerateNumberParts sets whether number subwords to be generated. For example:
// "500-42" => "500" "42".
// Defaults to true.
func (d *TokenFilterWordDelimiter) GenerateNumberParts(generateNumberParts bool) *TokenFilterWordDelimiter {
	d.generateNumberParts = &generateNumberParts
	return d
}

// CatenateWords sets whether maximum runs of word parts to be catenated. For example:
// "wi-fi" => "wifi".
// Defaults to false.
func (d *TokenFilterWordDelimiter) CatenateWords(catenateWords bool) *TokenFilterWordDelimiter {
	d.catenateWords = &catenateWords
	return d
}

// CatenateNumbers sets whether maximum runs of number parts to be catenated. For example:
// "500-42" => "50042".
// Defaults to false.
func (d *TokenFilterWordDelimiter) CatenateNumbers(catenateNumbers bool) *TokenFilterWordDelimiter {
	d.catenateNumbers = &catenateNumbers
	return d
}

// CatenateAll sets whether all subword parts to be catenated. For example:
// "wi-fi-4000" => "wifi4000".
// Defaults to false.
func (d *TokenFilterWordDelimiter) CatenateAll(catenateAll bool) *TokenFilterWordDelimiter {
	d.catenateAll = &catenateAll
	return d
}

// SplitOnCaseChange sets whether to split word on case change.
// Defaults to true.
func (d *TokenFilterWordDelimiter) SplitOnCaseChange(splitOnCaseChange bool) *TokenFilterWordDelimiter {
	d.splitOnCaseChange = &splitOnCaseChange
	return d
}

// PreserveOriginal sets whether to include original words in subwords. For example:
// "500-42" => "500-42" "500" "42".
// Defaults to false.
func (d *TokenFilterWordDelimiter) PreserveOriginal(preserveOriginal bool) *TokenFilterWordDelimiter {
	d.preserveOriginal = &preserveOriginal
	return d
}

// SplitOnNumerics sets whether to split on numerics part of a word into tokens.
// Defaults to true.
func (d *TokenFilterWordDelimiter) SplitOnNumerics(splitOnNumerics bool) *TokenFilterWordDelimiter {
	d.splitOnNumerics = &splitOnNumerics
	return d
}

// StemEnglishPossessive sets whether to remove trailing word for each subword.
// Defaults to true.
func (d *TokenFilterWordDelimiter) StemEnglishPossessive(stemEnglishPossessive bool) *TokenFilterWordDelimiter {
	d.stemEnglishPossessive = &stemEnglishPossessive
	return d
}

// ProtectedWords sets a list of protected words from being delimiter.
func (d *TokenFilterWordDelimiter) ProtectedWords(protectedWords ...string) *TokenFilterWordDelimiter {
	d.protectedWords = append(d.protectedWords, protectedWords...)
	return d
}

// ProtectedWordsPath sets a path to a file containing protected words.
// This path must be absolute or relative to the `config` location.
// The file must be UTF-8 encoded. Each word in the file must be separated
// by a line break.
func (d *TokenFilterWordDelimiter) ProtectedWordsPath(protectedWordsPath string) *TokenFilterWordDelimiter {
	d.protectedWordsPath = protectedWordsPath
	return d
}

// TypeTable sets a custom type mapping table. For example:
// $ => DIGIT
// \\u200D => ALPHANUM
func (d *TokenFilterWordDelimiter) TypeTable(typeTable ...string) *TokenFilterWordDelimiter {
	d.typeTable = append(d.typeTable, typeTable...)
	return d
}

// TypeTablePath sets a path to a file containing type table.
// This path must be absolute or relative to the `config` location.
// The file must be UTF-8 encoded. Each type map in the file must be separated
// by a line break.
func (d *TokenFilterWordDelimiter) TypeTablePath(typeTablePath string) *TokenFilterWordDelimiter {
	d.typeTablePath = typeTablePath
	return d
}

// Validate validates TokenFilterWordDelimiter.
func (d *TokenFilterWordDelimiter) Validate(includeName bool) error {
	var invalid []string
	if includeName && d.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (d *TokenFilterWordDelimiter) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "word_delimiter",
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
	options["type"] = "word_delimiter"

	if d.generateWordParts != nil {
		options["generate_word_parts"] = d.generateWordParts
	}
	if d.generateNumberParts != nil {
		options["generate_number_parts"] = d.generateNumberParts
	}
	if d.catenateWords != nil {
		options["catenate_words"] = d.catenateWords
	}
	if d.catenateNumbers != nil {
		options["catenate_numbers"] = d.catenateNumbers
	}
	if d.catenateAll != nil {
		options["catenate_all"] = d.catenateAll
	}
	if d.splitOnCaseChange != nil {
		options["split_on_case_change"] = d.splitOnCaseChange
	}
	if d.preserveOriginal != nil {
		options["preserve_original"] = d.preserveOriginal
	}
	if d.splitOnNumerics != nil {
		options["split_on_numerics"] = d.splitOnNumerics
	}
	if d.stemEnglishPossessive != nil {
		options["stem_english_possessive"] = d.stemEnglishPossessive
	}
	if len(d.protectedWords) > 0 {
		var protectedWords interface{}
		switch {
		case len(d.protectedWords) > 1:
			protectedWords = d.protectedWords
			break
		case len(d.protectedWords) == 1:
			protectedWords = d.protectedWords[0]
			break
		default:
			protectedWords = ""
		}
		options["protected_words"] = protectedWords
	}
	if d.protectedWordsPath != "" {
		options["protected_words_path"] = d.protectedWordsPath
	}
	if len(d.typeTable) > 0 {
		var typeTable interface{}
		switch {
		case len(d.typeTable) > 1:
			typeTable = d.typeTable
			break
		case len(d.typeTable) == 1:
			typeTable = d.typeTable[0]
			break
		default:
			typeTable = ""
		}
		options["type_table"] = typeTable
	}
	if d.typeTablePath != "" {
		options["type_table_path"] = d.typeTablePath
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[d.name] = options
	return source, nil
}
