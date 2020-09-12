package openstack

import "github.com/blushft/go-diagrams/node"

type optimizationContainer struct {
	path string
	opts []node.Option
}

var Optimization = &optimizationContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/optimization",
}

func (c *optimizationContainer) Watcher(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/optimization/watcher.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *optimizationContainer) Congress(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/optimization/congress.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *optimizationContainer) Rally(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/optimization/rally.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *optimizationContainer) Vitrage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/operations/optimization/vitrage.png")}, c.opts, opts)
	return node.New(nopts...)
}
