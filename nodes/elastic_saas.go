package nodes

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
	opts = append(opts, attr.AssetImage("assets/elastic/saas/cloud.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud", opts...)
}

func (c *saasContainer) Elastic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/saas/elastic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic", opts...)
}
