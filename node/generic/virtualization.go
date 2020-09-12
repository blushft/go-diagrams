package generic

import "github.com/blushft/go-diagrams/node"

type virtualizationContainer struct {
	path string
	opts []node.Option
}

var Virtualization = &virtualizationContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/virtualization",
}

func (c *virtualizationContainer) Virtualbox(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/virtualization/virtualbox.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *virtualizationContainer) Vmware(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/virtualization/vmware.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *virtualizationContainer) Xen(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/virtualization/xen.png")}, c.opts, opts)
	return node.New(nopts...)
}
