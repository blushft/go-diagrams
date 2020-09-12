package openstack

import "github.com/blushft/go-diagrams/node"

type sharedservicesContainer struct {
	path string
	opts []node.Option
}

var Sharedservices = &sharedservicesContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/sharedservices",
}

func (c *sharedservicesContainer) Barbican(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/sharedservices/barbican.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *sharedservicesContainer) Glance(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/sharedservices/glance.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *sharedservicesContainer) Karbor(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/sharedservices/karbor.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *sharedservicesContainer) Keystone(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/sharedservices/keystone.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *sharedservicesContainer) Searchlight(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/sharedservices/searchlight.png")}, c.opts, opts)
	return node.New(nopts...)
}
