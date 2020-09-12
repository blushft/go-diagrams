package firebase

import "github.com/blushft/go-diagrams/node"

type developContainer struct {
	path string
	opts []node.Option
}

var Develop = &developContainer{
	opts: node.OptionSet{node.Provider("firebase"), node.Shape("none")},
	path: "assets/firebase/develop",
}

func (c *developContainer) Functions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/develop/functions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *developContainer) Hosting(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/develop/hosting.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *developContainer) MlKit(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/develop/ml-kit.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *developContainer) RealtimeDatabase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/develop/realtime-database.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *developContainer) Storage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/develop/storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *developContainer) Authentication(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/develop/authentication.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *developContainer) Firestore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/develop/firestore.png")}, c.opts, opts)
	return node.New(nopts...)
}
