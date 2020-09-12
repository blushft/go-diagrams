package openstack

import "github.com/blushft/go-diagrams/node"

type operationsContainer struct {
	path string
	opts []node.Option
}

var Operations = &operationsContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/operations",
}
