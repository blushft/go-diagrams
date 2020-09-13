package generic

import "github.com/blushft/go-diagrams/diagram"

type deviceContainer struct {
	path string
	opts []diagram.NodeOption
}

var Device = &deviceContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/device",
}

func (c *deviceContainer) Mobile(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/device/mobile.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *deviceContainer) Tablet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/device/tablet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
