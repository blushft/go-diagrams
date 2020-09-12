package generic

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/storage",
}

func (c *storageContainer) Storage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/storage/storage.png")}, c.opts, opts)
	return node.New(nopts...)
}
