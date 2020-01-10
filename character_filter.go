// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// CharacterFilter represents the generic character filter interface.
// A character filter's only purpose is to return the source of the
// property in a mapping template as a JSON-serializable object.
// Returning a map[string]interface{} will do.
type CharacterFilter interface {
	Name() string
	Source(includeName bool) (interface{}, error)
}
