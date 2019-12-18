// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypePercolator Specialised Datatype that parses a JSON structure
// into a native query and stores that query, so that the percolate query
// can use it to match provided documents.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/percolator.html
// for details.
type DatatypePercolator struct {
	Datatype
	name string

	// fields specific to percolator datatype
}

// NewDatatypePercolator initializes a new DatatypePercolator.
func NewDatatypePercolator(name string) *DatatypePercolator {
	return &DatatypePercolator{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (p *DatatypePercolator) Name() string {
	return p.name
}

// Validate validates DatatypePercolator.
func (p *DatatypePercolator) Validate(includeName bool) error {
	var invalid []string
	if includeName && p.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *DatatypePercolator) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "percolator"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "percolator"

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[p.name] = options
	return source, nil
}
