package openstack

import "github.com/blushft/go-diagrams/diagram"

type apiproxiesContainer struct {
	path string
	opts []diagram.NodeOption
}

var Apiproxies = &apiproxiesContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/apiproxies",
}

func (c *apiproxiesContainer) Ec2Api(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/apiproxies/ec2api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
