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
	opts = append(opts, attr.AssetImage("assets/openstack/networking/neutron.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("neutron", opts...)
}

func (c *networkingContainer) Octavia(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/networking/octavia.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("octavia", opts...)
}

func (c *networkingContainer) Designate(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/networking/designate.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("designate", opts...)
}

func init() {
  const namespace = "openstack.networking"
  Register(namespace, "Neutron", OpenstackNetworking.Neutron)
  Register(namespace, "Octavia", OpenstackNetworking.Octavia)
  Register(namespace, "Designate", OpenstackNetworking.Designate)
}
