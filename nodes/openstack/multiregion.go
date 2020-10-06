package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type multiregionContainer struct {
	path  string
	attrs []attr.Attribute
}

var Multiregion = &multiregionContainer{path: "assets/openstack/multiregion"}

func (c *multiregionContainer) Tricircle(opts ...attr.Attribute) *node.Node {
	return node.New("tricircle", attr.AssetImage("assets/openstack/multiregion/tricircle.png"), attr.Shape(attr.None))
}
