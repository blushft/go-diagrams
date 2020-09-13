package saas

import "github.com/blushft/go-diagrams/diagram"

type cdnContainer struct {
	path string
	opts []diagram.NodeOption
}

var Cdn = &cdnContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/cdn",
}

func (c *cdnContainer) Cloudflare(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/cdn/cloudflare.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
