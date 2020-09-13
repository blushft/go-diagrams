package aws

import "github.com/blushft/go-diagrams/diagram"

type quantumContainer struct {
	path string
	opts []diagram.NodeOption
}

var Quantum = &quantumContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/quantum",
}

func (c *quantumContainer) Braket(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/quantum/braket.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
