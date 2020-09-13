package saas

import "github.com/blushft/go-diagrams/diagram"

type socialContainer struct {
	path string
	opts []diagram.NodeOption
}

var Social = &socialContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/social",
}

func (c *socialContainer) Facebook(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/social/facebook.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *socialContainer) Twitter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/social/twitter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
