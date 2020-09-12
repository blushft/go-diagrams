package k8s

import "github.com/blushft/go-diagrams/node"

type podconfigContainer struct {
	path string
	opts []node.Option
}

var Podconfig = &podconfigContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/podconfig",
}

func (c *podconfigContainer) Cm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/podconfig/cm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *podconfigContainer) Secret(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/podconfig/secret.png")}, c.opts, opts)
	return node.New(nopts...)
}
