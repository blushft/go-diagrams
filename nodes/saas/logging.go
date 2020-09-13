package saas

import "github.com/blushft/go-diagrams/diagram"

type loggingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Logging = &loggingContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/logging",
}

func (c *loggingContainer) Datadog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/logging/datadog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Papertrail(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/logging/papertrail.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
