package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type sharedservicesContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackSharedservices =&sharedservicesContainer{path: "assets/openstack/sharedservices"}

func (c *sharedservicesContainer) Barbican(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/sharedservices/barbican.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("barbican", opts...)
}

func (c *sharedservicesContainer) Glance(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/sharedservices/glance.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("glance", opts...)
}

func (c *sharedservicesContainer) Karbor(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/sharedservices/karbor.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("karbor", opts...)
}

func (c *sharedservicesContainer) Keystone(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/sharedservices/keystone.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("keystone", opts...)
}

func (c *sharedservicesContainer) Searchlight(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/sharedservices/searchlight.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("searchlight", opts...)
}

func init() {
  const namespace = "openstack.sharedservices"
  Register(namespace, "Barbican", OpenstackSharedservices.Barbican)
  Register(namespace, "Glance", OpenstackSharedservices.Glance)
  Register(namespace, "Karbor", OpenstackSharedservices.Karbor)
  Register(namespace, "Keystone", OpenstackSharedservices.Keystone)
  Register(namespace, "Searchlight", OpenstackSharedservices.Searchlight)
}
