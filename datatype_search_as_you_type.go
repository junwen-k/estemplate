// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeSearchAsYouType Specialised Datatype for text-like field that is optimized
// to provide out-of-the-box support for queries that serve an as-you-type completion
// use case.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-as-you-type.html
// for details.
type DatatypeSearchAsYouType struct {
	Datatype
	name string

	// fields specific search as you type datatype
	maxShingleSize      *int
	analyzer            string
	index               *bool
	indexOptions        string
	norms               *bool
	store               *bool
	searchAnalyzer      string
	searchQuoteAnalyzer string
	similarity          string
	termVector          string
}

// NewDatatypeSearchAsYouType initializes a new DatatypeSearchAsYouType.
func NewDatatypeSearchAsYouType(name string) *DatatypeSearchAsYouType {
	return &DatatypeSearchAsYouType{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (t *DatatypeSearchAsYouType) Name() string {
	return t.name
}

// MaxShingleSize sets the largest shingle size to index the input with and create subfields for,
// creating one subfield for each shingle size between 2 and `max_shingle_size`.
// Can be set between 2 and 4 inclusive.
// Defaults to 3.
func (t *DatatypeSearchAsYouType) MaxShingleSize(maxShingleSize int) *DatatypeSearchAsYouType {
	t.maxShingleSize = &maxShingleSize
	return t
}

// Analyzer sets which analyzer should be used for analyzed string fields.
// Defaults to default index analyzer, or the "standard" analyzer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analyzer.html
// for details.
func (t *DatatypeSearchAsYouType) Analyzer(analyzer string) *DatatypeSearchAsYouType {
	t.analyzer = analyzer
	return t
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (t *DatatypeSearchAsYouType) Index(index bool) *DatatypeSearchAsYouType {
	t.index = &index
	return t
}

// IndexOptions sets information which will be stored in the index for search and highlighting purposes.
// Can be set to the following values:
// "docs" - Index only Doc number.
// "freqs" - Index both Doc number and term frequencies.
// "positions" - Index Doc number, term frequencies, and term positions (or order).
// "offsets" - Index Doc number, term frequencies, positions and start and end character offsets.
// Defaults to "positions".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-options.html
// for details.
func (t *DatatypeSearchAsYouType) IndexOptions(indexOptions string) *DatatypeSearchAsYouType {
	t.indexOptions = indexOptions
	return t
}

// Norms sets whether if field-length should be taken into account when scoring queries.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/norms.html
// for details.
func (t *DatatypeSearchAsYouType) Norms(norms bool) *DatatypeSearchAsYouType {
	t.norms = &norms
	return t
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (t *DatatypeSearchAsYouType) Store(store bool) *DatatypeSearchAsYouType {
	t.store = &store
	return t
}

// SearchAnalyzer sets the analyzer that should be used at search time on analyzed fields
// Defaults to `analyzer` setting.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-analyzer.html
// for details.
func (t *DatatypeSearchAsYouType) SearchAnalyzer(searchAnalyzer string) *DatatypeSearchAsYouType {
	t.searchAnalyzer = searchAnalyzer
	return t
}

// SearchQuoteAnalyzer sets the analyzer that should be used at search time when a phrase
// is encountered. Defaults to `search_analyzer` setting.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analyzer.html#search-quote-analyzer
// for details.
func (t *DatatypeSearchAsYouType) SearchQuoteAnalyzer(searchQuoteAnalyzer string) *DatatypeSearchAsYouType {
	t.searchQuoteAnalyzer = searchQuoteAnalyzer
	return t
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
func (t *DatatypeSearchAsYouType) Similarity(similarity string) *DatatypeSearchAsYouType {
	t.similarity = similarity
	return t
}

// TermVector sets whether term vectors should be stored for an analyzed field.
// Can be set to the following values:
// "no" - No term vectors are stored.
// "yes" - Just the terms in the field are stored.
// "with_positions" - Terms and positions are stored.
// "with_offsets" - Terms and character offsets are stored.
// "with_positions_offsets" - Terms, positions, and character offsets are stored.
// "with_positions_payloads" - Terms, positions, and payloads are stored.
// "with_positions_offsets_payloads" - Terms, positions, offsets and payloads are stored.
// Defaults to "no".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/term-vector.html`
// for details.
func (t *DatatypeSearchAsYouType) TermVector(termVector string) *DatatypeSearchAsYouType {
	t.termVector = termVector
	return t
}

// Validate validates DatatypeSearchAsYouType.
func (t *DatatypeSearchAsYouType) Validate(includeName bool) error {
	var invalid []string
	if includeName && t.name == "" {
		invalid = append(invalid, "Name")
	}
	if t.maxShingleSize != nil && *t.maxShingleSize < 2 || *t.maxShingleSize > 4 {
		invalid = append(invalid, "MaxShingleSize")
	}
	if t.indexOptions != "" {
		for _, valid := range validIndexOptions {
			if t.indexOptions != valid {
				invalid = append(invalid, "IndexOptions")
				break
			}
		}
	}
	if t.similarity != "" {
		for _, valid := range validSimilarity {
			if t.similarity != valid {
				invalid = append(invalid, "Similarity")
				break
			}
		}
	}
	if t.termVector != "" {
		for _, valid := range validTermVector {
			if t.termVector != valid {
				invalid = append(invalid, "TermVector")
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
func (t *DatatypeSearchAsYouType) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "search_as_you_type",
	// 		"max_shingle_size": 3
	// 		"analyzer": "my_analyzer",
	// 		"index": true,
	// 		"index_options": "positions",
	// 		"norms": true,
	// 		"store": true,
	// 		"search_analyzer": "my_stop_analyzer",
	// 		"search_quote_analyzer": "my_analyzer",
	// 		"similarity": "BM25",
	// 		"term_vector": "no"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "search_as_you_type"

	if t.maxShingleSize != nil {
		options["max_shingle_size"] = t.maxShingleSize
	}
	if t.analyzer != "" {
		options["analyzer"] = t.analyzer
	}
	if t.index != nil {
		options["index"] = t.index
	}
	if t.indexOptions != "" {
		options["index_options"] = t.indexOptions
	}
	if t.norms != nil {
		options["norms"] = t.norms
	}
	if t.store != nil {
		options["store"] = t.store
	}
	if t.searchAnalyzer != "" {
		options["search_analyzer"] = t.searchAnalyzer
	}
	if t.searchQuoteAnalyzer != "" {
		options["search_quote_analyzer"] = t.searchQuoteAnalyzer
	}
	if t.similarity != "" {
		options["similarity"] = t.similarity
	}
	if t.termVector != "" {
		options["term_vector"] = t.termVector
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[t.name] = options
	return source, nil
}
