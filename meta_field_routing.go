// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// MetaFieldRouting Routing Meta-Field which routes a document to a particular
// shard.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-routing-field.html
// for details.
type MetaFieldRouting struct {
	required *bool
}

// NewMetaFieldRouting initializes a new MetaFieldRouting.
func NewMetaFieldRouting() *MetaFieldRouting {
	return &MetaFieldRouting{}
}

// Required sets whether if the `_routing` field is required for all CRUD operations to prevent
// routing value being forgotten and lead to a document being indexed on more than on shard.
func (r *MetaFieldRouting) Required(required bool) *MetaFieldRouting {
	r.required = &required
	return r
}

// Validate validates MetaFieldRouting.
func (r *MetaFieldRouting) Validate() error {
	return nil
}

// Source returns the serializable JSON for the source builder.
func (r *MetaFieldRouting) Source(includeName bool) (interface{}, error) {
	// {
	// 	"_routing": {
	// 		"required": true
	// 	}
	// }
	options := make(map[string]interface{})

	if r.required != nil {
		options["required"] = r.required
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["_routing"] = options
	return source, nil
}
