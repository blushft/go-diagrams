package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/openstack/storage"}

func (c *storageContainer) Cinder(opts ...attr.Attribute) *node.Node {
	return node.New("cinder", attr.AssetImage("assets/openstack/storage/cinder.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Manila(opts ...attr.Attribute) *node.Node {
	return node.New("manila", attr.AssetImage("assets/openstack/storage/manila.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Swift(opts ...attr.Attribute) *node.Node {
	return node.New("swift", attr.AssetImage("assets/openstack/storage/swift.png"), attr.Shape(attr.None))
}
