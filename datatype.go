// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// Datatype represents the generic datatype interface.
// A datatype's only purpose is to return the source of the
// property in a mapping template as a JSON-serializable object.
// Returning a map[string]interface{} will do.
type Datatype interface {
	Name() string
	Source(includeName bool) (interface{}, error)
}
