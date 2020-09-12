package gcp

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/storage",
}

func (c *storageContainer) Storage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/storage/storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Filestore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/storage/filestore.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) PersistentDisk(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/storage/persistent-disk.png")}, c.opts, opts)
	return node.New(nopts...)
}
