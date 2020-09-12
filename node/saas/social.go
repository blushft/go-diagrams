package saas

import "github.com/blushft/go-diagrams/node"

type socialContainer struct {
	path string
	opts []node.Option
}

var Social = &socialContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/social",
}

func (c *socialContainer) Facebook(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/social/facebook.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *socialContainer) Twitter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/social/twitter.png")}, c.opts, opts)
	return node.New(nopts...)
}
