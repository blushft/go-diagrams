package k8s

import "github.com/blushft/go-diagrams/node"

type othersContainer struct {
	path string
	opts []node.Option
}

var Others = &othersContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/others",
}

func (c *othersContainer) Crd(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/others/crd.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *othersContainer) Psp(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/others/psp.png")}, c.opts, opts)
	return node.New(nopts...)
}
