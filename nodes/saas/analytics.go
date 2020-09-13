package saas

import "github.com/blushft/go-diagrams/diagram"

type analyticsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Analytics = &analyticsContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/analytics",
}

func (c *analyticsContainer) Snowflake(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/analytics/snowflake.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Stitch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/analytics/stitch.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
