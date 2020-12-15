package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackStorage =&storageContainer{path: "assets/openstack/storage"}

func (c *storageContainer) Cinder(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/storage/cinder.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cinder", opts...)
}

func (c *storageContainer) Manila(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/storage/manila.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("manila", opts...)
}

func (c *storageContainer) Swift(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/storage/swift.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("swift", opts...)
}

func init() {
  const namespace = "openstack.storage"
  Register(namespace, "Cinder", OpenstackStorage.Cinder)
  Register(namespace, "Manila", OpenstackStorage.Manila)
  Register(namespace, "Swift", OpenstackStorage.Swift)
}
