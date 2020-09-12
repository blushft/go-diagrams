package openstack

import "github.com/blushft/go-diagrams/node"

type networkingContainer struct {
	path string
	opts []node.Option
}

var Networking = &networkingContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/networking",
}

func (c *networkingContainer) Neutron(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/networking/neutron.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkingContainer) Octavia(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/networking/octavia.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkingContainer) Designate(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/networking/designate.png")}, c.opts, opts)
	return node.New(nopts...)
}
