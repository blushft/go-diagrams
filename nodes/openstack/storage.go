package openstack

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/storage",
}

func (c *storageContainer) Swift(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/storage/swift.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Cinder(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/storage/cinder.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Manila(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/storage/manila.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
