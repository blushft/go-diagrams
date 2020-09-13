package openstack

import "github.com/blushft/go-diagrams/diagram"

type workloadprovisioningContainer struct {
	path string
	opts []diagram.NodeOption
}

var Workloadprovisioning = &workloadprovisioningContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/workloadprovisioning",
}

func (c *workloadprovisioningContainer) Magnum(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/workloadprovisioning/magnum.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *workloadprovisioningContainer) Sahara(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/workloadprovisioning/sahara.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *workloadprovisioningContainer) Trove(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/workloadprovisioning/trove.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
