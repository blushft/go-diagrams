package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type genericComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var GenericCompute = &genericComputeContainer{path: "assets/generic/compute"}

func (c *genericComputeContainer) Rack(opts ...attr.Attribute) *node.Node {
	return node.New("rack", attr.AssetImage("assets/generic/compute/rack.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.etl"
	Register(namespace, "Rack", GenericCompute.Rack)
}