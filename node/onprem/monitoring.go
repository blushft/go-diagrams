package onprem

import "github.com/blushft/go-diagrams/node"

type monitoringContainer struct {
	path string
	opts []node.Option
}

var Monitoring = &monitoringContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/monitoring",
}

func (c *monitoringContainer) Splunk(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/monitoring/splunk.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Thanos(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/monitoring/thanos.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Datadog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/monitoring/datadog.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Grafana(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/monitoring/grafana.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) PrometheusOperator(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/monitoring/prometheus-operator.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Prometheus(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/monitoring/prometheus.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Sentry(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/monitoring/sentry.png")}, c.opts, opts)
	return node.New(nopts...)
}
