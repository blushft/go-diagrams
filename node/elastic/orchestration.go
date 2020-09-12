package elastic

import "github.com/blushft/go-diagrams/node"

type orchestrationContainer struct {
	path string
	opts []node.Option
}

var Orchestration = &orchestrationContainer{
	opts: node.OptionSet{node.Provider("elastic"), node.Shape("none")},
	path: "assets/elastic/orchestration",
}

func (c *orchestrationContainer) Ece(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/orchestration/ece.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *orchestrationContainer) Eck(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/orchestration/eck.png")}, c.opts, opts)
	return node.New(nopts...)
}
