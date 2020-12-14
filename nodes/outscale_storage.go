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
	return node.New("simple-storage-service", attr.AssetImage("assets/outscale/storage/simple-storage-service.png"), attr.Shape(attr.None))
}

func (c *outscaleStorageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/outscale/storage/storage.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "outscale.storage"
  Register(namespace, "SimpleStorageService", OutscaleStorage.SimpleStorageService)
  Register(namespace, "Storage", OutscaleStorage.Storage)
}
