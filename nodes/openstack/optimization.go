package openstack

import "github.com/blushft/go-diagrams/diagram"

type optimizationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Optimization = &optimizationContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/optimization",
}

func (c *optimizationContainer) Congress(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/optimization/congress.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *optimizationContainer) Rally(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/optimization/rally.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *optimizationContainer) Vitrage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/optimization/vitrage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *optimizationContainer) Watcher(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/operations/optimization/watcher.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
