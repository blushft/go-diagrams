package openstack

import "github.com/blushft/go-diagrams/diagram"

type nfvContainer struct {
	path string
	opts []diagram.NodeOption
}

var Nfv = &nfvContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/nfv",
}

func (c *nfvContainer) Tacker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/adjacentenablers/nfv/tacker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
