package k8s

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/storage",
}

func (c *storageContainer) Vol(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/storage/vol.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Pv(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/storage/pv.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Pvc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/storage/pvc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Sc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/storage/sc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
