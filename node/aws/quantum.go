package aws

import "github.com/blushft/go-diagrams/node"

type quantumContainer struct {
	path string
	opts []node.Option
}

var Quantum = &quantumContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/quantum",
}

func (c *quantumContainer) Braket(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/quantum/braket.png")}, c.opts, opts)
	return node.New(nopts...)
}
