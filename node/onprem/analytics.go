package onprem

import "github.com/blushft/go-diagrams/node"

type analyticsContainer struct {
	path string
	opts []node.Option
}

var Analytics = &analyticsContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/analytics",
}

func (c *analyticsContainer) Hive(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/hive.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Metabase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/metabase.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Spark(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/spark.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Beam(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/beam.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Flink(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/flink.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Hadoop(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/hadoop.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Singer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/singer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Storm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/storm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Tableau(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/tableau.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Databricks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/databricks.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Dbt(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/dbt.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Norikra(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/analytics/norikra.png")}, c.opts, opts)
	return node.New(nopts...)
}
