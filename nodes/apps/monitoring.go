package apps

import "github.com/blushft/go-diagrams/diagram"

type monitoringContainer struct {
	path string
	opts []diagram.NodeOption
}

var Monitoring = &monitoringContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/monitoring",
}

func (c *monitoringContainer) Datadog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/monitoring/datadog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Grafana(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/monitoring/grafana.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) PrometheusOperator(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/monitoring/prometheus-operator.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Prometheus(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/monitoring/prometheus.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Sentry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/monitoring/sentry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Splunk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/monitoring/splunk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Thanos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/monitoring/thanos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
