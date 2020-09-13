package k8s

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/network",
}

func (c *networkContainer) Ep(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/network/ep.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Ing(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/network/ing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Netpol(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/network/netpol.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Svc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/network/svc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
