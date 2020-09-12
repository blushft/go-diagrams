package generic

import "github.com/blushft/go-diagrams/node"

type deviceContainer struct {
	path string
	opts []node.Option
}

var Device = &deviceContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/device",
}

func (c *deviceContainer) Mobile(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/device/mobile.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *deviceContainer) Tablet(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/device/tablet.png")}, c.opts, opts)
	return node.New(nopts...)
}
