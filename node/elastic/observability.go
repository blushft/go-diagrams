package elastic

import "github.com/blushft/go-diagrams/node"

type observabilityContainer struct {
	path string
	opts []node.Option
}

var Observability = &observabilityContainer{
	opts: node.OptionSet{node.Provider("elastic"), node.Shape("none")},
	path: "assets/elastic/observability",
}

func (c *observabilityContainer) Logs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/observability/logs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *observabilityContainer) Metrics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/observability/metrics.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *observabilityContainer) Observability(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/observability/observability.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *observabilityContainer) Uptime(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/observability/uptime.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *observabilityContainer) Apm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/observability/apm.png")}, c.opts, opts)
	return node.New(nopts...)
}
