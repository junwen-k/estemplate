// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"fmt"
	"strings"
)

// RoutingAllocation shard allocation filters to control where Elasticsearch allocates shards
// of a particular index. These per-index filters are applied in conjuction with cluster-wide
// allocation filtering and allocation awareness.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/shard-allocation-filtering.html
// for details.
type RoutingAllocation struct {
	allocationType string
	attribute      string
	values         []string
}

// NewRoutingAllocation initializes a new RoutingAllocation.
func NewRoutingAllocation(allocationType, attribute string, values ...string) *RoutingAllocation {
	return &RoutingAllocation{
		allocationType: allocationType,
		attribute:      attribute,
		values:         values,
	}
}

// NewRoutingAllocationInclude initializes a new include type RoutingAllocation.
func NewRoutingAllocationInclude(attribute string, values ...string) *RoutingAllocation {
	return NewRoutingAllocation("include", attribute, values...)
}

// NewRoutingAllocationRequire initializes a new require type RoutingAllocation.
func NewRoutingAllocationRequire(attribute string, values ...string) *RoutingAllocation {
	return NewRoutingAllocation("require", attribute, values...)
}

// NewRoutingAllocationExclude initializes a new exclude type RoutingAllocation.
func NewRoutingAllocationExclude(attribute string, values ...string) *RoutingAllocation {
	return NewRoutingAllocation("exclude", attribute, values...)
}

// AllocationType sets the type of allocation for this routing allocaiton.
// Can be set to the following values:
// "include" - At least one of the comma-separated values.
// "require" - All of the comma-separated values.
// "exclude" - None of the comma-separated values.
func (a *RoutingAllocation) AllocationType(allocationType string) *RoutingAllocation {
	a.allocationType = allocationType
	return a
}

// Attribute sets the attribute for this routing allocation. The index allocation settings support
// the following built-in attributes:
// "_name": Match nodes by node name.
// "_host_ip": Match nodes by host IP address (IP associated with hostname).
// "_publish_ip": Match nodes by publish IP address.
// "_ip": Match either `_host_ip` or `_publish_ip`.
// "_host": Match nodes by hostname.
// "_id": Match nodes by node id.
func (a *RoutingAllocation) Attribute(attribute string) *RoutingAllocation {
	a.attribute = attribute
	return a
}

// Values sets the attribute values for this routing allocation. Wildcards are supported. For example,
// "192.168.2.*" for ip attribute.
func (a *RoutingAllocation) Values(values ...string) *RoutingAllocation {
	a.values = append(a.values, values...)
	return a
}

// Source returns the serializable JSON for the source builder.
func (a *RoutingAllocation) Source(includeName bool) (interface{}, error) {
	// {
	// 	"routing": {
	// 		"allocation.include.size": "big,medium",
	// 		"allocation.include.rack": "rack1"
	// 	}
	// }
	options := make(map[string]interface{})

	if a.allocationType != "" && a.attribute != "" && len(a.values) > 0 {
		options[fmt.Sprintf("allocation.%s.%s", a.allocationType, a.attribute)] = strings.Join(a.values, ",")
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["routing"] = options
	return source, nil
}
