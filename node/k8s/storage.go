package k8s

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/storage",
}

func (c *storageContainer) Pv(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/storage/pv.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Pvc(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/storage/pvc.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Sc(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/storage/sc.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Vol(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/storage/vol.png")}, c.opts, opts)
	return node.New(nopts...)
}
