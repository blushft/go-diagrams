package openstack

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/storage",
}

func (c *storageContainer) Cinder(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/storage/cinder.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Manila(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/storage/manila.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Swift(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/storage/swift.png")}, c.opts, opts)
	return node.New(nopts...)
}
