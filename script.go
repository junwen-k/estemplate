// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// Script that evaluate custom expressions in Elasticsearch. For example, you
// could use a script to return "script fields" as part of a search request or
// evaluate a custom score for a query.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/modules-scripting.html
// for details.
type Script struct {
	lang   string
	source string
	id     string
	params map[string]interface{}
}

// NewScript initializes a new Script.
func NewScript(source string) *Script {
	return &Script{
		source: source,
		params: make(map[string]interface{}),
	}
}

// Lang sets the language the script is written in.
// Can be set to the following values:
// "painless" - general-purpose language
// "expression" - fast custom ranking and sorting
// "mustache" - templates
// "java" - expert API
// Defaults to "painless".
//
// Languages that are sandboxed are designed with security in mind. Non-sandboxed languages can be a security issue.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/modules-scripting-security.html
// for details.
func (s *Script) Lang(lang string) *Script {
	s.lang = lang
	return s
}

// ScriptSource sets the source of the inline script.
func (s *Script) ScriptSource(script string) *Script {
	s.source = script
	return s
}

// ID sets the ID of the stored script to be retrieved from the cluster state.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/modules-scripting-using.html#modules-scripting-stored-scripts
// for details.
func (s *Script) ID(id string) *Script {
	s.id = id
	return s
}

// RawParams sets the named parameters that are passed into the script as variables.
func (s *Script) RawParams(rawParams map[string]interface{}) *Script {
	s.params = rawParams
	return s
}

// Params sets a key-value pair into the named parameters that are passed into the script as variables.
func (s *Script) Params(key string, value interface{}) *Script {
	s.params[key] = value
	return s
}

// Validate validates Script.
func (s *Script) Validate() error {
	var invalid []string
	if s.source == "" && s.id == "" {
		invalid = append(invalid, "Source || ID")
	}
	if s.source != "" {
		if _, valid := map[string]bool{
			"painless":   true,
			"expression": true,
			"mustache":   true,
			"java":       true,
		}[s.source]; !valid {
			invalid = append(invalid, "Source")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *Script) Source(includeName bool) (interface{}, error) {
	// {
	// 	"lang": "painless",
	// 	"source": "doc['my_field'] * multiplier",
	// 	"id": "calculate-score",
	// 	"params": {
	// 		"multiplier": 2
	// 	}
	// }
	options := make(map[string]interface{})

	if s.lang != "" {
		options["lang"] = s.lang
	}
	if s.source != "" {
		options["source"] = s.source
	}
	if s.id != "" {
		options["id"] = s.id
	}
	if len(s.params) > 0 {
		options["params"] = s.params
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["script"] = options
	return source, nil
}
