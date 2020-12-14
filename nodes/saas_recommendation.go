package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type recommendationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Recommendation = &recommendationContainer{path: "assets/saas/recommendation"}

func (c *recommendationContainer) Recombee(opts ...attr.Attribute) *node.Node {
	return node.New("recombee", attr.AssetImage("assets/saas/recommendation/recombee.png"), attr.Shape(attr.None))
}
