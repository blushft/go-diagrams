package openstack

import "github.com/blushft/go-diagrams/diagram"

type billingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Billing = &billingContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/billing",
}

func (c *billingContainer) Cloudkitty(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/billing/cloudkitty.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
