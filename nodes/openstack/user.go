package openstack

import "github.com/blushft/go-diagrams/diagram"

type userContainer struct {
	path string
	opts []diagram.NodeOption
}

var User = &userContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/user",
}

func (c *userContainer) Openstackclient(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/user/openstackclient.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
