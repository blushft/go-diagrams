package openstack

import "github.com/blushft/go-diagrams/node"

type packagingContainer struct {
	path string
	opts []node.Option
}

var Packaging = &packagingContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/packaging",
}

func (c *packagingContainer) Loci(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/packaging/loci.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *packagingContainer) Puppet(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/packaging/puppet.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *packagingContainer) Rpm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/packaging/rpm.png")}, c.opts, opts)
	return node.New(nopts...)
}
