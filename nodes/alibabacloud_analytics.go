package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudAnalyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabaCloudAnalytics =&alibabaCloudAnalyticsContainer{path: "assets/alibabacloud/analytics"}

func (c *alibabaCloudAnalyticsContainer) DataLakeAnalytics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/analytics/data-lake-analytics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-lake-analytics", opts...)
}

func (c *alibabaCloudAnalyticsContainer) ElaticMapReduce(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/analytics/elatic-map-reduce.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elatic-map-reduce", opts...)
}

func (c *alibabaCloudAnalyticsContainer) OpenSearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/analytics/open-search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("open-search", opts...)
}

func (c *alibabaCloudAnalyticsContainer) AnalyticDb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/analytics/analytic-db.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("analytic-db", opts...)
}

func (c *alibabaCloudAnalyticsContainer) ClickHouse(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/analytics/click-house.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("click-house", opts...)
}
func init() {
  const namespace = "alibabacloud.analytics"
  Register(namespace, "DataLakeAnalytics", AlibabaCloudAnalytics.DataLakeAnalytics)
  Register(namespace, "ElaticMapReduce", AlibabaCloudAnalytics.ElaticMapReduce)
  Register(namespace, "OpenSearch", AlibabaCloudAnalytics.OpenSearch)
  Register(namespace, "AnalyticDb", AlibabaCloudAnalytics.AnalyticDb)
  Register(namespace, "ClickHouse", AlibabaCloudAnalytics.ClickHouse)
}
