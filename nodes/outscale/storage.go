package outscale

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/outscale/storage"}

func (c *storageContainer) SimpleStorageService(opts ...attr.Attribute) *node.Node {
	return node.New("simple-storage-service", attr.AssetImage("assets/outscale/storage/simple-storage-service.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/outscale/storage/storage.png"), attr.Shape(attr.None))
}
