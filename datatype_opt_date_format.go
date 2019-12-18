// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// DateFormat Datatype parameter for Date Datatype.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-date-format.html
// for details.
type DateFormat struct {
	strict bool
	format string
}

// NewDateFormat initializes a new DateFormat.
func NewDateFormat(format string) *DateFormat {
	return &DateFormat{
		format: format,
	}
}

// Format sets format for date format.
func (f *DateFormat) Format(format string) *DateFormat {
	f.format = format
	return f
}

// Strict sets strict parsing for date format. Useful when date fields are dynamically mapped
// in order to make sure to not accidentally map irrelevant strings as dates.
func (f *DateFormat) Strict(strict bool) *DateFormat {
	f.strict = strict
	return f
}

// Source returns the serializable JSON for the source builder.
func (f *DateFormat) Source() (interface{}, error) {
	// "strict_format"
	source := f.format

	if f.strict {
		source = "strict_" + f.format
	}

	return source, nil
}
