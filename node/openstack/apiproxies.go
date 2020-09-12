package openstack

import "github.com/blushft/go-diagrams/node"

type apiproxiesContainer struct {
	path string
	opts []node.Option
}

var Apiproxies = &apiproxiesContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/apiproxies",
}

func (c *apiproxiesContainer) Ec2Api(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/apiproxies/ec2api.png")}, c.opts, opts)
	return node.New(nopts...)
}
