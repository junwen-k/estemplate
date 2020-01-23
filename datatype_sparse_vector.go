// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeSparseVector Specialised Datatype that stores dense vectors of float
// values. The maximum number of dimensions that can be in a vector should not exceed
// 1024. The number of dimensions can be different across documents.
// ! Experimental and may be changed or removed completely in a future release.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/sparse-vector.html
// for details.
type DatatypeSparseVector struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to sparse vector datatype
}

// NewDatatypeSparseVector initializes a new DatatypeSparseVector.
func NewDatatypeSparseVector(name string) *DatatypeSparseVector {
	return &DatatypeSparseVector{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (v *DatatypeSparseVector) Name() string {
	return v.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (v *DatatypeSparseVector) CopyTo(copyTo ...string) *DatatypeSparseVector {
	v.copyTo = append(v.copyTo, copyTo...)
	return v
}

// Validate validates DatatypeSparseVector.
func (v *DatatypeSparseVector) Validate(includeName bool) error {
	var invalid []string
	if includeName && v.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (v *DatatypeSparseVector) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "sparse_vector",
	// 		"copy_to": ["field_1", "field_2"]
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "sparse_vector"

	if len(v.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(v.copyTo) > 1:
			copyTo = v.copyTo
			break
		case len(v.copyTo) == 1:
			copyTo = v.copyTo[0]
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
	source[v.name] = options
	return source, nil
}
