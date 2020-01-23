// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeText Core Datatype for string to index full-text values, such as the
// body of an email or the description of a product.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/text.html
// for details.
type DatatypeText struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to text datatype
	analyzer                 string
	boost                    *float32
	eagerGlobalOrdinals      *bool
	fielddata                *bool
	fielddataFrequencyFilter *FielddataFrequencyFilter
	fields                   []Datatype
	index                    *bool
	indexOptions             string
	indexPrefixes            *IndexPrefixes
	indexPhrases             *bool
	norms                    *bool
	positionIncrementGap     *int
	store                    *bool
	searchAnalyzer           string
	searchQuoteAnalyzer      string
	similarity               string
	termVector               string
}

// NewDatatypeText initializes a new DatatypeText.
func NewDatatypeText(name string) *DatatypeText {
	return &DatatypeText{
		name:   name,
		fields: make([]Datatype, 0),
	}
}

// Name returns field key for the Datatype.
func (t *DatatypeText) Name() string {
	return t.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (t *DatatypeText) CopyTo(copyTo ...string) *DatatypeText {
	t.copyTo = append(t.copyTo, copyTo...)
	return t
}

// Analyzer sets which analyzer should be used for analyzed string fields.
// Defaults to default index analyzer, or the "standard" analyzer.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analyzer.html
// for details.
func (t *DatatypeText) Analyzer(analyzer string) *DatatypeText {
	t.analyzer = analyzer
	return t
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (t *DatatypeText) Boost(boost float32) *DatatypeText {
	t.boost = &boost
	return t
}

// EagerGlobalOrdinals sets whether if global ordinals be loaded eagerly on refresh. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/eager-global-ordinals.html
// for details.
func (t *DatatypeText) EagerGlobalOrdinals(eagerGlobalOrdinals bool) *DatatypeText {
	t.eagerGlobalOrdinals = &eagerGlobalOrdinals
	return t
}

// Fielddata sets whether if the field use in-memory fielddata for sorting, aggregations,
// or scripting. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/fielddata.html
// for details.
func (t *DatatypeText) Fielddata(fielddata bool) *DatatypeText {
	t.fielddata = &fielddata
	return t
}

// FielddataFrequencyFilter sets frequency filter to reduce the number of terms loaded into memory,
// and thus reduce memory usage.
// By default all values are loaded.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/fielddata.html#field-data-filtering
// for details.
func (t *DatatypeText) FielddataFrequencyFilter(fielddataFrequencyFilter *FielddataFrequencyFilter) *DatatypeText {
	t.fielddataFrequencyFilter = fielddataFrequencyFilter
	return t
}

// Fields sets multi-fields which allow the same string value to be indexed in multiple
// ways for different purposes.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/multi-fields.html
// for details.
func (t *DatatypeText) Fields(fields ...Datatype) *DatatypeText {
	t.fields = append(t.fields, fields...)
	return t
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (t *DatatypeText) Index(index bool) *DatatypeText {
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
func (t *DatatypeText) IndexOptions(indexOptions string) *DatatypeText {
	t.indexOptions = indexOptions
	return t
}

// IndexPrefixes sets index prefixes to enable indexing of term prefixes to speed up prefix searches.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-prefixes.html
// for details.
func (t *DatatypeText) IndexPrefixes(indexPrefixes *IndexPrefixes) *DatatypeText {
	t.indexPrefixes = indexPrefixes
	return t
}

// IndexPhrases sets whether should two-term word combinations (shingles) indexed into a separate field.
// Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-phrases.html
// for details.
func (t *DatatypeText) IndexPhrases(indexPhrases bool) *DatatypeText {
	t.indexPhrases = &indexPhrases
	return t
}

// Norms sets whether if field-length should be taken into account when scoring queries.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/norms.html
// for details.
func (t *DatatypeText) Norms(norms bool) *DatatypeText {
	t.norms = &norms
	return t
}

// PositionIncrementGap sets the number of fake term position which should be inserted between
// each element of an array of strings.
// Defaults to the settings of `position_increment_gap` which defaults to 100.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/position-increment-gap.html
// for details.
func (t *DatatypeText) PositionIncrementGap(positionIncrementGap int) *DatatypeText {
	t.positionIncrementGap = &positionIncrementGap
	return t
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (t *DatatypeText) Store(store bool) *DatatypeText {
	t.store = &store
	return t
}

// SearchAnalyzer sets the analyzer that should be used at search time on analyzed fields
// Defaults to `analyzer` setting.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-analyzer.html
// for details.
func (t *DatatypeText) SearchAnalyzer(searchAnalyzer string) *DatatypeText {
	t.searchAnalyzer = searchAnalyzer
	return t
}

// SearchQuoteAnalyzer sets the analyzer that should be used at search time when a phrase
// is encountered. Defaults to `search_analyzer` setting.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analyzer.html#search-quote-analyzer
// for details.
func (t *DatatypeText) SearchQuoteAnalyzer(searchQuoteAnalyzer string) *DatatypeText {
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
func (t *DatatypeText) Similarity(similarity string) *DatatypeText {
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
func (t *DatatypeText) TermVector(termVector string) *DatatypeText {
	t.termVector = termVector
	return t
}

// Validate validates DatatypeText.
func (t *DatatypeText) Validate(includeName bool) error {
	var invalid []string
	if includeName && t.name == "" {
		invalid = append(invalid, "Name")
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
	if t.fielddataFrequencyFilter != nil && t.fielddata != nil && *t.fielddata {
		invalid = append(invalid, "FielddataFrequencyFilter")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (t *DatatypeText) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "text",
	// 		"copy_to": ["field_1", "field_2"],
	// 		"analyzer": "my_analyzer",
	// 		"boost": 2,
	// 		"eager_global_ordinals": true,
	// 		"fielddata": true,
	// 		"fielddata_frequency_filter": {
	// 			"min": 0.001,
	// 			"max": 0.1,
	// 			"min_segment_size": 500
	// 		},
	// 		"fields": {
	// 			"field_name": {
	// 				"type": "text",
	// 				"analyzer": "standard"
	// 			}
	// 		},
	// 		"index": true,
	// 		"index_options": "positions",
	// 		"index_prefixes": {
	// 			"min_chars": 2,
	// 			"max_chars": 20,
	// 		},
	// 		"index_phrases": true,
	// 		"norms": true,
	// 		"position_increment_gap": 1,
	// 		"store": true,
	// 		"search_analyzer": "my_stop_analyzer",
	// 		"search_quote_analyzer": "my_analyzer",
	// 		"similarity": "BM25",
	// 		"term_vector": "no"
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "text"

	if len(t.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(t.copyTo) > 1:
			copyTo = t.copyTo
			break
		case len(t.copyTo) == 1:
			copyTo = t.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if t.analyzer != "" {
		options["analyzer"] = t.analyzer
	}
	if t.boost != nil {
		options["boost"] = t.boost
	}
	if t.eagerGlobalOrdinals != nil {
		options["eager_global_ordinals"] = t.eagerGlobalOrdinals
	}
	if t.fielddata != nil {
		options["fielddata"] = t.fielddata
	}
	if t.fielddataFrequencyFilter != nil {
		fielddataFrequencyFilter, err := t.fielddataFrequencyFilter.Source(false)
		if err != nil {
			return nil, err
		}
		options["fielddata_frequency_filter"] = fielddataFrequencyFilter
	}
	if len(t.fields) > 0 {
		fields := make(map[string]interface{})
		for _, f := range t.fields {
			field, err := f.Source(false)
			if err != nil {
				return nil, err
			}
			fields[f.Name()] = field
		}
		options["fields"] = fields
	}
	if t.index != nil {
		options["index"] = t.index
	}
	if t.indexOptions != "" {
		options["index_options"] = t.indexOptions
	}
	if t.indexPrefixes != nil {
		indexPrefixes, err := t.indexPrefixes.Source(false)
		if err != nil {
			return nil, err
		}
		options["index_prefixes"] = indexPrefixes
	}
	if t.indexPhrases != nil {
		options["index_phrases"] = t.indexPhrases
	}
	if t.norms != nil {
		options["norms"] = t.norms
	}
	if t.positionIncrementGap != nil {
		options["position_increment_gap"] = t.positionIncrementGap
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
