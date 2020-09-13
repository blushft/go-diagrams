package gcp

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/storage",
}

func (c *storageContainer) Filestore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/storage/filestore.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) PersistentDisk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/storage/persistent-disk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Storage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/storage/storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
