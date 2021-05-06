package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/apps/compute"}

func (c *computeContainer) Server(opts ...attr.Attribute) *node.Node {
	return node.New("server", attr.AssetImage("assets/apps/compute/server.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Nomad(opts ...attr.Attribute) *node.Node {
	return node.New("nomad", attr.AssetImage("assets/apps/compute/nomad.png"), attr.Shape(attr.None))
}
