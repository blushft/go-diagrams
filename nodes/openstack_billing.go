package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type billingContainer struct {
	path  string
	attrs []attr.Attribute
}

var Billing = &billingContainer{path: "assets/openstack/billing"}

func (c *billingContainer) Cloudkitty(opts ...attr.Attribute) *node.Node {
	return node.New("cloudkitty", attr.AssetImage("assets/openstack/billing/cloudkitty.png"), attr.Shape(attr.None))
}
