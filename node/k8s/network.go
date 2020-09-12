package k8s

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/network",
}

func (c *networkContainer) Ep(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/network/ep.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Ing(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/network/ing.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Netpol(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/network/netpol.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Svc(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/network/svc.png")}, c.opts, opts)
	return node.New(nopts...)
}
