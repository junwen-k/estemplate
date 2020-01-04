// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"fmt"
)

// TokenizerPathHierarchy Structured Text Tokenizer which takes a hierarchical value like a
// filesystem path, splits on the path separator, and emits a term for each component in the
// tree.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/analysis-pathhierarchy-tokenizer.html
// for details.
type TokenizerPathHierarchy struct {
	Tokenizer
	name string

	// fields specific to path hierarchy tokenizer
	delimiter   string
	replacement string
	bufferSize  *int
	reverse     *bool
	skip        *int
}

// NewTokenizerPathHierarchy initializes a new TokenizerPathHierarchy.
func NewTokenizerPathHierarchy(name string) *TokenizerPathHierarchy {
	return &TokenizerPathHierarchy{
		name: name,
	}
}

// Name returns field key for the Tokenizer.
func (h *TokenizerPathHierarchy) Name() string {
	return h.name
}

// Delimiter sets the character to use as the path separator.
// Defaults to "/".
func (h *TokenizerPathHierarchy) Delimiter(delimiter string) *TokenizerPathHierarchy {
	h.delimiter = delimiter
	return h
}

// Replacement sets an optional replacement character to use for the delimiter.
// Defaults to the `delimiter`.
func (h *TokenizerPathHierarchy) Replacement(replacement string) *TokenizerPathHierarchy {
	h.replacement = replacement
	return h
}

// BufferSize sets the number of characters read into the term buffer in a single
// pass. The term buffer will grow by this size until all the text has been consumed.
// It is advisable not to change this setting.
// Defaults to 1024.
func (h *TokenizerPathHierarchy) BufferSize(bufferSize int) *TokenizerPathHierarchy {
	h.bufferSize = &bufferSize
	return h
}

// Reverse sets whether to emits the tokens in reverse order or not.
// Defaults to false.
func (h *TokenizerPathHierarchy) Reverse(reverse bool) *TokenizerPathHierarchy {
	h.reverse = &reverse
	return h
}

// Skip sets the number of initial tokens to skip.
// Defaults to 0.
func (h *TokenizerPathHierarchy) Skip(skip int) *TokenizerPathHierarchy {
	h.skip = &skip
	return h
}

// Validate validates TokenizerPathHierarchy.
func (h *TokenizerPathHierarchy) Validate(includeName bool) error {
	var invalid []string
	if includeName && h.name == "" {
		invalid = append(invalid, "Name")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Source returns the serializable JSON for the source builder.
func (h *TokenizerPathHierarchy) Source(includeName bool) (interface{}, error) {
	// {
	// 	"test": {
	// 		"type": "path_hierarchy",
	// 		"delimiter": "-",
	// 		"replacement": "/",
	// 		"butter_size": 1024,
	// 		"reverse": true,
	// 		"skip": 2
	// 	}
	// }
	options := make(map[string]interface{})
	options["type"] = "path_hierarchy"

	if h.delimiter != "" {
		options["delimiter"] = h.delimiter
	}
	if h.replacement != "" {
		options["replacement"] = h.replacement
	}
	if h.bufferSize != nil {
		options["buffer_size"] = h.bufferSize
	}
	if h.reverse != nil {
		options["reverse"] = h.reverse
	}
	if h.skip != nil {
		options["skip"] = h.skip
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source[h.name] = options
	return source, nil
}
