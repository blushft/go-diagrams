package k8s

import "github.com/blushft/go-diagrams/node"

type groupContainer struct {
	path string
	opts []node.Option
}

var Group = &groupContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/group",
}

func (c *groupContainer) Ns(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/group/ns.png")}, c.opts, opts)
	return node.New(nopts...)
}
