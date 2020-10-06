package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type arContainer struct {
	path  string
	attrs []attr.Attribute
}

var Ar = &arContainer{path: "assets/aws/ar"}

func (c *arContainer) Sumerian(opts ...attr.Attribute) *node.Node {
	return node.New("sumerian", attr.AssetImage("assets/aws/ar/sumerian.png"), attr.Shape(attr.None))
}
