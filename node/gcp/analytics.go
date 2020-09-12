package gcp

import "github.com/blushft/go-diagrams/node"

type analyticsContainer struct {
	path string
	opts []node.Option
}

var Analytics = &analyticsContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/analytics",
}

func (c *analyticsContainer) Datalab(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/datalab.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Pubsub(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/pubsub.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Composer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/composer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataCatalog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/data-catalog.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Dataflow(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/dataflow.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Dataproc(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/dataproc.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Genomics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/genomics.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Bigquery(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/bigquery.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataFusion(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/data-fusion.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Dataprep(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/analytics/dataprep.png")}, c.opts, opts)
	return node.New(nopts...)
}
