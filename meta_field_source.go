// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// MetaFieldSource Document source Meta-Field that contains the original
// JSON representing the body of the document. The `_source` field itself is
// not indexed (and thus is not searchable), but it is stored so that it can be
// returned when executing fetch requests, like get or search.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-source-field.html
// for details.
type MetaFieldSource struct {
	MetaField

	// fields specific to source meta field
	enabled  *bool
	includes []string
	excludes []string
}

// NewMetaFieldSource initializes a NewMetaFieldSource.
func NewMetaFieldSource() *MetaFieldSource {
	return &MetaFieldSource{}
}

// Enabled sets whether to disable the `_source` field to prevent incur storage
// overhead.
func (s *MetaFieldSource) Enabled(enabled bool) *MetaFieldSource {
	s.enabled = &enabled
	return s
}

// Includes sets the contents to be included after the document has been indexed,
// but before the `_source` field is stored. Accepts wildcards.
func (s *MetaFieldSource) Includes(includes ...string) *MetaFieldSource {
	s.includes = append(s.includes, includes...)
	return s
}

// Excludes sets the contents to be excluded after the document has been indexed,
// but before the `_source` field is stored. Accepts wildcards.
func (s *MetaFieldSource) Excludes(excludes ...string) *MetaFieldSource {
	s.excludes = append(s.excludes, excludes...)
	return s
}

// Validate validates MetaFieldSource.
func (s *MetaFieldSource) Validate() error {
	return nil
}

// Source returns the serializable JSON for the source builder.
func (s *MetaFieldSource) Source() (interface{}, error) {
	// "_source": {
	// 	"enabled": true,
	// 	"includes": [
	// 		"*.count",
	// 		"meta.*"
	// 	],
	// 	"excludes": [
	// 		"meta.description",
	// 		"meta.other.*"
	// 	]
	// }
	options := make(map[string]interface{})

	if s.enabled != nil {
		options["enabled"] = s.enabled
	}
	if len(s.includes) > 0 {
		options["includes"] = s.includes
	}
	if len(s.excludes) > 0 {
		options["excludes"] = s.excludes
	}

	source := make(map[string]interface{})
	source["_source"] = options
	return source, nil
}
