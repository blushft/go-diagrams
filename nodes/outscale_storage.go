package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type outscaleStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var OutscaleStorage =&outscaleStorageContainer{path: "assets/outscale/storage"}

func (c *outscaleStorageContainer) SimpleStorageService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/storage/simple-storage-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("simple-storage-service", opts...)
}

func (c *outscaleStorageContainer) Storage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/storage/storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage", opts...)
}

func init() {
  const namespace = "outscale.storage"
  Register(namespace, "SimpleStorageService", OutscaleStorage.SimpleStorageService)
  Register(namespace, "Storage", OutscaleStorage.Storage)
}
