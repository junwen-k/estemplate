// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterElision token filter that removes specified elisions from the beginning
// of tokens. For example, you can use this filter to change "l'avion" to "avion".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-elision-tokenfilter.html
// for details.
type TokenFilterElision struct {
	TokenFilter
	name string

	// fields specific to elision token filter
	articles     []string
	articlesPath string
	articlesCase *bool
}

// NewTokenFilterElision initializes a new TokenFilterElision.
func NewTokenFilterElision(name string) *TokenFilterElision {
	return &TokenFilterElision{
		name:     name,
		articles: make([]string, 0),
	}
}

// Name returns field key for the Token Filter.
func (e *TokenFilterElision) Name() string {
	return e.name
}

// Articles sets a list of elisions to remove. To be removed, the elision must
// be at the beginning of a token and be immediately followed by an apostrophe ("'").
// Both the elision and apostrophe are removed.
// Defaults to [l', m', t', qu', n', s', j', d', c', jusqu', quoiqu', lorsqu', puisqu'] (French).
// Either this or `articles_path` parameter must be specified.
func (e *TokenFilterElision) Articles(articles ...string) *TokenFilterElision {
	e.articles = append(e.articles, articles...)
	return e
}

// ArticlesPath sets a path to a file containing a list of elisions to remove.
// This path must be absolute or relative to the `config` location.
// The file must be UTF-8 encoded. Each token in the file must be separated by a line break.
// To be removed, the elision must be at the beginning of a token and be immediately
// followed by an apostrophe ("'").  Both the elision and apostrophe are removed.
// Either this or `articles` parameter must be specified.
func (e *TokenFilterElision) ArticlesPath(articlesPath string) *TokenFilterElision {
	e.articlesPath = articlesPath
	return e
}

// ArticlesCase sets whether provided elisions should be case sensitive or not.
// Defaults to false.
func (e *TokenFilterElision) ArticlesCase(articlesCase bool) *TokenFilterElision {
	e.articlesCase = &articlesCase
	return e
}

// Validate validates TokenFilterElision.
func (e *TokenFilterElision) Validate(includeName bool) error {
	var invalid []string
	if includeName && e.name == "" {
		invalid = append(invalid, "Name")
	}
	if !(len(e.articles) > 0) && e.articlesPath == "" {
		invalid = append(invalid, "Articles || ArticlesPath")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (e *TokenFilterElision) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "elision",
	// 		"articles": ["l", "m", "t", "qu", "n", "s", "j"],
	// 		"articles_path": "token_filter/example_elision_list.txt",
	// 		"articles_case": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "elision"

	if len(e.articles) > 0 {
		options["articles"] = e.articles
	}
	if e.articlesPath != "" {
		options["articles_path"] = e.articlesPath
	}
	if e.articlesCase != nil {
		options["articles_case"] = e.articlesCase
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[e.name] = options
	return source, nil
}
