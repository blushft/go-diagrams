package openstack

import "github.com/blushft/go-diagrams/diagram"

type sharedservicesContainer struct {
	path string
	opts []diagram.NodeOption
}

var Sharedservices = &sharedservicesContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/sharedservices",
}

func (c *sharedservicesContainer) Glance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/sharedservices/glance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *sharedservicesContainer) Karbor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/sharedservices/karbor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *sharedservicesContainer) Keystone(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/sharedservices/keystone.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *sharedservicesContainer) Searchlight(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/sharedservices/searchlight.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *sharedservicesContainer) Barbican(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/sharedservices/barbican.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
