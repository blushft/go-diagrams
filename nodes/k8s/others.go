package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type othersContainer struct {
	path  string
	attrs []attr.Attribute
}

var Others = &othersContainer{path: "assets/k8s/others"}

func (c *othersContainer) Crd(opts ...attr.Attribute) *node.Node {
	return node.New("crd", attr.AssetImage("assets/k8s/others/crd.png"), attr.Shape(attr.None))
}

func (c *othersContainer) Psp(opts ...attr.Attribute) *node.Node {
	return node.New("psp", attr.AssetImage("assets/k8s/others/psp.png"), attr.Shape(attr.None))
}
