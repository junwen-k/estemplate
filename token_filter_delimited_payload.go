// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterDelimitedPayload token filter that separates a token stream into tokens
// and payloads based on a specified delimiter.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-delimited-payload-tokenfilter.html
// for details.
type TokenFilterDelimitedPayload struct {
	TokenFilter
	name string

	// fields specific to delimited payload token filter
	delimiter string
	encoding  string
}

// NewTokenFilterDelimitedPayload initializes a new TokenFilterDelimitedPayload.
func NewTokenFilterDelimitedPayload(name string) *TokenFilterDelimitedPayload {
	return &TokenFilterDelimitedPayload{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (p *TokenFilterDelimitedPayload) Name() string {
	return p.name
}

// Delimiter sets the character used to separate tokens from payloads.
// Defaults to "|".
func (p *TokenFilterDelimitedPayload) Delimiter(delimiter string) *TokenFilterDelimitedPayload {
	p.delimiter = delimiter
	return p
}

// Encoding sets the Datatype for the stored payload.
// Can be set to the following values:
// "float" - Float
// "identity" - Characters
// "int" - Integer
// Defaults to "float".
func (p *TokenFilterDelimitedPayload) Encoding(encoding string) *TokenFilterDelimitedPayload {
	p.encoding = encoding
	return p
}

// Validate validates TokenFilterDelimitedPayload.
func (p *TokenFilterDelimitedPayload) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if p.encoding != "" {
		if _, valid := map[string]bool{
			"float":    true,
			"identity": true,
			"int":      true,
		}[p.encoding]; !valid {
			invalid = append(invalid, "Encoding")
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *TokenFilterDelimitedPayload) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "delimited_payload",
	// 		"delimiter": "+",
	// 		"encoding": "int"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "delimited_payload"

	if p.delimiter != "" {
		options["delimiter"] = p.delimiter
	}
	if p.encoding != "" {
		options["encoding"] = p.encoding
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
