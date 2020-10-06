package openstack

import attr "github.com/blushft/go-diagrams/attr"

type lifecyclemanagementContainer struct {
	path  string
	attrs []attr.Attribute
}

var Lifecyclemanagement = &lifecyclemanagementContainer{path: "assets/openstack/lifecyclemanagement"}
