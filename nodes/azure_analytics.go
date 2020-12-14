package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azureAnalyticsController struct {
	path  string
	attrs []attr.Attribute
}

var azureAnalytics = &azureAnalyticsController{path: "assets/azure/analytics"}

func (c *azureAnalyticsController) DataFactories(opts ...attr.Attribute) *node.Node {
	return node.New("data-factories", attr.AssetImage("assets/azure/analytics/data-factories.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) DataLakeStoreGen1(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake-store-gen1", attr.AssetImage("assets/azure/analytics/data-lake-store-gen1.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) Databricks(opts ...attr.Attribute) *node.Node {
	return node.New("databricks", attr.AssetImage("assets/azure/analytics/databricks.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) EventHubClusters(opts ...attr.Attribute) *node.Node {
	return node.New("event-hub-clusters", attr.AssetImage("assets/azure/analytics/event-hub-clusters.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) LogAnalyticsWorkspaces(opts ...attr.Attribute) *node.Node {
	return node.New("log-analytics-workspaces", attr.AssetImage("assets/azure/analytics/log-analytics-workspaces.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) AnalysisServices(opts ...attr.Attribute) *node.Node {
	return node.New("analysis-services", attr.AssetImage("assets/azure/analytics/analysis-services.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) DataExplorerClusters(opts ...attr.Attribute) *node.Node {
	return node.New("data-explorer-clusters", attr.AssetImage("assets/azure/analytics/data-explorer-clusters.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) DataLakeAnalytics(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake-analytics", attr.AssetImage("assets/azure/analytics/data-lake-analytics.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) EventHubs(opts ...attr.Attribute) *node.Node {
	return node.New("event-hubs", attr.AssetImage("assets/azure/analytics/event-hubs.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) Hdinsightclusters(opts ...attr.Attribute) *node.Node {
	return node.New("hdinsightclusters", attr.AssetImage("assets/azure/analytics/hdinsightclusters.png"), attr.Shape(attr.None))
}

func (c *azureAnalyticsController) StreamAnalyticsJobs(opts ...attr.Attribute) *node.Node {
	return node.New("stream-analytics-jobs", attr.AssetImage("assets/azure/analytics/stream-analytics-jobs.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.analytics"
  Register(namespace, "DataFactories", azureAnalytics.DataFactories)
  Register(namespace, "DataLakeStoreGen1", azureAnalytics.DataLakeStoreGen1)
  Register(namespace, "Databricks", azureAnalytics.Databricks)
  Register(namespace, "EventHubClusters", azureAnalytics.EventHubClusters)
  Register(namespace, "LogAnalyticsWorkspaces", azureAnalytics.LogAnalyticsWorkspaces)
  Register(namespace, "AnalysisServices", azureAnalytics.AnalysisServices)
  Register(namespace, "DataExplorerClusters", azureAnalytics.DataExplorerClusters)
  Register(namespace, "DataLakeAnalytics", azureAnalytics.DataLakeAnalytics)
  Register(namespace, "EventHubs", azureAnalytics.EventHubs)
  Register(namespace, "Hdinsightclusters", azureAnalytics.Hdinsightclusters)
  Register(namespace, "StreamAnalyticsJobs", azureAnalytics.StreamAnalyticsJobs)
}
