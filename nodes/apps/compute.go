package apps

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/compute",
}

func (c *computeContainer) Server(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/compute/server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Nomad(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/compute/nomad.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
