package openstack

import "github.com/blushft/go-diagrams/node"

type workloadprovisioningContainer struct {
	path string
	opts []node.Option
}

var Workloadprovisioning = &workloadprovisioningContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/workloadprovisioning",
}

func (c *workloadprovisioningContainer) Magnum(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/workloadprovisioning/magnum.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *workloadprovisioningContainer) Sahara(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/workloadprovisioning/sahara.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *workloadprovisioningContainer) Trove(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/workloadprovisioning/trove.png")}, c.opts, opts)
	return node.New(nopts...)
}
