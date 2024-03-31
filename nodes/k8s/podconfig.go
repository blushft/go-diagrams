package k8s

import "github.com/blushft/go-diagrams/diagram"

type podconfigContainer struct {
	path string
	opts []diagram.NodeOption
}

var Podconfig = &podconfigContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/podconfig",
}

func (c *podconfigContainer) Cm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/podconfig/cm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *podconfigContainer) Secret(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/podconfig/secret.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
