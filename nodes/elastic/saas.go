package elastic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasContainer struct {
	path  string
	attrs []attr.Attribute
}

var Saas = &saasContainer{path: "assets/elastic/saas"}

func (c *saasContainer) Cloud(opts ...attr.Attribute) *node.Node {
	return node.New("cloud", attr.AssetImage("assets/elastic/saas/cloud.png"), attr.Shape(attr.None))
}

func (c *saasContainer) Elastic(opts ...attr.Attribute) *node.Node {
	return node.New("elastic", attr.AssetImage("assets/elastic/saas/elastic.png"), attr.Shape(attr.None))
}
