// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeNested Complex Datatype for JSON array of objects to be indexed
// in a way they can be queried independently of each other.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/nested.html
// for details.
type DatatypeNested struct {
	Datatype
	name string

	// fields specific to nested datatype
	dynamic    *bool
	strict     bool
	properties []Datatype
}

// NewDatatypeNested initializes a new DatatypeNested.
func NewDatatypeNested(name string) *DatatypeNested {
	return &DatatypeNested{
		name:       name,
		properties: make([]Datatype, 0),
	}
}

// Name returns field key for the Datatype.
func (n *DatatypeNested) Name() string {
	return n.name
}

// Dynamic sets whether if fields can be added dynamically to a document.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic.html
// for details.
func (n *DatatypeNested) Dynamic(dynamic bool) *DatatypeNested {
	n.dynamic = &dynamic
	return n
}

// Strict sets dynamic setting to use "strict" which throw an exception if new
// fields are detected.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic.html
// for details.
func (n *DatatypeNested) Strict(strict bool) *DatatypeNested {
	n.strict = strict
	return n
}

// Properties sets the fields within the object, which can be of any datatype, including
// object. New properties may be added to an existing object.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/properties.html
// for details.
func (n *DatatypeNested) Properties(properties ...Datatype) *DatatypeNested {
	n.properties = append(n.properties, properties...)
	return n
}

// Validate validates DatatypeNested.
func (n *DatatypeNested) Validate(includeName bool) error {
	var invalid []string
	if includeName && n.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (n *DatatypeNested) Source(includeName bool) (interface{}, error) {
	// {
	// 	"name": {
	// 		"type": "nested",
	// 		"dynamic": true,
	// 		"properties": {
	// 			"field_name": {
	// 				"type": "text",
	// 				"analzyer": "standard"
	// 			}
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "object"

	if n.dynamic != nil {
		options["dynamic"] = n.dynamic
	}
	if n.strict {
		options["dynamic"] = "strict"
	}
	if len(n.properties) > 0 {
		properties := make(map[string]interface{})
		for _, f := range n.properties {
			property, err := f.Source(false)
			if err != nil {
				return nil, err
			}
			properties[f.Name()] = property
		}
		options["properties"] = properties
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[n.name] = options
	return source, nil
}
