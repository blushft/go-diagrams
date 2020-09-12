package saas

import "github.com/blushft/go-diagrams/node"

type recommendationContainer struct {
	path string
	opts []node.Option
}

var Recommendation = &recommendationContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/recommendation",
}

func (c *recommendationContainer) Recombee(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/recommendation/recombee.png")}, c.opts, opts)
	return node.New(nopts...)
}
