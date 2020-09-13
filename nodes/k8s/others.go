package k8s

import "github.com/blushft/go-diagrams/diagram"

type othersContainer struct {
	path string
	opts []diagram.NodeOption
}

var Others = &othersContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/others",
}

func (c *othersContainer) Crd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/others/crd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *othersContainer) Psp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/others/psp.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
