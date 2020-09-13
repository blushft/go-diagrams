package k8s

import "github.com/blushft/go-diagrams/diagram"

type groupContainer struct {
	path string
	opts []diagram.NodeOption
}

var Group = &groupContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/group",
}

func (c *groupContainer) Ns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/group/ns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
