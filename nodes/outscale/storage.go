package outscale

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("outscale"), diagram.NodeShape("none")},
	path: "assets/outscale/storage",
}

func (c *storageContainer) SimpleStorageService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/storage/simple-storage-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Storage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/storage/storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
