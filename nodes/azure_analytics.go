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
	opts = append(opts, attr.AssetImage("assets/azure/analytics/data-factories.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-factories", opts...)
}

func (c *azureAnalyticsController) DataLakeStoreGen1(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/data-lake-store-gen1.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-lake-store-gen1", opts...)
}

func (c *azureAnalyticsController) Databricks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/databricks.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("databricks", opts...)
}

func (c *azureAnalyticsController) EventHubClusters(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/event-hub-clusters.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("event-hub-clusters", opts...)
}

func (c *azureAnalyticsController) LogAnalyticsWorkspaces(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/log-analytics-workspaces.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("log-analytics-workspaces", opts...)
}

func (c *azureAnalyticsController) AnalysisServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/analysis-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("analysis-services", opts...)
}

func (c *azureAnalyticsController) DataExplorerClusters(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/data-explorer-clusters.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-explorer-clusters", opts...)
}

func (c *azureAnalyticsController) DataLakeAnalytics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/data-lake-analytics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-lake-analytics", opts...)
}

func (c *azureAnalyticsController) EventHubs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/event-hubs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("event-hubs", opts...)
}

func (c *azureAnalyticsController) Hdinsightclusters(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/hdinsightclusters.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hdinsightclusters", opts...)
}

func (c *azureAnalyticsController) StreamAnalyticsJobs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/analytics/stream-analytics-jobs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("stream-analytics-jobs", opts...)
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
