package openstack

import "github.com/blushft/go-diagrams/diagram"

type frontendContainer struct {
	path string
	opts []diagram.NodeOption
}

var Frontend = &frontendContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/frontend",
}

func (c *frontendContainer) Horizon(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/frontend/horizon.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
