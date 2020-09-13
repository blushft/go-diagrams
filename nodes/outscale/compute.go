package outscale

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("outscale"), diagram.NodeShape("none")},
	path: "assets/outscale/compute",
}

func (c *computeContainer) Compute(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/compute/compute.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) DirectConnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/compute/direct-connect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
