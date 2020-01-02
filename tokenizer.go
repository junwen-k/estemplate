// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// Tokenizer represents the generic tokenizer interface.
// A tokenizer's only purpose is to return the source of the
// property in a mapping template as a JSON-serializable object.
// Returning a map[string]interface{} will do.
type Tokenizer interface {
	Name() string
	Source(includeName bool) (interface{}, error)
}
