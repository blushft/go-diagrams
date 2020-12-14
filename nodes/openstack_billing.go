package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type billingContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackBilling =&billingContainer{path: "assets/openstack/billing"}

func (c *billingContainer) Cloudkitty(opts ...attr.Attribute) *node.Node {
	return node.New("cloudkitty", attr.AssetImage("assets/openstack/billing/cloudkitty.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.billing"
  Register(namespace, "Cloudkitty", OpenstackBilling.Cloudkitty)
}
