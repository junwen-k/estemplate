// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// IndexPrefixes Datatype parameter that enables the indexing of term prefixes to speed
// up prefix searches.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-prefixes.html
// for details.
type IndexPrefixes struct {
	DatatypeOption

	// fields specific to index prefixes datatype option
	minChars int
	maxChars int
}

// NewIndexPrefixes initializes a new IndexPrefixes.
func NewIndexPrefixes(minChars, maxChars int) *IndexPrefixes {
	return &IndexPrefixes{
		minChars: minChars,
		maxChars: maxChars,
	}
}

// MinChars sets the minimum prefix length to index, must be greater than 0.
// Defaults to 2.
func (p *IndexPrefixes) MinChars(minChars int) *IndexPrefixes {
	p.minChars = minChars
	return p
}

// MaxChars sets the maximum prefix length to index, must be less than 20.
// Defaults to 5.
func (p *IndexPrefixes) MaxChars(maxChars int) *IndexPrefixes {
	p.maxChars = maxChars
	return p
}

// Validate validates IndexPrefixes.
func (p *IndexPrefixes) Validate() error {
	var invalid []string
	if p.minChars <= 0 {
		invalid = append(invalid, "MinChars")
	}
	if p.maxChars > 20 {
		invalid = append(invalid, "MaxChars")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (p *IndexPrefixes) Source() (interface{}, error) {
	// {
	// 	"min_chars": 2,
	// 	"max_chars": 20,
	// }
	source := make(map[string]interface{})

	if p.minChars > 0 {
		source["min_chars"] = p.minChars
	}
	if p.maxChars > 0 {
		source["max_chars"] = p.maxChars
	}

	return source, nil
}
