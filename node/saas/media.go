package saas

import "github.com/blushft/go-diagrams/node"

type mediaContainer struct {
	path string
	opts []node.Option
}

var Media = &mediaContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/media",
}

func (c *mediaContainer) Cloudinary(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/media/cloudinary.png")}, c.opts, opts)
	return node.New(nopts...)
}
