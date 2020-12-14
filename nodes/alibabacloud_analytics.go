package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudAnalyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabaAnalytics = &alibabaCloudAnalyticsContainer{path: "assets/alibabacloud/analytics"}

func (c *alibabaCloudAnalyticsContainer) DataLakeAnalytics(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake-analytics", attr.AssetImage("assets/alibabacloud/analytics/data-lake-analytics.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudAnalyticsContainer) ElaticMapReduce(opts ...attr.Attribute) *node.Node {
	return node.New("elatic-map-reduce", attr.AssetImage("assets/alibabacloud/analytics/elatic-map-reduce.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudAnalyticsContainer) OpenSearch(opts ...attr.Attribute) *node.Node {
	return node.New("open-search", attr.AssetImage("assets/alibabacloud/analytics/open-search.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudAnalyticsContainer) AnalyticDb(opts ...attr.Attribute) *node.Node {
	return node.New("analytic-db", attr.AssetImage("assets/alibabacloud/analytics/analytic-db.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudAnalyticsContainer) ClickHouse(opts ...attr.Attribute) *node.Node {
	return node.New("click-house", attr.AssetImage("assets/alibabacloud/analytics/click-house.png"), attr.Shape(attr.None))
}


func init() {
	const namespace = "alibabacloud.analytics"
	Register()
