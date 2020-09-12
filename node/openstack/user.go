package openstack

import "github.com/blushft/go-diagrams/node"

type userContainer struct {
	path string
	opts []node.Option
}

var User = &userContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/user",
}

func (c *userContainer) Openstackclient(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/user/openstackclient.png")}, c.opts, opts)
	return node.New(nopts...)
}
