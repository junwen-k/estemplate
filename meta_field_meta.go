// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "encoding/json"

// MetaFieldMeta Other Meta-Field which sets application specific metadata.
// A mapping type can have custom meta data associated with it. THese are not
// used at all by Elasticsearch, but can be used to store application-specific
// metadata, such as the class that a document belongs to.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-meta-field.html
// for details.
type MetaFieldMeta struct {
	value   interface{}
	rawJSON string
}

// NewMetaFieldMeta initializes a new MetaFieldMeta.
func NewMetaFieldMeta() *MetaFieldMeta {
	return &MetaFieldMeta{}
}

// Value sets a value (interface{}) for the meta data, in which will later
// be marshalled into JSON string.
func (m *MetaFieldMeta) Value(value interface{}) *MetaFieldMeta {
	m.value = value
	return m
}

// RawJSON sets the Raw JSON string for the meta data.
func (m *MetaFieldMeta) RawJSON(rawJSON string) *MetaFieldMeta {
	m.rawJSON = rawJSON
	return m
}

// Validate validates MetaFieldMeta.
func (m *MetaFieldMeta) Validate() error {
	return nil
}

// Source returns the serializable JSON for the source builder.
func (m *MetaFieldMeta) Source(includeName bool) (interface{}, error) {
	// {
	// 	"_meta": {
	// 		"class": "MyApp::User",
	// 		"version": {
	// 			"min": "1.0",
	// 			"max": "1.3"
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})

	var (
		b   []byte
		err error
	)
	if m.value != nil {
		b, err = json.Marshal(m.value)
		if err != nil {
			return nil, err
		}
	}
	if m.rawJSON != "" {
		b = []byte(m.rawJSON)
	}

	if b != nil {
		err = json.Unmarshal(b, &options)
		if err != nil {
			return nil, err
		}
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["_meta"] = options
	return source, nil
}
