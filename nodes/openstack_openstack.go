package nodes

import attr "github.com/blushft/go-diagrams/attr"

type openstackContainer struct {
	path  string
	attrs []attr.Attribute
}

var Openstack = &openstackContainer{path: "assets/openstack/openstack"}
