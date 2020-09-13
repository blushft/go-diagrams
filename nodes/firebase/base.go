package firebase

import "github.com/blushft/go-diagrams/diagram"

type baseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Base = &baseContainer{
	opts: diagram.OptionSet{diagram.Provider("firebase"), diagram.NodeShape("none")},
	path: "assets/firebase/base",
}

func (c *baseContainer) Firebase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/base/firebase.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
