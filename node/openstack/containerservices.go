package openstack

import "github.com/blushft/go-diagrams/node"

type containerservicesContainer struct {
	path string
	opts []node.Option
}

var Containerservices = &containerservicesContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/containerservices",
}

func (c *containerservicesContainer) Kuryr(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/adjacentenablers/containerservices/kuryr.png")}, c.opts, opts)
	return node.New(nopts...)
}
