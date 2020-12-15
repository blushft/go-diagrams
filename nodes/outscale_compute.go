package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type outscaleComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var OutscaleCompute =&outscaleComputeContainer{path: "assets/outscale/compute"}

func (c *outscaleComputeContainer) Compute(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/compute/compute.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("compute", opts...)
}

func (c *outscaleComputeContainer) DirectConnect(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/compute/direct-connect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("direct-connect", opts...)
}

func init() {
  const namespace = "outscale.compute"
  Register(namespace, "Compute", OutscaleCompute.Compute)
  Register(namespace, "DirectConnect", OutscaleCompute.DirectConnect)
}
