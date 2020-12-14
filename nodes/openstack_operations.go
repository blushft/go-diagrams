package nodes

import attr "github.com/blushft/go-diagrams/attr"

type operationsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Operations = &operationsContainer{path: "assets/openstack/operations"}
