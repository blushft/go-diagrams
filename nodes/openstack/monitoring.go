package openstack

import "github.com/blushft/go-diagrams/diagram"

type monitoringContainer struct {
	path string
	opts []diagram.NodeOption
}

var Monitoring = &monitoringContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/monitoring",
}

func (c *monitoringContainer) Monasca(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/monitoring/monasca.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Telemetry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/monitoring/telemetry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
