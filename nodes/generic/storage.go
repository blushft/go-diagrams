package generic

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/storage",
}

func (c *storageContainer) Storage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/storage/storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
