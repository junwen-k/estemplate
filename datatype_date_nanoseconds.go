// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"fmt"
	"strings"
)

// DatatypeDateNanoseconds Core Datatype for date.
// - Date Datatype stores dates in millisecond resolution.
// - DateNanoseconds Datatype stores dates in nanosecond resolution.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/date_nanos.html
// for details.
type DatatypeDateNanoseconds struct {
	Datatype
	name string

	// fields specific to date nanoseconds datatype
	boost           *int
	docValues       *bool
	format          []*DateFormat
	rawFormat       string
	locale          string
	ignoreMalformed *bool
	index           *bool
	nullValue       interface{}
	store           *bool
}

// NewDatatypeDateNanoseconds initializes a new DatatypeDateNanoseconds.
func NewDatatypeDateNanoseconds(name string) *DatatypeDateNanoseconds {
	return &DatatypeDateNanoseconds{
		name:   name,
		format: make([]*DateFormat, 0),
	}
}

// Name returns field key for the Datatype.
func (d *DatatypeDateNanoseconds) Name() string {
	return d.name
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (d *DatatypeDateNanoseconds) Boost(boost int) *DatatypeDateNanoseconds {
	d.boost = &boost
	return d
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (d *DatatypeDateNanoseconds) DocValues(docValues bool) *DatatypeDateNanoseconds {
	d.docValues = &docValues
	return d
}

// RawFormat sets string type raw format and overwrites the current format value.
func (d *DatatypeDateNanoseconds) RawFormat(rawFormat string) *DatatypeDateNanoseconds {
	d.rawFormat = rawFormat
	return d
}

// Format sets date format for Elasticsearch to recognize and parse date string values.
// The first format will be used to convert the milliseconds-since-the-epoch value back into a string.
// Defaults to "strict_date_optional_time||epoch_millis"
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-date-format.html
// for details.
func (d *DatatypeDateNanoseconds) Format(format ...*DateFormat) *DatatypeDateNanoseconds {
	d.format = append(d.format, format...)
	return d
}

// Locale sets the locale to use when parsing dates since months do not have the same names
// and/or abbreviations in all languages.
// Defaults to "ROOT".
//
// See https://docs.oracle.com/javase/8/docs/api/java/util/Locale.html#ROOT
// for details.
func (d *DatatypeDateNanoseconds) Locale(locale string) *DatatypeDateNanoseconds {
	d.locale = locale
	return d
}

// IgnoreMalformed sets whether if the field should ignore malformed numbers.
// Defatuls to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-malformed.html
// for details.
func (d *DatatypeDateNanoseconds) IgnoreMalformed(ignoreMalformed bool) *DatatypeDateNanoseconds {
	d.ignoreMalformed = &ignoreMalformed
	return d
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (d *DatatypeDateNanoseconds) Index(index bool) *DatatypeDateNanoseconds {
	d.index = &index
	return d
}

// NullValue sets a date value in one of the configured format(s) as the field
// which is substituted for any explicit null values. Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (d *DatatypeDateNanoseconds) NullValue(nullValue interface{}) *DatatypeDateNanoseconds {
	d.nullValue = &nullValue
	return d
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (d *DatatypeDateNanoseconds) Store(store bool) *DatatypeDateNanoseconds {
	d.store = &store
	return d
}

// Validate validates DatatypeDateNanoseconds.
func (d *DatatypeDateNanoseconds) Validate(includeName bool) error {
	var invalid []string
	if includeName && d.name == "" {
		invalid = append(invalid, "Name")
	}
	if d.locale != "" {
		for _, valid := range validLocale {
			if d.locale != valid {
				invalid = append(invalid, "Locale")
				break
			}
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (d *DatatypeDateNanoseconds) Source(includeName bool) (interface{}, error) {
	// 	"test": {
	// 		"type": "date_nanos",
	// 		"boost": 2,
	// 		"doc_values": true,
	// 		"format": "strict_date_optional_time||epoch_millis",
	// 		"locale": "ROOT",
	// 		"ignore_malformed": true,
	// 		"index": true,
	// 		"null_value": 0,
	// 		"store": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "date_nanos"

	if d.boost != nil {
		options["boost"] = d.boost
	}
	if d.docValues != nil {
		options["doc_values"] = d.docValues
	}
	if len(d.format) > 0 {
		formats := make([]string, 0)
		for _, f := range d.format {
			f, err := f.Source()
			if err != nil {
				return nil, err
			}
			formats = append(formats, fmt.Sprintf("%s", f))
		}
		options["format"] = strings.Join(formats, "||")
	}
	if d.rawFormat != "" {
		options["format"] = d.rawFormat
	}
	if d.locale != "" {
		options["locale"] = d.locale
	}
	if d.ignoreMalformed != nil {
		options["ignore_malformed"] = d.ignoreMalformed
	}
	if d.index != nil {
		options["index"] = d.index
	}
	if d.nullValue != nil {
		options["null_value"] = d.nullValue
	}
	if d.store != nil {
		options["store"] = d.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[d.name] = options
	return source, nil
}