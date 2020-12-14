package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkingContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackNetworking =&networkingContainer{path: "assets/openstack/networking"}

func (c *networkingContainer) Neutron(opts ...attr.Attribute) *node.Node {
	return node.New("neutron", attr.AssetImage("assets/openstack/networking/neutron.png"), attr.Shape(attr.None))
}

func (c *networkingContainer) Octavia(opts ...attr.Attribute) *node.Node {
	return node.New("octavia", attr.AssetImage("assets/openstack/networking/octavia.png"), attr.Shape(attr.None))
}

func (c *networkingContainer) Designate(opts ...attr.Attribute) *node.Node {
	return node.New("designate", attr.AssetImage("assets/openstack/networking/designate.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.networking"
  Register(namespace, "Neutron", OpenstackNetworking.Neutron)
  Register(namespace, "Octavia", OpenstackNetworking.Octavia)
  Register(namespace, "Designate", OpenstackNetworking.Designate)
}
