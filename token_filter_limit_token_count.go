// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import "fmt"

// TokenFilterLimitTokenCount token filter that limits the number of output tokens.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-limit-token-count-tokenfilter.html
// for details.
type TokenFilterLimitTokenCount struct {
	TokenFilter
	name string

	// fields specific to limit token count token filter
	maxTokenCount    *int
	consumeAllTokens *bool
}

// NewTokenFilterLimitTokenCount initializes a new TokenFilterLimitTokenCount.
func NewTokenFilterLimitTokenCount(name string) *TokenFilterLimitTokenCount {
	return &TokenFilterLimitTokenCount{
		name: name,
	}
}

// Name returns field key for the Token Filter.
func (c *TokenFilterLimitTokenCount) Name() string {
	return c.name
}

// MaxTokenCount sets the maximum number of tokens to keep. Once this limit is reached, any remaining
// tokens are excluded from the output.
// Defaults to 1.
func (c *TokenFilterLimitTokenCount) MaxTokenCount(maxTokenCount int) *TokenFilterLimitTokenCount {
	c.maxTokenCount = &maxTokenCount
	return c
}

// ConsumeAllTokens sets whether filter exhausts the token stream, even if the `max_token_count` has
// already been reached.
// Defaults to false.
func (c *TokenFilterLimitTokenCount) ConsumeAllTokens(consumeAllTokens bool) *TokenFilterLimitTokenCount {
	c.consumeAllTokens = &consumeAllTokens
	return c
}

// Validate validates TokenFilterLimitTokenCount.
func (c *TokenFilterLimitTokenCount) Validate(includeName bool) error {
	var invalid []string
	if includeName && c.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (c *TokenFilterLimitTokenCount) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "limit",
	// 		"max_token_count": 5,
	// 		"consume_all_tokens": true
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "limit"

	if c.maxTokenCount != nil {
		options["max_token_count"] = c.maxTokenCount
	}
	if c.consumeAllTokens != nil {
		options["consume_all_tokens"] = c.consumeAllTokens
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[c.name] = options
	return source, nil
}
