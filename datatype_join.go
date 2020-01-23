// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeJoin Specialised Datatype that creates parent/child relation
// within documents of the same index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/parent-join.html
// for details.
type DatatypeJoin struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to join datatype
	relations           []*Relation
	eagerGlobalOrdinals *bool
}

// NewDatatypeJoin initializes a new DatatypeJoin.
func NewDatatypeJoin(name string) *DatatypeJoin {
	return &DatatypeJoin{
		name:      name,
		relations: make([]*Relation, 0),
	}
}

// Name returns field key for the Datatype.
func (j *DatatypeJoin) Name() string {
	return j.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (j *DatatypeJoin) CopyTo(copyTo ...string) *DatatypeJoin {
	j.copyTo = append(j.copyTo, copyTo...)
	return j
}

// Relations sets parent/child relation within the documents.
func (j *DatatypeJoin) Relations(relations ...*Relation) *DatatypeJoin {
	j.relations = append(j.relations, relations...)
	return j
}

// EagerGlobalOrdinals sets global ordinals to speed up joins. When the join field is
// used infrequently and writes occur frequently it may make sense to disable eager loading.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/parent-join.html#_global_ordinals
// for details.
func (j *DatatypeJoin) EagerGlobalOrdinals(eagerGlobalOrdinals bool) *DatatypeJoin {
	j.eagerGlobalOrdinals = &eagerGlobalOrdinals
	return j
}

// Validate validates DatatypeJoin.
func (j *DatatypeJoin) Validate(includeName bool) error {
	var invalid []string
	if includeName && j.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (j *DatatypeJoin) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "join",
	// 		"copy_to": ["field_1", "field_2"],
	// 		"relations": {
	// 			"question": ["answer", "comment"],
	// 			"answer": "vote"
	// 		}
	// 		"doc_values": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "join"

	if len(j.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(j.copyTo) > 1:
			copyTo = j.copyTo
			break
		case len(j.copyTo) == 1:
			copyTo = j.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if len(j.relations) > 0 {
		relations := make(map[string]interface{})
		for _, r := range j.relations {
			relation, err := r.Source(false)
			if err != nil {
				return nil, err
			}
			for parent, value := range relation.(map[string]interface{}) {
				relations[parent] = value
			}
		}
		options["relations"] = relations
	}
	if j.eagerGlobalOrdinals != nil {
		options["eager_global_ordinals"] = j.eagerGlobalOrdinals
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[j.name] = options
	return source, nil
}
