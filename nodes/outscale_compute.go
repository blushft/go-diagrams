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
	return node.New("compute", attr.AssetImage("assets/outscale/compute/compute.png"), attr.Shape(attr.None))
}

func (c *outscaleComputeContainer) DirectConnect(opts ...attr.Attribute) *node.Node {
	return node.New("direct-connect", attr.AssetImage("assets/outscale/compute/direct-connect.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "outscale.compute"
  Register(namespace, "Compute", OutscaleCompute.Compute)
  Register(namespace, "DirectConnect", OutscaleCompute.DirectConnect)
}
