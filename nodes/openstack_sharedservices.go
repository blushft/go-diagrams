package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type sharedservicesContainer struct {
	path  string
	attrs []attr.Attribute
}

var Sharedservices = &sharedservicesContainer{path: "assets/openstack/sharedservices"}

func (c *sharedservicesContainer) Barbican(opts ...attr.Attribute) *node.Node {
	return node.New("barbican", attr.AssetImage("assets/openstack/sharedservices/barbican.png"), attr.Shape(attr.None))
}

func (c *sharedservicesContainer) Glance(opts ...attr.Attribute) *node.Node {
	return node.New("glance", attr.AssetImage("assets/openstack/sharedservices/glance.png"), attr.Shape(attr.None))
}

func (c *sharedservicesContainer) Karbor(opts ...attr.Attribute) *node.Node {
	return node.New("karbor", attr.AssetImage("assets/openstack/sharedservices/karbor.png"), attr.Shape(attr.None))
}

func (c *sharedservicesContainer) Keystone(opts ...attr.Attribute) *node.Node {
	return node.New("keystone", attr.AssetImage("assets/openstack/sharedservices/keystone.png"), attr.Shape(attr.None))
}

func (c *sharedservicesContainer) Searchlight(opts ...attr.Attribute) *node.Node {
	return node.New("searchlight", attr.AssetImage("assets/openstack/sharedservices/searchlight.png"), attr.Shape(attr.None))
}
