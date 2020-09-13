package apps

import "github.com/blushft/go-diagrams/diagram"

type containerContainer struct {
	path string
	opts []diagram.NodeOption
}

var Container = &containerContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/container",
}

func (c *containerContainer) Docker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/container/docker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Rkt(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/container/rkt.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
