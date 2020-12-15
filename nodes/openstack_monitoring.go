package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type openstackMonitoringContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackMonitoring =&openstackMonitoringContainer{path: "assets/openstack/monitoring"}

func (c *openstackMonitoringContainer) Monasca(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/monitoring/monasca.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("monasca", opts...)
}

func (c *openstackMonitoringContainer) Telemetry(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/monitoring/telemetry.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("telemetry", opts...)
}

func init() {
  const namespace = "openstack.monitoring"
  Register(namespace, "Monasca", OpenstackMonitoring.Monasca)
  Register(namespace, "Telemetry", OpenstackMonitoring.Telemetry)
}
