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
	opts = append(opts, attr.AssetImage("assets/generic/storage/storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage", opts...)
}

func init() {
	const namespace = "generic.storage"
	Register(namespace, "GenericStorage", GenericStorage.Storage)
}
