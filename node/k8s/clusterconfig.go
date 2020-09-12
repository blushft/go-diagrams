package k8s

import "github.com/blushft/go-diagrams/node"

type clusterconfigContainer struct {
	path string
	opts []node.Option
}

var Clusterconfig = &clusterconfigContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/clusterconfig",
}

func (c *clusterconfigContainer) Quota(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/clusterconfig/quota.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *clusterconfigContainer) Hpa(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/clusterconfig/hpa.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *clusterconfigContainer) Limits(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/clusterconfig/limits.png")}, c.opts, opts)
	return node.New(nopts...)
}
