package k8s

import "github.com/blushft/go-diagrams/diagram"

type clusterconfigContainer struct {
	path string
	opts []diagram.NodeOption
}

var Clusterconfig = &clusterconfigContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/clusterconfig",
}

func (c *clusterconfigContainer) Limits(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/clusterconfig/limits.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *clusterconfigContainer) Quota(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/clusterconfig/quota.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *clusterconfigContainer) Hpa(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/clusterconfig/hpa.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
