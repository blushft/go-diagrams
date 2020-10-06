package generic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/generic/compute"}

func (c *computeContainer) Rack(opts ...attr.Attribute) *node.Node {
	return node.New("rack", attr.AssetImage("assets/generic/compute/rack.png"), attr.Shape(attr.None))
}
