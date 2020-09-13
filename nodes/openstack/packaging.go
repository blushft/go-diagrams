package openstack

import "github.com/blushft/go-diagrams/diagram"

type packagingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Packaging = &packagingContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/packaging",
}

func (c *packagingContainer) Puppet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/packaging/puppet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *packagingContainer) Rpm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/packaging/rpm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *packagingContainer) Loci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/packaging/loci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
