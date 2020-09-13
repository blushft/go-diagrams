package openstack

import "github.com/blushft/go-diagrams/diagram"

type containerservicesContainer struct {
	path string
	opts []diagram.NodeOption
}

var Containerservices = &containerservicesContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/containerservices",
}

func (c *containerservicesContainer) Kuryr(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/adjacentenablers/containerservices/kuryr.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
