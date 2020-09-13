package elastic

import "github.com/blushft/go-diagrams/diagram"

type observabilityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Observability = &observabilityContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/observability",
}

func (c *observabilityContainer) Apm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/observability/apm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *observabilityContainer) Logs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/observability/logs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *observabilityContainer) Metrics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/observability/metrics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *observabilityContainer) Observability(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/observability/observability.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *observabilityContainer) Uptime(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/observability/uptime.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
