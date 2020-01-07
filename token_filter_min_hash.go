// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterMinHash token filter that hashes each token of the token stream
// and divdes the resulting hashes into buckets, keeping the lowest-valued hashes
// per bucket. It then returns these hashes as tokens.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-minhash-tokenfilter.html
// for details.
type TokenFilterMinHash struct {
	TokenFilter
	name string

	// fields specific to min hash token filter
	hashCount    *int
	bucketCount  *int
	hashSetSize  *int
	withRotation *bool
}

// NewTokenFilterMinHash initializes a new TokenFilterMinHash.
func NewTokenFilterMinHash(name string) *TokenFilterMinHash {
	return &TokenFilterMinHash{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (h *TokenFilterMinHash) Name() string {
	return h.name
}

// HashCount sets the number of hases to has the token stream with.
// Defaults to 1.
func (h *TokenFilterMinHash) HashCount(hashCount int) *TokenFilterMinHash {
	h.hashCount = &hashCount
	return h
}

// BucketCount sets the number of buckets to divde the minhashes into.
// Defaults to 512.
func (h *TokenFilterMinHash) BucketCount(bucketCount int) *TokenFilterMinHash {
	h.bucketCount = &bucketCount
	return h
}

// HashSetSize sets the number of minhashes to keep per bucket.
// Defaults to 1.
func (h *TokenFilterMinHash) HashSetSize(hashSetSize int) *TokenFilterMinHash {
	h.hashSetSize = &hashSetSize
	return h
}

// WithRotation sets whether or not to kill empty buckets with the value of the
// first non-empty bucket to its circular right. Only takes effect if `hash_set_size`
// is equal to one.
// Defaults to true if `bucket_count` is greater than one, else false.
func (h *TokenFilterMinHash) WithRotation(withRotation bool) *TokenFilterMinHash {
	h.withRotation = &withRotation
	return h
}

// Validate validates TokenFilterMinHash.
func (h *TokenFilterMinHash) Validate(includeName bool) error {
	var invalid []string
	if includeName && h.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields or invalid values: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (h *TokenFilterMinHash) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "min_hash",
	// 		"hash_count": 1,
	// 		"bucket_count": 512,
	// 		"hash_set_size": 1,
	// 		"with_rotation": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "min_hash"

	if h.hashCount != nil {
		options["hash_count"] = h.hashCount
	}
	if h.bucketCount != nil {
		options["bucket_count"] = h.bucketCount
	}
	if h.hashSetSize != nil {
		options["hash_set_size"] = h.hashSetSize
	}
	if h.withRotation != nil {
		options["with_rotation"] = h.withRotation
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[h.name] = options
	return source, nil
}
