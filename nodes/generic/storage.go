package generic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/generic/storage"}

func (c *storageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/generic/storage/storage.png"), attr.Shape(attr.None))
}
