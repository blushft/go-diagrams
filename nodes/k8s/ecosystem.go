package k8s

import "github.com/blushft/go-diagrams/diagram"

type ecosystemContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ecosystem = &ecosystemContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/ecosystem",
}

func (c *ecosystemContainer) Helm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/ecosystem/helm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ecosystemContainer) Krew(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/ecosystem/krew.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ecosystemContainer) Kustomize(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/ecosystem/kustomize.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
