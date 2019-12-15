// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeKeyword Core Datatype for string to index structured content
// such as email addresses, hostnames, status codes, zip codes or tags.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/keyword.html
// for details.
type DatatypeKeyword struct {
	Datatype
	name string

	// fields specific to keyword datatype
	boost                    *int
	docValues                *bool
	eagerGlobalOrdinals      *bool
	fields                   []Datatype
	ignoreAbove              *int
	index                    *bool
	indexOptions             string
	norms                    *bool
	nullValue                string
	store                    *bool
	similarity               string
	normalizer               string
	splitQueriesOnWhitespace *bool
}

// NewDatatypeKeyword initializes a new DatatypeKeyword.
func NewDatatypeKeyword(name string) *DatatypeKeyword {
	return &DatatypeKeyword{
		name:   name,
		fields: make([]Datatype, 0),
	}
}

// Name returns field key for the Datatype.
func (k *DatatypeKeyword) Name() string {
	return k.name
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (k *DatatypeKeyword) Boost(boost int) *DatatypeKeyword {
	k.boost = &boost
	return k
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (k *DatatypeKeyword) DocValues(docValues bool) *DatatypeKeyword {
	k.docValues = &docValues
	return k
}

// EagerGlobalOrdinals sets whether if global ordinals be loaded eagerly on refresh. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/eager-global-ordinals.html
// for details.
func (k *DatatypeKeyword) EagerGlobalOrdinals(eagerGlobalOrdinals bool) *DatatypeKeyword {
	k.eagerGlobalOrdinals = &eagerGlobalOrdinals
	return k
}

// Fields sets multi-fields which allow the same string value to be indexed in multiple
// ways for different purposes.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/multi-fields.html
// for details.
func (k *DatatypeKeyword) Fields(fields ...Datatype) *DatatypeKeyword {
	k.fields = append(k.fields, fields...)
	return k
}

// IgnoreAbove sets the limit for the string length to be indexed, strings longer than
// the `ignore_above` setting will not be indexed or stored.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ignore-above.html
// for details.
func (k *DatatypeKeyword) IgnoreAbove(ignoreAbove int) *DatatypeKeyword {
	k.ignoreAbove = &ignoreAbove
	return k
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (k *DatatypeKeyword) Index(index bool) *DatatypeKeyword {
	k.index = &index
	return k
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
func (k *DatatypeKeyword) IndexOptions(indexOptions string) *DatatypeKeyword {
	k.indexOptions = indexOptions
	return k
}

// Norms sets whether if field-length should be taken into account when scoring queries.
// Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/norms.html
// for details.
func (k *DatatypeKeyword) Norms(norms bool) *DatatypeKeyword {
	k.norms = &norms
	return k
}

// NullValue sets a string value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (k *DatatypeKeyword) NullValue(nullValue string) *DatatypeKeyword {
	k.nullValue = nullValue
	return k
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (k *DatatypeKeyword) Store(store bool) *DatatypeKeyword {
	k.store = &store
	return k
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
func (k *DatatypeKeyword) Similarity(similarity string) *DatatypeKeyword {
	k.similarity = similarity
	return k
}

// Normalizer sets the normalizer that should be used.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/normalizer.html
// for details.
func (k *DatatypeKeyword) Normalizer(normalizer string) *DatatypeKeyword {
	k.normalizer = normalizer
	return k
}

// SplitQueriesOnWhitespace sets whether full text queries should split the input on
// whitespace when building a query for this field.
// Defaults to false.
func (k *DatatypeKeyword) SplitQueriesOnWhitespace(splitQueriesOnWhitespace bool) *DatatypeKeyword {
	k.splitQueriesOnWhitespace = &splitQueriesOnWhitespace
	return k
}

// Validate validates DatatypeKeyword.
func (k *DatatypeKeyword) Validate(includeName bool) error {
	var invalid []string
	if includeName && k.name == "" {
		invalid = append(invalid, "Name")
	}
	if k.indexOptions != "" {
		for _, valid := range validIndexOptions {
			if k.indexOptions != valid {
				invalid = append(invalid, "IndexOptions")
				break
			}
		}
	}
	if k.similarity != "" {
		for _, valid := range validSimilarity {
			if k.similarity != valid {
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
func (k *DatatypeKeyword) Source(includeName bool) (interface{}, error) {
	// {
	// 	"name": {
	// 		"type": "keyword",
	// 		"boost": 2,
	// 		"doc_values": true,
	// 		"eager_global_ordinals": true,
	// 		"fields": {
	// 			"field_name": {
	// 				"type": "text",
	// 				"analzyer": "standard"
	// 			}
	// 		},
	// 		"ignore_above": 256,
	// 		"index": true,
	// 		"index_options": "docs",
	// 		"norms": true,
	// 		"null_value": "NULL",
	// 		"store": true,
	// 		"similarity": "BM25",
	// 		"normalizer": "my_normalizer",
	// 		"split_queries_on_whitespace": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "keyword"

	if k.boost != nil {
		options["boost"] = k.boost
	}
	if k.docValues != nil {
		options["doc_values"] = k.docValues
	}
	if k.eagerGlobalOrdinals != nil {
		options["eager_global_ordinals"] = k.eagerGlobalOrdinals
	}
	if len(k.fields) > 0 {
		fields := make(map[string]interface{})
		for _, f := range k.fields {
			field, err := f.Source(false)
			if err != nil {
				return nil, err
			}
			fields[f.Name()] = field
		}
		options["fields"] = fields
	}
	if k.ignoreAbove != nil {
		options["ignore_above"] = k.ignoreAbove
	}
	if k.index != nil {
		options["index"] = k.index
	}
	if k.indexOptions != "" {
		options["index_options"] = k.IndexOptions
	}
	if k.norms != nil {
		options["norms"] = k.norms
	}
	if k.nullValue != "" {
		options["null_value"] = k.nullValue
	}
	if k.store != nil {
		options["store"] = k.store
	}
	if k.similarity != "" {
		options["similarity"] = k.similarity
	}
	if k.normalizer != "" {
		options["normalizer"] = k.normalizer
	}
	if k.splitQueriesOnWhitespace != nil {
		options["split_queries_on_whitespace"] = k.splitQueriesOnWhitespace
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[k.name] = options
	return source, nil
}
