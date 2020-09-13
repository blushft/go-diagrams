package k8s

import "github.com/blushft/go-diagrams/diagram"

type infraContainer struct {
	path string
	opts []diagram.NodeOption
}

var Infra = &infraContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/infra",
}

func (c *infraContainer) Etcd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/infra/etcd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infraContainer) Master(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/infra/master.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infraContainer) Node(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/infra/node.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
