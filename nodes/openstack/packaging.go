package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type packagingContainer struct {
	path  string
	attrs []attr.Attribute
}

var Packaging = &packagingContainer{path: "assets/openstack/packaging"}

func (c *packagingContainer) Loci(opts ...attr.Attribute) *node.Node {
	return node.New("loci", attr.AssetImage("assets/openstack/packaging/loci.png"), attr.Shape(attr.None))
}

func (c *packagingContainer) Puppet(opts ...attr.Attribute) *node.Node {
	return node.New("puppet", attr.AssetImage("assets/openstack/packaging/puppet.png"), attr.Shape(attr.None))
}

func (c *packagingContainer) Rpm(opts ...attr.Attribute) *node.Node {
	return node.New("rpm", attr.AssetImage("assets/openstack/packaging/rpm.png"), attr.Shape(attr.None))
}
