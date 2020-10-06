package outscale

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/outscale/compute"}

func (c *computeContainer) Compute(opts ...attr.Attribute) *node.Node {
	return node.New("compute", attr.AssetImage("assets/outscale/compute/compute.png"), attr.Shape(attr.None))
}

func (c *computeContainer) DirectConnect(opts ...attr.Attribute) *node.Node {
	return node.New("direct-connect", attr.AssetImage("assets/outscale/compute/direct-connect.png"), attr.Shape(attr.None))
}
