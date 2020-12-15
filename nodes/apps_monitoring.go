package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type monitoringContainer struct {
	path  string
	attrs []attr.Attribute
}

var Monitoring = &monitoringContainer{path: "assets/apps/monitoring"}

func (c *monitoringContainer) PrometheusOperator(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/monitoring/prometheus-operator.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("prometheus-operator", opts...)
}

func (c *monitoringContainer) Prometheus(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/monitoring/prometheus.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("prometheus", opts...)
}

func (c *monitoringContainer) Sentry(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/monitoring/sentry.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sentry", opts...)
}

func (c *monitoringContainer) Splunk(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/monitoring/splunk.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("splunk", opts...)
}

func (c *monitoringContainer) Thanos(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/monitoring/thanos.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thanos", opts...)
}

func (c *monitoringContainer) Datadog(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/monitoring/datadog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("datadog", opts...)
}

func (c *monitoringContainer) Grafana(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/monitoring/grafana.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("grafana", opts...)
}

func init() {
	const namespace = "apps.monitoring"
	Register(namespace, "PrometheusOperator", Monitoring.PrometheusOperator)
	Register(namespace, "Prometheus", Monitoring.Prometheus)
	Register(namespace, "Sentry", Monitoring.Sentry)
	Register(namespace, "Splunk", Monitoring.Splunk)
	Register(namespace, "Thanos", Monitoring.Thanos)
	Register(namespace, "Datadog", Monitoring.Datadog)
	Register(namespace, "Grafana", Monitoring.Grafana)
}