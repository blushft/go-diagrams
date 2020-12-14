package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type genericStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var GenericStorage = &genericStorageContainer{path: "assets/generic/storage"}

func (c *genericStorageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/generic/storage/storage.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.storage"
	Register(namespace, "GenericStorage", GenericStorage.Storage)
}
