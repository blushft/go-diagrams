package onprem

import "github.com/blushft/go-diagrams/node"

type cdContainer struct {
	path string
	opts []node.Option
}

var Cd = &cdContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/cd",
}

func (c *cdContainer) Spinnaker(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/cd/spinnaker.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *cdContainer) TektonCli(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/cd/tekton-cli.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *cdContainer) Tekton(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/cd/tekton.png")}, c.opts, opts)
	return node.New(nopts...)
}
