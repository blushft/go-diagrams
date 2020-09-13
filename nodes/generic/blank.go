package generic

import "github.com/blushft/go-diagrams/diagram"

type blankContainer struct {
	path string
	opts []diagram.NodeOption
}

var Blank = &blankContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/blank",
}

func (c *blankContainer) Blank(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/blank/blank.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
