package openstack

import "github.com/blushft/go-diagrams/node"

type applicationlifecycleContainer struct {
	path string
	opts []node.Option
}

var Applicationlifecycle = &applicationlifecycleContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/applicationlifecycle",
}

func (c *applicationlifecycleContainer) Murano(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/applicationlifecycle/murano.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationlifecycleContainer) Solum(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/applicationlifecycle/solum.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationlifecycleContainer) Freezer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/applicationlifecycle/freezer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationlifecycleContainer) Masakari(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/applicationlifecycle/masakari.png")}, c.opts, opts)
	return node.New(nopts...)
}
