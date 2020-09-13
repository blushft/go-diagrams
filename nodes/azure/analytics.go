package azure

import "github.com/blushft/go-diagrams/diagram"

type analyticsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Analytics = &analyticsContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/analytics",
}

func (c *analyticsContainer) AnalysisServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/analysis-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) DataLakeAnalytics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/data-lake-analytics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Hdinsightclusters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/hdinsightclusters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) LogAnalyticsWorkspaces(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/log-analytics-workspaces.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) StreamAnalyticsJobs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/stream-analytics-jobs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) DataExplorerClusters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/data-explorer-clusters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) DataFactories(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/data-factories.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) DataLakeStoreGen1(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/data-lake-store-gen1.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Databricks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/databricks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) EventHubClusters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/event-hub-clusters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) EventHubs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/analytics/event-hubs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
