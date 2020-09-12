package saas

import "github.com/blushft/go-diagrams/node"

type analyticsContainer struct {
	path string
	opts []node.Option
}

var Analytics = &analyticsContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/analytics",
}

func (c *analyticsContainer) Stitch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/analytics/stitch.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Snowflake(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/analytics/snowflake.png")}, c.opts, opts)
	return node.New(nopts...)
}
