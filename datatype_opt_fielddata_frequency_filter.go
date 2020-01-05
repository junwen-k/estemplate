// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// FielddataFrequencyFilter Datatype parameter to reduce the number of terms loaded into memory.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/fielddata.html#field-data-filtering
// for details.
type FielddataFrequencyFilter struct {
	min            float32
	max            float32
	minSegmentSize int
}

// NewFielddataFrequencyFilter initializes a new FielddataFrequencyFilter.
func NewFielddataFrequencyFilter(min, max float32) *FielddataFrequencyFilter {
	return &FielddataFrequencyFilter{
		min: min,
		max: max,
	}
}

// Min sets minimum document frequency for frequency filter.
func (f *FielddataFrequencyFilter) Min(min float32) *FielddataFrequencyFilter {
	f.min = min
	return f
}

// Max sets maximum document frequency for frequency filter.
func (f *FielddataFrequencyFilter) Max(max float32) *FielddataFrequencyFilter {
	f.max = max
	return f
}

// MinSegmentSize sets minimum number of docs for frequency filter to exclude small segments.
func (f *FielddataFrequencyFilter) MinSegmentSize(minSegmentSize int) *FielddataFrequencyFilter {
	f.minSegmentSize = minSegmentSize
	return f
}

// Source returns the serializable JSON for the source builder.
func (f *FielddataFrequencyFilter) Source(includeName bool) (interface{}, error) {
	// {
	// 	"fielddata_frequency_filter": {
	// 		"min": 0.001,
	// 		"max": 0.1,
	// 		"min_segment_size": 500
	// 	}
	// }
	options := make(map[string]interface{})

	if f.min > 0 {
		options["min"] = f.min
	}
	if f.max > 0 {
		options["max"] = f.max
	}
	if f.minSegmentSize > 0 {
		options["min_segment_size"] = f.minSegmentSize
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["fielddata_frequency_filter"] = options
	return source, nil
}
