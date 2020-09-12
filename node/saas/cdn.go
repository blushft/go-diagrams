package saas

import "github.com/blushft/go-diagrams/node"

type cdnContainer struct {
	path string
	opts []node.Option
}

var Cdn = &cdnContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/cdn",
}

func (c *cdnContainer) Cloudflare(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/cdn/cloudflare.png")}, c.opts, opts)
	return node.New(nopts...)
}
