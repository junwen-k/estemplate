// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeDenseVector Specialised Datatype that stores dense vectors of float
// values. The maximum number of dimensions that can be in a vector should not exceed
// 1024.
// ! Experimental and may be changed or removed completely in a future release.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dense-vector.html
// for details.
type DatatypeDenseVector struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to dense vector datatype
	dims *int
}

// NewDatatypeDenseVector initializes a new DatatypeDenseVector.
func NewDatatypeDenseVector(name string) *DatatypeDenseVector {
	return &DatatypeDenseVector{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (v *DatatypeDenseVector) Name() string {
	return v.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (v *DatatypeDenseVector) CopyTo(copyTo ...string) *DatatypeDenseVector {
	v.copyTo = append(v.copyTo, copyTo...)
	return v
}

// Dims sets the number of dimensions in the vector. Internally, each document's dense
// vector is encoded as a binary doc value. Its size in bytes is equal to 4 * dims + 4,
// where dims - the number of the vector's dimensions.
func (v *DatatypeDenseVector) Dims(dims int) *DatatypeDenseVector {
	v.dims = &dims
	return v
}

// Validate validates DatatypeDenseVector.
func (v *DatatypeDenseVector) Validate(includeName bool) error {
	var invalid []string
	if includeName && v.name == "" {
		invalid = append(invalid, "Name")
	}
	// TODO: validate dims
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (v *DatatypeDenseVector) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "dense_vector",
	// 		"copy_to": ["field_1", "field_2"],
	// 		"dims": 3
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "dense_vector"

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
	if v.dims != nil {
		options["dims"] = v.dims
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[v.name] = options
	return source, nil
}
