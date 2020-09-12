package onprem

import "github.com/blushft/go-diagrams/node"

type containerContainer struct {
	path string
	opts []node.Option
}

var Container = &containerContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/container",
}

func (c *containerContainer) Rkt(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/container/rkt.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *containerContainer) Docker(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/container/docker.png")}, c.opts, opts)
	return node.New(nopts...)
}
