// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DynamicTemplate defines custom mappings that can be applied to dynamically
// added fields based on:
// - the datatype detected by Elasticsearch, with `match_mapping_type`.
// - the name of the field, with `match` and `unmatch` or `match_pattern`.
// - the full dotted path to the field, with `path_match` and `path_unmatch`.
//
// The original field name {name} and the detected datatype {dynamic_type}
// template variables can be used in the mapping specification as placeholders.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html
// for details.
type DynamicTemplate struct {
	name string

	// fields specific to dynamic template
	matchMappingType string
	match            string
	unmatch          string
	matchPattern     string
	pathMatch        string
	pathUnmatch      string
	mapping          Datatype
}

// NewDynamicTemplate initializes a new DynamicTemplate.
func NewDynamicTemplate(name string) *DynamicTemplate {
	return &DynamicTemplate{
		name: name,
	}
}

// Name returns field key for the DynamicTemplate.
func (t *DynamicTemplate) Name() string {
	return t.name
}

// MatchMappingType sets the datatype detected by the json parser. Since JSON doesn't allow
// to distinguish a `long` from an `integer` or a `double` from a `float`, it will always choose
// the wider datatype, i.e. `long` for integers and `double` for floating-point numbers.
// "*" may also be used in order to match all datatypes.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html#match-mapping-type
// for details.
func (t *DynamicTemplate) MatchMappingType(matchMappingType string) *DynamicTemplate {
	t.matchMappingType = matchMappingType
	return t
}

// Match sets the pattern to match on the field name. Wildcard "*" is supported.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html#match-unmatch
// for details.
func (t *DynamicTemplate) Match(match string) *DynamicTemplate {
	t.match = match
	return t
}

// Unmatch sets the pattern to exclude fields matched by `match`. Wildcard "*" is supported.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html#match-unmatch
// for details.
func (t *DynamicTemplate) Unmatch(unmatch string) *DynamicTemplate {
	t.unmatch = unmatch
	return t
}

// MatchPattern sets the behavior of the match parameter such that it supports full
// Java regular expression ("regex") matching on the field name instead of simple wildcards.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html#match-pattern
// for details.
func (t *DynamicTemplate) MatchPattern(matchPattern string) *DynamicTemplate {
	t.matchPattern = matchPattern
	return t
}

// PathMatch sets the full dotted path pattern to match on the field, not just the final name.
// Wildcard "*" is supported.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html#path-match-unmatch
// for details.
func (t *DynamicTemplate) PathMatch(pathMatch string) *DynamicTemplate {
	t.pathMatch = pathMatch
	return t
}

// PathUnmatch sets the full dotted path pattern to exclude fields matched by `match`, not just the final name.
// Wildcard "*" is supported.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html#path-match-unmatch
// for details.
func (t *DynamicTemplate) PathUnmatch(pathUnmatch string) *DynamicTemplate {
	t.pathUnmatch = pathUnmatch
	return t
}

// Mapping sets the mapping that the matched field should use.
func (t *DynamicTemplate) Mapping(mapping Datatype) *DynamicTemplate {
	t.mapping = mapping
	return t
}

// Validate validates DynamicTemplate.
func (t *DynamicTemplate) Validate(includeName bool) error {
	var invalid []string
	if includeName && t.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (t *DynamicTemplate) Source(includeName bool) (interface{}, error) {
	// {
	// 	"strings": {
	// 		"match_mapping_type": "string",
	// 		"match_pattern": "regex",
	// 		"match": "^profit_\d+$",
	// 		"unmatch": "*_text",
	// 		"path_match":   "name.*",
	// 		"path_unmatch": "*.middle",
	// 		"mapping": {
	// 			"type": "{dynamic_type}",
	// 			"analyzer": "{name}"
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})

	if t.matchMappingType != "" {
		options["match_mapping_type"] = t.matchMappingType
	}
	if t.matchPattern != "" {
		options["match_pattern"] = t.matchPattern
	}
	if t.match != "" {
		options["match"] = t.match
	}
	if t.unmatch != "" {
		options["unmatch"] = t.unmatch
	}
	if t.pathMatch != "" {
		options["path_match"] = t.pathMatch
	}
	if t.pathUnmatch != "" {
		options["path_unmatch"] = t.pathUnmatch
	}
	if t.mapping != nil {
		mapping, err := t.mapping.Source(false)
		if err != nil {
			return nil, err
		}
		options["mapping"] = mapping
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[t.name] = options
	return source, nil
}
