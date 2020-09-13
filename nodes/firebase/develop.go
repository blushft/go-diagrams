package firebase

import "github.com/blushft/go-diagrams/diagram"

type developContainer struct {
	path string
	opts []diagram.NodeOption
}

var Develop = &developContainer{
	opts: diagram.OptionSet{diagram.Provider("firebase"), diagram.NodeShape("none")},
	path: "assets/firebase/develop",
}

func (c *developContainer) Authentication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/develop/authentication.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *developContainer) Firestore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/develop/firestore.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *developContainer) Functions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/develop/functions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *developContainer) Hosting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/develop/hosting.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *developContainer) MlKit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/develop/ml-kit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *developContainer) RealtimeDatabase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/develop/realtime-database.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *developContainer) Storage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/develop/storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
