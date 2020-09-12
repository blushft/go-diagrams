package generic

import "github.com/blushft/go-diagrams/node"

type blankContainer struct {
	path string
	opts []node.Option
}

var Blank = &blankContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/blank",
}

func (c *blankContainer) Blank(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/blank/blank.png")}, c.opts, opts)
	return node.New(nopts...)
}
