package outscale

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("outscale"), node.Shape("none")},
	path: "assets/outscale/compute",
}

func (c *computeContainer) DirectConnect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/compute/direct-connect.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Compute(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/compute/compute.png")}, c.opts, opts)
	return node.New(nopts...)
}
