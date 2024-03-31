package openstack

import "github.com/blushft/go-diagrams/diagram"

type orchestrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Orchestration = &orchestrationContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/orchestration",
}

func (c *orchestrationContainer) Blazar(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/orchestration/blazar.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *orchestrationContainer) Heat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/orchestration/heat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *orchestrationContainer) Mistral(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/orchestration/mistral.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *orchestrationContainer) Senlin(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/orchestration/senlin.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *orchestrationContainer) Zaqar(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/orchestration/zaqar.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
