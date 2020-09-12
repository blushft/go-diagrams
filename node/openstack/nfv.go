package openstack

import "github.com/blushft/go-diagrams/node"

type nfvContainer struct {
	path string
	opts []node.Option
}

var Nfv = &nfvContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/nfv",
}

func (c *nfvContainer) Tacker(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/adjacentenablers/nfv/tacker.png")}, c.opts, opts)
	return node.New(nopts...)
}
