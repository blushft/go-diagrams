package openstack

import "github.com/blushft/go-diagrams/node"

type billingContainer struct {
	path string
	opts []node.Option
}

var Billing = &billingContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/billing",
}

func (c *billingContainer) Cloudkitty(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/billing/cloudkitty.png")}, c.opts, opts)
	return node.New(nopts...)
}
