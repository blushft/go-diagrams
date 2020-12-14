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
	return node.New("monasca", attr.AssetImage("assets/openstack/monitoring/monasca.png"), attr.Shape(attr.None))
}

func (c *openstackMonitoringContainer) Telemetry(opts ...attr.Attribute) *node.Node {
	return node.New("telemetry", attr.AssetImage("assets/openstack/monitoring/telemetry.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.monitoring"
  Register(namespace, "Monasca", OpenstackMonitoring.Monasca)
  Register(namespace, "Telemetry", OpenstackMonitoring.Telemetry)
}
