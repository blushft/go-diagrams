package openstack

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/compute",
}

func (c *computeContainer) Nova(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/compute/nova.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Qinling(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/compute/qinling.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Zun(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/compute/zun.png")}, c.opts, opts)
	return node.New(nopts...)
}
