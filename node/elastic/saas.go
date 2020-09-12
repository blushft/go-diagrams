package elastic

import "github.com/blushft/go-diagrams/node"

type saasContainer struct {
	path string
	opts []node.Option
}

var Saas = &saasContainer{
	opts: node.OptionSet{node.Provider("elastic"), node.Shape("none")},
	path: "assets/elastic/saas",
}

func (c *saasContainer) Cloud(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/saas/cloud.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *saasContainer) Elastic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/saas/elastic.png")}, c.opts, opts)
	return node.New(nopts...)
}
