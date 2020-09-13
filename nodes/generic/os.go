package generic

import "github.com/blushft/go-diagrams/diagram"

type osContainer struct {
	path string
	opts []diagram.NodeOption
}

var Os = &osContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/os",
}

func (c *osContainer) Ubuntu(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/os/ubuntu.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *osContainer) Windows(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/os/windows.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *osContainer) Android(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/os/android.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *osContainer) Centos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/os/centos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *osContainer) Ios(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/os/ios.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *osContainer) LinuxGeneral(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/os/linux-general.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *osContainer) Suse(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/os/suse.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
