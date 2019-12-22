// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// MetaFieldFieldNames Indexing Meta-Field which index the names of every
// field in a document that contains any value other than `null`,
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-field-names-field.html
// for details.
type MetaFieldFieldNames struct {
	enabled *bool
}

// NewMetaFieldFieldNames initializes a new MetaFieldFieldNames.
func NewMetaFieldFieldNames() *MetaFieldFieldNames {
	return &MetaFieldFieldNames{}
}

// Enabled sets whether to enable or disable the `_field_names` meta-field.
// Disabling `field_names` is usually not necessary because it no longer carries
// the index overhead it once did.
// ! Disabling `_field_names` has been deprecated and will be removed in a future major version.
func (n *MetaFieldFieldNames) Enabled(enabled bool) *MetaFieldFieldNames {
	n.enabled = &enabled
	return n
}

// Validate validates MetaFieldFieldNames.
func (n *MetaFieldFieldNames) Validate() error {
	return nil
}

// Source returns the serializable JSON for the source builder.
func (n *MetaFieldFieldNames) Source() (interface{}, error) {
	// {
	// 	"_field_names": {
	// 		"enabled": true
	// 	}
	// }
	options := make(map[string]interface{})

	if n.enabled != nil {
		options["enabled"] = n.enabled
	}

	source := make(map[string]interface{})
	source["_field_names"] = options
	return source, nil
}
