package generic

import "github.com/blushft/go-diagrams/diagram"

type virtualizationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Virtualization = &virtualizationContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/virtualization",
}

func (c *virtualizationContainer) Virtualbox(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/virtualization/virtualbox.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *virtualizationContainer) Vmware(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/virtualization/vmware.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *virtualizationContainer) Xen(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/virtualization/xen.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
