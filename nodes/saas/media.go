package saas

import "github.com/blushft/go-diagrams/diagram"

type mediaContainer struct {
	path string
	opts []diagram.NodeOption
}

var Media = &mediaContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/media",
}

func (c *mediaContainer) Cloudinary(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/media/cloudinary.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
