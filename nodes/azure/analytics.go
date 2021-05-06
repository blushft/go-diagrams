package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type analyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Analytics = &analyticsContainer{path: "assets/azure/analytics"}

func (c *analyticsContainer) DataExplorerClusters(opts ...attr.Attribute) *node.Node {
	return node.New("data-explorer-clusters", attr.AssetImage("assets/azure/analytics/data-explorer-clusters.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) DataLakeStoreGen1(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake-store-gen1", attr.AssetImage("assets/azure/analytics/data-lake-store-gen1.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) LogAnalyticsWorkspaces(opts ...attr.Attribute) *node.Node {
	return node.New("log-analytics-workspaces", attr.AssetImage("assets/azure/analytics/log-analytics-workspaces.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) StreamAnalyticsJobs(opts ...attr.Attribute) *node.Node {
	return node.New("stream-analytics-jobs", attr.AssetImage("assets/azure/analytics/stream-analytics-jobs.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) AnalysisServices(opts ...attr.Attribute) *node.Node {
	return node.New("analysis-services", attr.AssetImage("assets/azure/analytics/analysis-services.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) DataFactories(opts ...attr.Attribute) *node.Node {
	return node.New("data-factories", attr.AssetImage("assets/azure/analytics/data-factories.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) DataLakeAnalytics(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake-analytics", attr.AssetImage("assets/azure/analytics/data-lake-analytics.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Databricks(opts ...attr.Attribute) *node.Node {
	return node.New("databricks", attr.AssetImage("assets/azure/analytics/databricks.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) EventHubClusters(opts ...attr.Attribute) *node.Node {
	return node.New("event-hub-clusters", attr.AssetImage("assets/azure/analytics/event-hub-clusters.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) EventHubs(opts ...attr.Attribute) *node.Node {
	return node.New("event-hubs", attr.AssetImage("assets/azure/analytics/event-hubs.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Hdinsightclusters(opts ...attr.Attribute) *node.Node {
	return node.New("hdinsightclusters", attr.AssetImage("assets/azure/analytics/hdinsightclusters.png"), attr.Shape(attr.None))
}
