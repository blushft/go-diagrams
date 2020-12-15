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
	opts = append(opts, attr.AssetImage("assets/openstack/billing/cloudkitty.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudkitty", opts...)
}

func init() {
  const namespace = "openstack.billing"
  Register(namespace, "Cloudkitty", OpenstackBilling.Cloudkitty)
}
