package openstack

import "github.com/blushft/go-diagrams/node"

type frontendContainer struct {
	path string
	opts []node.Option
}

var Frontend = &frontendContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/frontend",
}

func (c *frontendContainer) Horizon(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/frontend/horizon.png")}, c.opts, opts)
	return node.New(nopts...)
}
