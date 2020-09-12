package aws

import "github.com/blushft/go-diagrams/node"

type arContainer struct {
	path string
	opts []node.Option
}

var Ar = &arContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/ar",
}

func (c *arContainer) Sumerian(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ar/sumerian.png")}, c.opts, opts)
	return node.New(nopts...)
}
