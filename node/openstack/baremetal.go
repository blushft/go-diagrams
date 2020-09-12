package openstack

import "github.com/blushft/go-diagrams/node"

type baremetalContainer struct {
	path string
	opts []node.Option
}

var Baremetal = &baremetalContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/baremetal",
}

func (c *baremetalContainer) Ironic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/baremetal/ironic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *baremetalContainer) Cyborg(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/baremetal/cyborg.png")}, c.opts, opts)
	return node.New(nopts...)
}
