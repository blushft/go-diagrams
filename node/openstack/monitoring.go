package openstack

import "github.com/blushft/go-diagrams/node"

type monitoringContainer struct {
	path string
	opts []node.Option
}

var Monitoring = &monitoringContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/monitoring",
}

func (c *monitoringContainer) Monasca(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/monitoring/monasca.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Telemetry(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/monitoring/telemetry.png")}, c.opts, opts)
	return node.New(nopts...)
}
