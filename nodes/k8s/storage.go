package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/k8s/storage"}

func (c *storageContainer) Vol(opts ...attr.Attribute) *node.Node {
	return node.New("vol", attr.AssetImage("assets/k8s/storage/vol.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Pv(opts ...attr.Attribute) *node.Node {
	return node.New("pv", attr.AssetImage("assets/k8s/storage/pv.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Pvc(opts ...attr.Attribute) *node.Node {
	return node.New("pvc", attr.AssetImage("assets/k8s/storage/pvc.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Sc(opts ...attr.Attribute) *node.Node {
	return node.New("sc", attr.AssetImage("assets/k8s/storage/sc.png"), attr.Shape(attr.None))
}
