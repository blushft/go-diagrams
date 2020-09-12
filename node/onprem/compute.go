package onprem

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/compute",
}

func (c *computeContainer) Nomad(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/compute/nomad.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Server(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/compute/server.png")}, c.opts, opts)
	return node.New(nopts...)
}
