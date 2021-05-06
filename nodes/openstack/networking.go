package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkingContainer struct {
	path  string
	attrs []attr.Attribute
}

var Networking = &networkingContainer{path: "assets/openstack/networking"}

func (c *networkingContainer) Octavia(opts ...attr.Attribute) *node.Node {
	return node.New("octavia", attr.AssetImage("assets/openstack/networking/octavia.png"), attr.Shape(attr.None))
}

func (c *networkingContainer) Designate(opts ...attr.Attribute) *node.Node {
	return node.New("designate", attr.AssetImage("assets/openstack/networking/designate.png"), attr.Shape(attr.None))
}

func (c *networkingContainer) Neutron(opts ...attr.Attribute) *node.Node {
	return node.New("neutron", attr.AssetImage("assets/openstack/networking/neutron.png"), attr.Shape(attr.None))
}
