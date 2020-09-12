package outscale

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("outscale"), node.Shape("none")},
	path: "assets/outscale/storage",
}

func (c *storageContainer) SimpleStorageService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/storage/simple-storage-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Storage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/storage/storage.png")}, c.opts, opts)
	return node.New(nopts...)
}
