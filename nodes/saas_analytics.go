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
	opts = append(opts, attr.AssetImage("assets/saas/analytics/snowflake.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("snowflake", opts...)
}

func (c *saasAnalyticsContainer) Stitch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/saas/analytics/stitch.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("stitch", opts...)
}

func init() {
  const namespace = "saas.analytics"
  Register(namespace, "Snowflake", SaasAnalytics.Snowflake)
  Register(namespace, "Stitch", SaasAnalytics.Stitch)
}
