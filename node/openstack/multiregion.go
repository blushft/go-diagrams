package openstack

import "github.com/blushft/go-diagrams/node"

type multiregionContainer struct {
	path string
	opts []node.Option
}

var Multiregion = &multiregionContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/multiregion",
}

func (c *multiregionContainer) Tricircle(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/multiregion/tricircle.png")}, c.opts, opts)
	return node.New(nopts...)
}
