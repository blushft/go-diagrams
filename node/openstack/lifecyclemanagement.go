package openstack

import "github.com/blushft/go-diagrams/node"

type lifecyclemanagementContainer struct {
	path string
	opts []node.Option
}

var Lifecyclemanagement = &lifecyclemanagementContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/lifecyclemanagement",
}
