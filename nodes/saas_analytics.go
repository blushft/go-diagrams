package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasAnalyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasAnalytics =&saasAnalyticsContainer{path: "assets/saas/analytics"}

func (c *saasAnalyticsContainer) Snowflake(opts ...attr.Attribute) *node.Node {
	return node.New("snowflake", attr.AssetImage("assets/saas/analytics/snowflake.png"), attr.Shape(attr.None))
}

func (c *saasAnalyticsContainer) Stitch(opts ...attr.Attribute) *node.Node {
	return node.New("stitch", attr.AssetImage("assets/saas/analytics/stitch.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.analytics"
  Register(namespace, "Snowflake", SaasAnalytics.Snowflake)
  Register(namespace, "Stitch", SaasAnalytics.Stitch)
}
