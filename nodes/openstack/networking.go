package openstack

import "github.com/blushft/go-diagrams/diagram"

type networkingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Networking = &networkingContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/networking",
}

func (c *networkingContainer) Designate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/networking/designate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkingContainer) Neutron(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/networking/neutron.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkingContainer) Octavia(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/networking/octavia.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
