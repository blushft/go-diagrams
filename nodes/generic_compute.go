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
	opts = append(opts, attr.AssetImage("assets/generic/compute/rack.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rack", opts...)
}

func init() {
	const namespace = "apps.etl"
	Register(namespace, "Rack", GenericCompute.Rack)
}