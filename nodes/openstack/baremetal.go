package openstack

import "github.com/blushft/go-diagrams/diagram"

type baremetalContainer struct {
	path string
	opts []diagram.NodeOption
}

var Baremetal = &baremetalContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/baremetal",
}

func (c *baremetalContainer) Cyborg(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/baremetal/cyborg.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *baremetalContainer) Ironic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/baremetal/ironic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
