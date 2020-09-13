package firebase

import "github.com/blushft/go-diagrams/diagram"

type extentionsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Extentions = &extentionsContainer{
	opts: diagram.OptionSet{diagram.Provider("firebase"), diagram.NodeShape("none")},
	path: "assets/firebase/extentions",
}

func (c *extentionsContainer) Extensions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/extentions/extensions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
