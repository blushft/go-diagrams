package apps

import "github.com/blushft/go-diagrams/diagram"

type analyticsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Analytics = &analyticsContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/analytics",
}

func (c *analyticsContainer) Dbt(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/dbt.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Hadoop(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/hadoop.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Hive(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/hive.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Metabase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/metabase.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Tableau(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/tableau.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Spark(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/spark.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Storm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/storm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Beam(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/beam.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Databricks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/databricks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Flink(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/flink.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Norikra(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/norikra.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Singer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/analytics/singer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
