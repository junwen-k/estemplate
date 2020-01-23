// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// DatatypeIP Specialised Datatype for IPv4 and IPv6 addresses.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ip.html
// for details.
type DatatypeIP struct {
	Datatype
	name   string
	copyTo []string

	// fields specific to ip datatype
	boost     *float32
	docValues *bool
	index     *bool
	nullValue string
	store     *bool
}

// NewDatatypeIP initializes a new DatatypeIP.
func NewDatatypeIP(name string) *DatatypeIP {
	return &DatatypeIP{
		name: name,
	}
}

// Name returns field key for the Datatype.
func (ip *DatatypeIP) Name() string {
	return ip.name
}

// CopyTo sets the field(s) to copy to which allows the values of multiple fields to be
// queried as a single field.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/copy-to.html
// for details.
func (ip *DatatypeIP) CopyTo(copyTo ...string) *DatatypeIP {
	ip.copyTo = append(ip.copyTo, copyTo...)
	return ip
}

// Boost sets Mapping field-level query time boosting. Defaults to 1.0.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-boost.html
// for details.
func (ip *DatatypeIP) Boost(boost float32) *DatatypeIP {
	ip.boost = &boost
	return ip
}

// DocValues sets whether if the field should be stored on disk in a column-stride fashion
// so that it can later be used for sorting, aggregations, or scripting.
// Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/doc-values.html
// for details.
func (ip *DatatypeIP) DocValues(docValues bool) *DatatypeIP {
	ip.docValues = &docValues
	return ip
}

// Index sets whether if the field should be searchable. Defaults to true.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-index.html
// for details.
func (ip *DatatypeIP) Index(index bool) *DatatypeIP {
	ip.index = &index
	return ip
}

// NullValue sets an IPv4 value which is substituted for any explicit null values.
// Defaults to null.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/null-value.html
// for details.
func (ip *DatatypeIP) NullValue(nullValue string) *DatatypeIP {
	ip.nullValue = nullValue
	return ip
}

// Store sets whether if the field value should be stored and retrievable separately
// from the `_source` field. Defaults to false.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-store.html
// for details.
func (ip *DatatypeIP) Store(store bool) *DatatypeIP {
	ip.store = &store
	return ip
}

// Validate validates DatatypeIP.
func (ip *DatatypeIP) Validate(includeName bool) error {
	var invalid []string
	if includeName && ip.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (ip *DatatypeIP) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "ip",
	// 		"copy_to": ["field_1", "field_2"],
	// 		"boost": 2,
	// 		"doc_values": true,
	// 		"index": true,
	// 		"null_value": "192.168.0.0/16",
	// 		"store" true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "ip"

	if len(ip.copyTo) > 0 {
		var copyTo interface{}
		switch {
		case len(ip.copyTo) > 1:
			copyTo = ip.copyTo
			break
		case len(ip.copyTo) == 1:
			copyTo = ip.copyTo[0]
			break
		default:
			copyTo = ""
		}
		options["copy_to"] = copyTo
	}
	if ip.boost != nil {
		options["boost"] = ip.boost
	}
	if ip.docValues != nil {
		options["doc_values"] = ip.docValues
	}
	if ip.index != nil {
		options["index"] = ip.index
	}
	if ip.nullValue != "" {
		options["null_value"] = ip.nullValue
	}
	if ip.store != nil {
		options["store"] = ip.store
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[ip.name] = options
	return source, nil
}
