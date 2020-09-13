package generic

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/compute",
}

func (c *computeContainer) Rack(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/compute/rack.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
