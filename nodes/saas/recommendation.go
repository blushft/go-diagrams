package saas

import "github.com/blushft/go-diagrams/diagram"

type recommendationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Recommendation = &recommendationContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/recommendation",
}

func (c *recommendationContainer) Recombee(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/recommendation/recombee.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
