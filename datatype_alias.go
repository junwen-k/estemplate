// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeAlias Specialised Datatype defines an alternate name for a field
// in the index. The alias can be used in place of the target field in search
// requests, and selected other APIs like field capabilities.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/alias.html
// for details.
type DatatypeAlias struct {
	Datatype
	name string

	// fields specific to alias datatype
	path string
}

// NewDatatypeAlias initializes a new DatatypeAlias.
func NewDatatypeAlias(name string) *DatatypeAlias {
	return &DatatypeAlias{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (a *DatatypeAlias) Name() string {
	return a.name
}

// Path sets the path to the target field. Note that this must be the full path,
// including any parent objects (e.g. object1.object2.field).
func (a *DatatypeAlias) Path(path string) *DatatypeAlias {
	a.path = path
	return a
}

// Validate validates DatatypeAlias.
func (a *DatatypeAlias) Validate(includeName bool) error {
	var invalid []string
	if includeName && a.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (a *DatatypeAlias) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "alias",
	// 		"path": "distance"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "alias"

	if a.path != "" {
		options["path"] = a.path
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[a.name] = options
	return source, nil
}
