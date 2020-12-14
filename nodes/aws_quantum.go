package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type quantumContainer struct {
	path  string
	attrs []attr.Attribute
}

var Quantum = &quantumContainer{path: "assets/aws/quantum"}

func (c *quantumContainer) Braket(opts ...attr.Attribute) *node.Node {
	return node.New("braket", attr.AssetImage("assets/aws/quantum/braket.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.quantum"
  Register(namespace, "Braket", Quantum.Braket)
}
