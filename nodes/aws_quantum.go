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
	opts = append(opts, attr.AssetImage("assets/aws/quantum/braket.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("braket", opts...)
}

func init() {
  const namespace= "aws.quantum"
  Register(namespace, "Braket", Quantum.Braket)
}
