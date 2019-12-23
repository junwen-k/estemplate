// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// Mappings defines how a document, and the fields it contains, are stored and indexed.
// For instance, use mappings to define:
// - which string fields should be treated as full text fields.
// - which fields contain numbers, dates, or geolocations.
// - the format of date values.
// - custom rules to control the mapping for dynamically added fields.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping.html
// for details.
type Mappings struct {
	// dynamic templates
	dynamicTemplates []*DynamicTemplate

	// dynamic mapping fields
	dateDetection      *bool
	dynamicDateFormats []*DateFormat
	numericDetection   *bool

	// meta fields
	source     *MetaFieldSource
	size       *MetaFieldSize
	fieldNames *MetaFieldFieldNames
	routing    *MetaFieldRouting
	meta       *MetaFieldMeta

	// properties fields
	properties []Datatype
}

// NewMappings initializes a new Mappings.
func NewMappings() *Mappings {
	return &Mappings{
		dynamicTemplates:   make([]*DynamicTemplate, 0),
		dynamicDateFormats: make([]*DateFormat, 0),
		properties:         make([]Datatype, 0),
	}
}

// DynamicTemplates sets custom rules to configure the mapping for dynamically added fields.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-templates.html
// for details.
func (m *Mappings) DynamicTemplates(dynamicTemplates ...*DynamicTemplate) *Mappings {
	m.dynamicTemplates = append(m.dynamicTemplates, dynamicTemplates...)
	return m
}

// DateDetection sets whether new string fields are checked to see whether their contents match
// any of the date patterns specified in `dynamic_date_formats`. If a match is found, a new `date`
// field is added with the corresponding format. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-field-mapping.html#date-detection
// for details.
func (m *Mappings) DateDetection(dateDetection bool) *Mappings {
	m.dateDetection = &dateDetection
	return m
}

// DynamicDateFormats sets date formats for date detections.
// Defaults to ["strict_date_optional_time","yyyy/MM/dd HH:mm:ss Z||yyyy/MM/dd Z"].
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-field-mapping.html#_customising_detected_date_formats
// for details.
func (m *Mappings) DynamicDateFormats(dynamicDateFormats ...*DateFormat) *Mappings {
	m.dynamicDateFormats = append(m.dynamicDateFormats, dynamicDateFormats...)
	return m
}

// NumericDetection sets whether new numeric like fields should be added as numeric type.
// Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/dynamic-field-mapping.html#numeric-detection
// for details.
func (m *Mappings) NumericDetection(numericDetection bool) *Mappings {
	m.numericDetection = &numericDetection
	return m
}

// MetaSource sets the original JSON representing the body of the document.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-source-field.html
// for details.
func (m *Mappings) MetaSource(metaSource *MetaFieldSource) *Mappings {
	m.source = metaSource
	return m
}

// Size sets the size of the `_source` field in bytes, provided by the `mapper-size` plugin.
//
// See https://www.elastic.co/guide/en/elasticsearch/plugins/7.5/mapper-size.html
// for details.
func (m *Mappings) Size(size *MetaFieldSize) *Mappings {
	m.size = size
	return m
}

// FieldNames sets wether to index the names of every field in a document that contains any value
// other than null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-field-names-field.html
// for details.
func (m *Mappings) FieldNames(fieldNames *MetaFieldFieldNames) *Mappings {
	m.fieldNames = fieldNames
	return m
}

// Routing sets whether if custom routing value which routes a document to a particular shard
// is required.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-routing-field.html
// for details.
func (m *Mappings) Routing(routing *MetaFieldRouting) *Mappings {
	m.routing = routing
	return m
}

// Meta sets application specific metadata.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-meta-field.html
// for details.
func (m *Mappings) Meta(meta *MetaFieldMeta) *Mappings {
	m.meta = meta
	return m
}

// Properties sets the list of fields or properties pertinent to the document.
func (m *Mappings) Properties(properties ...Datatype) *Mappings {
	m.properties = append(m.properties, properties...)
	return m
}

// Validate validates Mappings.
func (m *Mappings) Validate() error {
	return nil
}

// Source returns the serializable JSON for the source builder.
func (m *Mappings) Source(includeName bool) (interface{}, error) {
	// {
	// 	"mappings": {
	// 		"dynamic_templates": [
	// 			{
	// 				"integers": {
	// 					"match_mapping_type": "long",
	// 					"mapping": {
	// 						"type": "integer"
	// 					}
	// 				}
	// 			}
	// 		],
	// 		"date_detection": false,
	// 		"dynamic_date_formats": ["strict_date_optional_time","yyyy/MM/dd HH:mm:ss Z||yyyy/MM/dd Z"],
	// 		"numeric_detection": false,
	// 		"_source": {
	// 			"enabled": true,
	// 			"includes": [
	// 				"*.count",
	// 				"meta.*"
	// 			],
	// 			"excludes": [
	// 				"meta.description",
	// 				"meta.other.*"
	// 			]
	// 		},
	// 		"_size": {
	// 			"enabled": true
	// 		},
	// 		"_field_names": {
	// 			"enabled": true
	// 		},
	// 		"_routing": {
	// 			"required": true
	// 		},
	// 		"_meta": {
	// 			"class": "MyApp::User",
	// 			"version": {
	// 				"min": "1.0",
	// 				"max": "1.3"
	// 			}
	// 		},
	// 		"properties": {
	// 			"field_name": {
	// 				"type": "text",
	// 				"analyzer": "standard"
	// 			}
	// 		}
	// 	}
	// }
	options := make(map[string]interface{})

	if len(m.dynamicTemplates) > 0 {
		dynamicTemplates := make([]interface{}, 0)
		for _, t := range m.dynamicTemplates {
			tpl := make(map[string]interface{})
			dynamicTemplate, err := t.Source(false)
			if err != nil {
				return nil, err
			}
			tpl[t.Name()] = dynamicTemplate
			dynamicTemplates = append(dynamicTemplates, tpl)
		}
		options["dynamic_templates"] = dynamicTemplates
	}
	if m.dateDetection != nil {
		options["date_detection"] = m.dateDetection
	}
	if len(m.dynamicDateFormats) > 0 {
		dynamicDateFormats := make([]string, 0)
		for _, f := range m.dynamicDateFormats {
			format, err := f.Source()
			if err != nil {
				return nil, err
			}
			dynamicDateFormats = append(dynamicDateFormats, fmt.Sprintf("%s", format))
		}
		// Omit join with "||"
		options["dynamic_date_formats"] = dynamicDateFormats
	}
	if m.numericDetection != nil {
		options["numeric_detection"] = m.numericDetection
	}
	if m.source != nil {
		source, err := m.source.Source(false)
		if err != nil {
			return nil, err
		}
		options["_source"] = source
	}
	if m.size != nil {
		size, err := m.size.Source(false)
		if err != nil {
			return nil, err
		}
		options["_size"] = size
	}
	if m.fieldNames != nil {
		fieldNames, err := m.fieldNames.Source(false)
		if err != nil {
			return nil, err
		}
		options["_field_names"] = fieldNames
	}
	if m.routing != nil {
		routing, err := m.routing.Source(false)
		if err != nil {
			return nil, err
		}
		options["_routing"] = routing
	}
	if m.meta != nil {
		meta, err := m.meta.Source(false)
		if err != nil {
			return nil, err
		}
		options["_meta"] = meta
	}
	if len(m.properties) > 0 {
		properties := make(map[string]interface{})
		for _, p := range m.properties {
			property, err := p.Source(false)
			if err != nil {
				return nil, err
			}
			properties[p.Name()] = property
		}
		options["properties"] = properties
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["mappings"] = options
	return source, nil
}
