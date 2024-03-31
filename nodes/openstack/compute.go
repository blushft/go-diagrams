package openstack

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/compute",
}

func (c *computeContainer) Nova(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/compute/nova.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Qinling(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/compute/qinling.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Zun(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/compute/zun.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
