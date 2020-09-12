package k8s

import "github.com/blushft/go-diagrams/node"

type ecosystemContainer struct {
	path string
	opts []node.Option
}

var Ecosystem = &ecosystemContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/ecosystem",
}

func (c *ecosystemContainer) Helm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/ecosystem/helm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ecosystemContainer) Krew(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/ecosystem/krew.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ecosystemContainer) Kustomize(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/ecosystem/kustomize.png")}, c.opts, opts)
	return node.New(nopts...)
}
