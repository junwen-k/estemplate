// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeFlattened Specialised Datatype that allows an entire JSON object
// to be indexed as a single field. This data type can be useful for indexing
// objects with a large or unknown number of unique keys.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/flattened.html
// for details.
type DatatypeFlattened struct {
	Datatype
	name string

	// fields specific to flattened datatype
	boost                    *float32
	depthLimit               *int
	docValues                *bool
	eagerGlobalOrdinals      *bool
	ignoreAbove              *int
	index                    *bool
	indexOptions             string
	nullValue                string
	similarity               string
	splitQueriesOnWhitespace *bool
}

// NewDatatypeFlattened initializes a new DatatypeFlattened.
func NewDatatypeFlattened(name string) *DatatypeFlattened {
	return &DatatypeFlattened{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (f *DatatypeFlattened) Name() string {
	return f.name
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (f *DatatypeFlattened) Boost(boost float32) *DatatypeFlattened {
	f.boost = &boost
	return f
}

// DepthLimit sets the maximum allowed depth of the flattened object field,
// in terms of nested inner objects. Note that `depth_limit` can be updated
// dynamically through the put mapping API.
// Defaults to 20.
func (f *DatatypeFlattened) DepthLimit(depthLimit int) *DatatypeFlattened {
	f.depthLimit = &depthLimit
	return f
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (f *DatatypeFlattened) DocValues(docValues bool) *DatatypeFlattened {
	f.docValues = &docValues
	return f
}

// EagerGlobalOrdinals sets whether if global ordinals be loaded eagerly on refresh. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/eager-global-ordinals.html
// for details.
func (f *DatatypeFlattened) EagerGlobalOrdinals(eagerGlobalOrdinals bool) *DatatypeFlattened {
	f.eagerGlobalOrdinals = &eagerGlobalOrdinals
	return f
}

// IgnoreAbove sets the limit for the leaf values to be indexed, leaf values longer than
// the `ignore_above` setting will not be indexed or stored. Note that this limit applies
// to the leaf values within the flattened object field, and not the length of the entire
// field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-above.html
// for details.
func (f *DatatypeFlattened) IgnoreAbove(ignoreAbove int) *DatatypeFlattened {
	f.ignoreAbove = &ignoreAbove
	return f
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (f *DatatypeFlattened) Index(index bool) *DatatypeFlattened {
	f.index = &index
	return f
}

// IndexOptions sets information which will be stored in the index for search and highlighting purposes.
// Can be set to the following values:
// "docs" - Index only Doc number.
// "freqs" - Index both Doc number and term frequencies.
// "positions" - Index Doc number, term frequencies, and term positions (or order).
// "offsets" - Index Doc number, term frequencies, positions and start and end character offsets.
// Defaults to "docs".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-options.html
// for details.
func (f *DatatypeFlattened) IndexOptions(indexOptions string) *DatatypeFlattened {
	f.indexOptions = indexOptions
	return f
}

// NullValue sets a string value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (f *DatatypeFlattened) NullValue(nullValue string) *DatatypeFlattened {
	f.nullValue = nullValue
	return f
}

// Similarity sets the scoring algorithm or similarity that should be used.
// Can be set to the following values:
// "BM25" - Okapi BM25 algorithm.
// "classic" - TF/IDF algorithm.
// "boolean" - Simple boolean similarity.
// Defaults to "BM25".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/similarity.html
// for details.
func (f *DatatypeFlattened) Similarity(similarity string) *DatatypeFlattened {
	f.similarity = similarity
	return f
}

// SplitQueriesOnWhitespace sets whether full text queries should split the input on
// whitespace when building a query for this field.
// Defaults to false.
func (f *DatatypeFlattened) SplitQueriesOnWhitespace(splitQueriesOnWhitespace bool) *DatatypeFlattened {
	f.splitQueriesOnWhitespace = &splitQueriesOnWhitespace
	return f
}

// Validate validates DatatypeFlattened.
func (f *DatatypeFlattened) Validate(includeName bool) error {
	var invalid []string
	if includeName && f.name == "" {
		invalid = append(invalid, "Name")
	}
	if f.indexOptions != "" {
		for _, valid := range validIndexOptions {
			if f.indexOptions != valid {
				invalid = append(invalid, "IndexOptions")
				break
			}
		}
	}
	if f.similarity != "" {
		for _, valid := range validSimilarity {
			if f.similarity != valid {
				invalid = append(invalid, "Similarity")
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
func (f *DatatypeFlattened) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "flattened",
	// 		"boost": 2,
	// 		"depth_limit": 20,
	// 		"doc_values": true,
	// 		"eager_global_ordinals": true,
	// 		"ignore_above": 256,
	// 		"index": true,
	// 		"index_options": "docs",
	// 		"null_value": "NULL",
	// 		"similarity": "BM25",
	// 		"split_queries_on_whitespace": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "flattened"

	if f.boost != nil {
		options["boost"] = f.boost
	}
	if f.depthLimit != nil {
		options["depth_limit"] = f.depthLimit
	}
	if f.docValues != nil {
		options["doc_values"] = f.docValues
	}
	if f.eagerGlobalOrdinals != nil {
		options["eager_global_ordinals"] = f.eagerGlobalOrdinals
	}
	if f.ignoreAbove != nil {
		options["ignore_above"] = f.ignoreAbove
	}
	if f.index != nil {
		options["index"] = f.index
	}
	if f.indexOptions != "" {
		options["index_options"] = f.IndexOptions
	}
	if f.nullValue != "" {
		options["null_value"] = f.nullValue
	}
	if f.similarity != "" {
		options["similarity"] = f.similarity
	}
	if f.splitQueriesOnWhitespace != nil {
		options["split_queries_on_whitespace"] = f.splitQueriesOnWhitespace
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[f.name] = options
	return source, nil
}
