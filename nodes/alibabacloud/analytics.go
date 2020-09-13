package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type analyticsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Analytics = &analyticsContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/analytics",
}

func (c *analyticsContainer) AnalyticDb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/analytics/analytic-db.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) ClickHouse(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/analytics/click-house.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) DataLakeAnalytics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/analytics/data-lake-analytics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) ElaticMapReduce(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/analytics/elatic-map-reduce.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) OpenSearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/analytics/open-search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
