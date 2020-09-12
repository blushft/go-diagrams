package generic

import "github.com/blushft/go-diagrams/node"

type osContainer struct {
	path string
	opts []node.Option
}

var Os = &osContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/os",
}

func (c *osContainer) Android(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/os/android.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *osContainer) Centos(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/os/centos.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *osContainer) Ios(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/os/ios.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *osContainer) LinuxGeneral(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/os/linux-general.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *osContainer) Suse(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/os/suse.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *osContainer) Ubuntu(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/os/ubuntu.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *osContainer) Windows(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/os/windows.png")}, c.opts, opts)
	return node.New(nopts...)
}
