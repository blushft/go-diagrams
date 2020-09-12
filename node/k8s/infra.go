package k8s

import "github.com/blushft/go-diagrams/node"

type infraContainer struct {
	path string
	opts []node.Option
}

var Infra = &infraContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/infra",
}

func (c *infraContainer) Node(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/infra/node.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *infraContainer) Etcd(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/infra/etcd.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *infraContainer) Master(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/infra/master.png")}, c.opts, opts)
	return node.New(nopts...)
}
