package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type podconfigContainer struct {
	path  string
	attrs []attr.Attribute
}

var Podconfig = &podconfigContainer{path: "assets/k8s/podconfig"}

func (c *podconfigContainer) Secret(opts ...attr.Attribute) *node.Node {
	return node.New("secret", attr.AssetImage("assets/k8s/podconfig/secret.png"), attr.Shape(attr.None))
}

func (c *podconfigContainer) Cm(opts ...attr.Attribute) *node.Node {
	return node.New("cm", attr.AssetImage("assets/k8s/podconfig/cm.png"), attr.Shape(attr.None))
}
