package generic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type deviceContainer struct {
	path  string
	attrs []attr.Attribute
}

var Device = &deviceContainer{path: "assets/generic/device"}

func (c *deviceContainer) Tablet(opts ...attr.Attribute) *node.Node {
	return node.New("tablet", attr.AssetImage("assets/generic/device/tablet.png"), attr.Shape(attr.None))
}

func (c *deviceContainer) Mobile(opts ...attr.Attribute) *node.Node {
	return node.New("mobile", attr.AssetImage("assets/generic/device/mobile.png"), attr.Shape(attr.None))
}
