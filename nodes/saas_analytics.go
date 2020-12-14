package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type analyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Analytics = &analyticsContainer{path: "assets/saas/analytics"}

func (c *analyticsContainer) Snowflake(opts ...attr.Attribute) *node.Node {
	return node.New("snowflake", attr.AssetImage("assets/saas/analytics/snowflake.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Stitch(opts ...attr.Attribute) *node.Node {
	return node.New("stitch", attr.AssetImage("assets/saas/analytics/stitch.png"), attr.Shape(attr.None))
}
