// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterStemmer token filter that provides access to (almost) all available
// stemming token filters through a single unified interface.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-stemmer-tokenfilter.html
// for details.
type TokenFilterStemmer struct {
	TokenFilter
	name string

	// fields specific to stemmer token filter
	language string
	// name field is omitted
}

// NewTokenFilterStemmer initializes a new TokenFilterStemmer.
func NewTokenFilterStemmer(name string) *TokenFilterStemmer {
	return &TokenFilterStemmer{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (s *TokenFilterStemmer) Name() string {
	return s.name
}

// Language sets the stemmer to be used by the filter.
// Can be set to the following values:
// "arabic" - Arabic
// "armenian" - Armenian
// "basque" - Basque
// "bengali" || "light_bengali" - Bengali
// "brazilian" - Brazlian Portuguese
// "bulgarian" - Bulgarian
// "catalan" - Catalan
// "czech" - Czech
// "danish" - Danish
// "dutch" || "dutch_kp" - Dutch
// "english" || "light_english" || "minimal_english" || "possessive_english" || "porter2" || "lovins" - English
// "finnish" || "light_finnish" - Finnish
// "french" || "light_french" || "minimal_french" - French
// "galician" || "minimal_galician" - Galician (Plural step only)
// "german" || "german2" || "light_german" || "minimal_german" - German
// "greek" - Greek
// "hindi" - Hindi
// "hungarian" || "light_hungarian" - Hungarian
// "indonesian" - Indonesian
// "irish" - Irish
// "italian" || "light_italian" - Italian
// "sorani" - Kurdish (Sorani)
// "latvian" - Latvian
// "lithuanian" - Lithuanian
// "norwegian" || "light_norwegian" || "minimal_norwegian" - Norwegian (Bokmål)
// "light_nyrorsk" || "minimal_nynorsk" - Norwegian (Nynorsk)
// "portuguese" || "light_portuguese" || "minimal_portuguese" || "portuguese_rslp" - Portuguese
// "romanian" - Romanian
// "russian" || "light_russian" - Russian
// "spanish" || "light_spanish" - Spanish
// "swedish" || "light_swedish" - Swedish
// "turkish" - Turkish
func (s *TokenFilterStemmer) Language(language string) *TokenFilterStemmer {
	s.language = language
	return s
}

// Validate validates TokenFilterStemmer.
func (s *TokenFilterStemmer) Validate(includeName bool) error {
	var invalid []string
	if includeName && s.name == "" {
		invalid = append(invalid, "Name")
	}
	if s.language != "" {
		if _, valid := map[string]bool{
			"arabic":             true,
			"armenian":           true,
			"basque":             true,
			"bengali":            true,
			"light_bengali":      true,
			"brazilian":          true,
			"bulgarian":          true,
			"catalan":            true,
			"czech":              true,
			"danish":             true,
			"dutch":              true,
			"dutch_kp":           true,
			"english":            true,
			"light_english":      true,
			"minimal_english":    true,
			"possessive_english": true,
			"porter2":            true,
			"lovins":             true,
			"finnish":            true,
			"light_finnish":      true,
			"french":             true,
			"light_french":       true,
			"minimal_french":     true,
			"galician":           true,
			"minimal_galician":   true,
			"german":             true,
			"german2":            true,
			"light_german":       true,
			"minimal_german":     true,
			"greek":              true,
			"hindi":              true,
			"hungarian":          true,
			"light_hungarian":    true,
			"indonesian":         true,
			"irish":              true,
			"italian":            true,
			"light_italian":      true,
			"sorani":             true,
			"latvian":            true,
			"lithuanian":         true,
			"norwegian":          true,
			"light_norwegian":    true,
			"minimal_norwegian":  true,
			"light_nyrorsk":      true,
			"minimal_nyrorsk":    true,
			"portuguese":         true,
			"light_portuguese":   true,
			"minimal_portuguese": true,
			"portuguese_rslp":    true,
			"romanian":           true,
			"russian":            true,
			"light_russian":      true,
			"spanish":            true,
			"light_spanish":      true,
			"swedish":            true,
			"light_swedish":      true,
			"turkish":            true,
		}[s.language]; !valid {
			invalid = append(invalid, "Language")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *TokenFilterStemmer) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "stemmer",
	// 		"language": "arabic"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "stemmer"

	if s.language != "" {
		options["language"] = s.language
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[s.name] = options
	return source, nil
}
