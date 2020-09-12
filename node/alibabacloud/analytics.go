package alibabacloud

import "github.com/blushft/go-diagrams/node"

type analyticsContainer struct {
	path string
	opts []node.Option
}

var Analytics = &analyticsContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/analytics",
}

func (c *analyticsContainer) AnalyticDb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/analytics/analytic-db.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) ClickHouse(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/analytics/click-house.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataLakeAnalytics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/analytics/data-lake-analytics.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) ElaticMapReduce(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/analytics/elatic-map-reduce.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) OpenSearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/analytics/open-search.png")}, c.opts, opts)
	return node.New(nopts...)
}
