package aws

import "github.com/blushft/go-diagrams/diagram"

type arContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ar = &arContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/ar",
}

func (c *arContainer) Sumerian(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ar/sumerian.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
