package openstack

import "github.com/blushft/go-diagrams/diagram"

type multiregionContainer struct {
	path string
	opts []diagram.NodeOption
}

var Multiregion = &multiregionContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/multiregion",
}

func (c *multiregionContainer) Tricircle(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/multiregion/tricircle.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
