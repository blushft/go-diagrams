package openstack

import "github.com/blushft/go-diagrams/diagram"

type lifecyclemanagementContainer struct {
	path string
	opts []diagram.NodeOption
}

var Lifecyclemanagement = &lifecyclemanagementContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/lifecyclemanagement",
}
