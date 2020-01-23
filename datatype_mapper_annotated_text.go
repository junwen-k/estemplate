// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeMapperAnnotatedText (Plugin) Specialised Datatype that tokenizes text content
// as per the more common `text` field but also injects any marked-up annotation tokens directly
// into the search index.
//
// See https://www.elastic.co/guide/en/elasticsearch/plugins/7.5/mapper-annotated-text.html
// for details.
type DatatypeMapperAnnotatedText struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to mapper annotated text datatype
}

// NewDatatypeMapperAnnotatedText initializes a new DatatypeMapperAnnotatedText.
func NewDatatypeMapperAnnotatedText(name string) *DatatypeMapperAnnotatedText {
	return &DatatypeMapperAnnotatedText{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (t *DatatypeMapperAnnotatedText) Name() string {
	return t.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (t *DatatypeMapperAnnotatedText) CopyTo(copyTo ...string) *DatatypeMapperAnnotatedText {
	t.copyTo = append(t.copyTo, copyTo...)
	return t
}

// Validate validates DatatypeMapperAnnotatedText.
func (t *DatatypeMapperAnnotatedText) Validate(includeName bool) error {
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
func (t *DatatypeMapperAnnotatedText) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "annotated_text",
	// 		"copy_to": ["field_1", "field_2"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "annotated_text"

	if len(t.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(t.copyTo) > 1:
			copyTo = t.copyTo
			break
		case len(t.copyTo) == 1:
			copyTo = t.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[t.name] = options
	return source, nil
}
