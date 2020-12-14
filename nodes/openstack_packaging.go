package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type openstackPackagingContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackPackaging =&openstackPackagingContainer{path: "assets/openstack/packaging"}

func (c *openstackPackagingContainer) Loci(opts ...attr.Attribute) *node.Node {
	return node.New("loci", attr.AssetImage("assets/openstack/packaging/loci.png"), attr.Shape(attr.None))
}

func (c *openstackPackagingContainer) Puppet(opts ...attr.Attribute) *node.Node {
	return node.New("puppet", attr.AssetImage("assets/openstack/packaging/puppet.png"), attr.Shape(attr.None))
}

func (c *openstackPackagingContainer) Rpm(opts ...attr.Attribute) *node.Node {
	return node.New("rpm", attr.AssetImage("assets/openstack/packaging/rpm.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.packaging"
  Register(namespace, "Loci", OpenstackPackaging.Loci)
  Register(namespace, "Puppet", OpenstackPackaging.Puppet)
  Register(namespace, "Rpm", OpenstackPackaging.Rpm)
}
