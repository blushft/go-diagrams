package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type monitoringContainer struct {
	path  string
	attrs []attr.Attribute
}

var Monitoring = &monitoringContainer{path: "assets/apps/monitoring"}

func (c *monitoringContainer) Grafana(opts ...attr.Attribute) *node.Node {
	return node.New("grafana", attr.AssetImage("assets/apps/monitoring/grafana.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) PrometheusOperator(opts ...attr.Attribute) *node.Node {
	return node.New("prometheus-operator", attr.AssetImage("assets/apps/monitoring/prometheus-operator.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Prometheus(opts ...attr.Attribute) *node.Node {
	return node.New("prometheus", attr.AssetImage("assets/apps/monitoring/prometheus.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Sentry(opts ...attr.Attribute) *node.Node {
	return node.New("sentry", attr.AssetImage("assets/apps/monitoring/sentry.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Splunk(opts ...attr.Attribute) *node.Node {
	return node.New("splunk", attr.AssetImage("assets/apps/monitoring/splunk.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Thanos(opts ...attr.Attribute) *node.Node {
	return node.New("thanos", attr.AssetImage("assets/apps/monitoring/thanos.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Datadog(opts ...attr.Attribute) *node.Node {
	return node.New("datadog", attr.AssetImage("assets/apps/monitoring/datadog.png"), attr.Shape(attr.None))
}
