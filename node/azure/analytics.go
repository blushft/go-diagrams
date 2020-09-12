package azure

import "github.com/blushft/go-diagrams/node"

type analyticsContainer struct {
	path string
	opts []node.Option
}

var Analytics = &analyticsContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/analytics",
}

func (c *analyticsContainer) AnalysisServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/analysis-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataExplorerClusters(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/data-explorer-clusters.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataLakeAnalytics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/data-lake-analytics.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Databricks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/databricks.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Hdinsightclusters(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/hdinsightclusters.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) LogAnalyticsWorkspaces(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/log-analytics-workspaces.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataFactories(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/data-factories.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataLakeStoreGen1(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/data-lake-store-gen1.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) EventHubClusters(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/event-hub-clusters.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) EventHubs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/event-hubs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) StreamAnalyticsJobs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/analytics/stream-analytics-jobs.png")}, c.opts, opts)
	return node.New(nopts...)
}
