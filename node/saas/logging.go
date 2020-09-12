package saas

import "github.com/blushft/go-diagrams/node"

type loggingContainer struct {
	path string
	opts []node.Option
}

var Logging = &loggingContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/logging",
}

func (c *loggingContainer) Datadog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/logging/datadog.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *loggingContainer) Papertrail(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/logging/papertrail.png")}, c.opts, opts)
	return node.New(nopts...)
}
