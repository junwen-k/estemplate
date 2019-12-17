// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeMapperMurmur3 (Plugin) Specialised Datatype to compute hashes of values at index-time
// and store them in the index. Typically used within a multi-field, so that both the original value
// and its hash are stored in the index.
//
// See https://www.elastic.co/guide/en/elasticsearch/plugins/7.5/mapper-murmur3.html
// for details.
type DatatypeMapperMurmur3 struct {
	Datatype
	name string

	// fields specific to mapper murmur3 datatype
}

// NewDatatypeMapperMurmur3 initializes a new DatatypeMapperMurmur3.
func NewDatatypeMapperMurmur3(name string) *DatatypeMapperMurmur3 {
	return &DatatypeMapperMurmur3{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (m3 *DatatypeMapperMurmur3) Name() string {
	return m3.name
}

// Validate validates DatatypeMapperMurmur3.
func (m3 *DatatypeMapperMurmur3) Validate(includeName bool) error {
	var invalid []string
	if includeName && m3.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (m3 *DatatypeMapperMurmur3) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "murmur3"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "murmur3"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[m3.name] = options
	return source, nil
}
