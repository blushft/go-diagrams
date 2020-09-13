package openstack

import "github.com/blushft/go-diagrams/diagram"

type applicationlifecycleContainer struct {
	path string
	opts []diagram.NodeOption
}

var Applicationlifecycle = &applicationlifecycleContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/applicationlifecycle",
}

func (c *applicationlifecycleContainer) Murano(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/applicationlifecycle/murano.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationlifecycleContainer) Solum(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/applicationlifecycle/solum.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationlifecycleContainer) Freezer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/applicationlifecycle/freezer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationlifecycleContainer) Masakari(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/applicationlifecycle/masakari.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
