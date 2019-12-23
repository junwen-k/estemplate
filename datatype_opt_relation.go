// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// Relation Datatype parameter that defines a set of possible relations within
// the documents, each relation being a parent name and a child name.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/parent-join.html#parent-join
// for details.
type Relation struct {
	parent   string
	children []string
}

// NewRelation initializes a new Relation.
func NewRelation(parent string, childrens ...string) *Relation {
	return &Relation{
		parent:   parent,
		children: childrens,
	}
}

// Name returns field key (parent) for Datatype option.
func (r *Relation) Name() string {
	return r.parent
}

// Parent sets parent for the join relations.
func (r *Relation) Parent(parent string) *Relation {
	r.parent = parent
	return r
}

// Children sets children(s) for the join relations.
func (r *Relation) Children(childrens ...string) *Relation {
	r.children = append(r.children, childrens...)
	return r
}

// Source returns the serializable JSON for the source builder.
func (r *Relation) Source(includeName bool) (interface{}, error) {
	// {
	// 	"relations": {
	// 		"parent_1": ["children_1", "children_2"]
	// 	}
	// }
	options := make(map[string]interface{})

	if len(r.children) > 0 {
		var value interface{}
		switch {
		case len(r.children) > 1:
			value = r.children
			break
		case len(r.children) == 1:
			value = r.children[0]
			break
		default:
			value = ""
		}
		options[r.parent] = value
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["relations"] = options
	return source, nil
}
