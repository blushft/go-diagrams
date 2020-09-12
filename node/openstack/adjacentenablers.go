package openstack

import "github.com/blushft/go-diagrams/node"

type adjacentenablersContainer struct {
	path string
	opts []node.Option
}

var Adjacentenablers = &adjacentenablersContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/adjacentenablers",
}
