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
	name string

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
	// 	"name": {
	// 		"type": "join",
	// 		"relations": {
	// 			"question": ["answer", "comment"],
	// 			"answer": "vote"
	// 		}
	// 		"doc_values": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "join"

	if len(j.relations) > 0 {
		relations := make(map[string]interface{})
		for _, r := range j.relations {
			relation, err := r.Source(false)
			if err != nil {
				return nil, err
			}
			relations[r.Name()] = relation
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
