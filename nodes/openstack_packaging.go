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
	opts = append(opts, attr.AssetImage("assets/openstack/packaging/loci.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("loci", opts...)
}

func (c *openstackPackagingContainer) Puppet(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/packaging/puppet.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("puppet", opts...)
}

func (c *openstackPackagingContainer) Rpm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/packaging/rpm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rpm", opts...)
}

func init() {
  const namespace = "openstack.packaging"
  Register(namespace, "Loci", OpenstackPackaging.Loci)
  Register(namespace, "Puppet", OpenstackPackaging.Puppet)
  Register(namespace, "Rpm", OpenstackPackaging.Rpm)
}
