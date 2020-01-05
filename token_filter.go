// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// TokenFilter represents the generic token filter interface.
// A token filter's only purpose is to return the source of the
// property in a mapping template as a JSON-serializable object.
// Returning a map[string]interface{} will do.
type TokenFilter interface {
	Name() string
	Source(includeName bool) (interface{}, error)
}
