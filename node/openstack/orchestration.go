package openstack

import "github.com/blushft/go-diagrams/node"

type orchestrationContainer struct {
	path string
	opts []node.Option
}

var Orchestration = &orchestrationContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/orchestration",
}

func (c *orchestrationContainer) Blazar(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/orchestration/blazar.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *orchestrationContainer) Heat(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/orchestration/heat.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *orchestrationContainer) Mistral(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/orchestration/mistral.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *orchestrationContainer) Senlin(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/orchestration/senlin.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *orchestrationContainer) Zaqar(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/orchestration/zaqar.png")}, c.opts, opts)
	return node.New(nopts...)
}
