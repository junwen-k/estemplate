// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeObject Complex Datatype for JSON object. The document may
// contain inner objects which, in turn, may contain inner objects
// themselves.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/object.html
// for details
type DatatypeObject struct {
	Datatype
	name string

	// fields specific to object datatype
	dynamic    *bool
	strict     *bool
	enabled    *bool
	properties []Datatype
}

// NewDatatypeObject initializes a new DatatypeObject.
func NewDatatypeObject(name string) *DatatypeObject {
	return &DatatypeObject{
		name:       name,
		properties: make([]Datatype, 0),
	}
}

// Name returns field key for the Datatype.
func (o *DatatypeObject) Name() string {
	return o.name
}

// Dynamic sets whether if fields can be added dynamically to a document.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic.html
// for details.
func (o *DatatypeObject) Dynamic(dynamic bool) *DatatypeObject {
	o.dynamic = &dynamic
	return o
}

// Strict sets dynamic setting to use "strict" which throw an exception if new
// fields are detected.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic.html
// for details.
func (o *DatatypeObject) Strict(strict bool) *DatatypeObject {
	o.strict = &strict
	return o
}

// Enabled sets whether if the JSON value given for the object field should be parsed
// and indexed, or completely ignored. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/enabled.html
// for details.
func (o *DatatypeObject) Enabled(enabled bool) *DatatypeObject {
	o.enabled = &enabled
	return o
}

// Properties sets the fields within the object, which can be of any datatype, including
// object. New properties may be added to an existing object.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/properties.html
// for details.
func (o *DatatypeObject) Properties(properties ...Datatype) *DatatypeObject {
	o.properties = append(o.properties, properties...)
	return o
}

// Validate validates DatatypeObject.
func (o *DatatypeObject) Validate(includeName bool) error {
	var invalid []string
	if includeName && o.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (o *DatatypeObject) Source(includeName bool) (interface{}, error) {
	// {
	// 	"name": {
	// 		"type": "object",
	// 		"dynamic": true,
	// 		"enabled": true,
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

	if o.dynamic != nil {
		options["dynamic"] = o.dynamic
	}
	if o.strict != nil && *o.strict {
		options["dynamic"] = "strict"
	}
	if o.enabled != nil {
		options["enabled"] = o.enabled
	}
	if len(o.properties) > 0 {
		properties := make(map[string]interface{})
		for _, f := range o.properties {
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
	source[o.name] = options
	return source, nil
}
