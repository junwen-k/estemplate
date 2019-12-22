// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// MetaFieldSize Document size Meta-Field which indexes the size in bytes
// of the original `_source` field.
//
// See https://www.elastic.co/guide/en/elasticsearch/plugins/7.5/mapper-size.html
// for details.
type MetaFieldSize struct {
	MetaField

	// fields specific to size meta field
	enabled *bool
}

// NewMetaFieldSize initializes a NewMetaFieldSize.
func NewMetaFieldSize() *MetaFieldSize {
	return &MetaFieldSize{}
}

// Enabled sets whether to enable or disable the `_size` meta-field.
func (s *MetaFieldSize) Enabled(enabled bool) *MetaFieldSize {
	s.enabled = &enabled
	return s
}

// Validate validates MetaFieldSize.
func (s *MetaFieldSize) Validate() error {
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *MetaFieldSize) Source() (interface{}, error) {
	// "_size": {
	// 	"enabled": true
	// }
	options := make(map[string]interface{})

	if s.enabled != nil {
		options["enabled"] = s.enabled
	}

	source := make(map[string]interface{})
	source["_size"] = options
	return source, nil
}
