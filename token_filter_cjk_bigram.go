// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterCJKBigram token filter that forms bigrams out of CJK (Chinese, Japanese, and Korean) tokens.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-cjk-bigram-tokenfilter.html
// for details.
type TokenFilterCJKBigram struct {
	TokenFilter
	name string

	// fields specific to cjk bigram token filter
	ignoredScripts []string
	outputUnigrams *bool
}

// NewTokenFilterCJKBigram initializes a new TokenFilterCJKBigram.
func NewTokenFilterCJKBigram(name string) *TokenFilterCJKBigram {
	return &TokenFilterCJKBigram{
		name:           name,
		ignoredScripts: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (b *TokenFilterCJKBigram) Name() string {
	return b.name
}

// IgnoredScripts sets an array of character scripts for which to disable bigrams.
// Can be set to the following values:
// "han"
// "hangul"
// "hiragana"
// "katakana"
// All non-CJK input is passed through unmodified.
func (b *TokenFilterCJKBigram) IgnoredScripts(ignoredScripts ...string) *TokenFilterCJKBigram {
	b.ignoredScripts = append(b.ignoredScripts, ignoredScripts...)
	return b
}

// OutputUnigrams sets whether to emit tokens in both bigram and unigram form.
// If false, a CJK character is output in unigram form when it has no adjacent
// characters.
// Defaults to false.
func (b *TokenFilterCJKBigram) OutputUnigrams(outputUnigrams bool) *TokenFilterCJKBigram {
	b.outputUnigrams = &outputUnigrams
	return b
}

// Validate validates TokenFilterCJKBigram.
func (b *TokenFilterCJKBigram) Validate(includeName bool) error {
	var invalid []string
	if includeName && b.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(b.ignoredScripts) > 0 {
		for _, ignoredScript := range b.ignoredScripts {
			if _, valid := map[string]bool{
				"han":      true,
				"hangul":   true,
				"hiragana": true,
				"katakana": true,
			}[ignoredScript]; !valid {
				invalid = append(invalid, "IgnoredScripts")
				break
			}
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (b *TokenFilterCJKBigram) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "cjk_bigram",
	// 		"ignored_scripts": ["han","hangul","hiragana","katakana"],
	// 		"output_unigrams": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "cjk_bigram"

	if len(b.ignoredScripts) > 0 {
		options["ignored_scripts"] = b.ignoredScripts
	}
	if b.outputUnigrams != nil {
		options["output_unigrams"] = b.outputUnigrams
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[b.name] = options
	return source, nil
}
