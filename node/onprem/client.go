package onprem

import "github.com/blushft/go-diagrams/node"

type clientContainer struct {
	path string
	opts []node.Option
}

var Client = &clientContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/client",
}

func (c *clientContainer) Client(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/client/client.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *clientContainer) User(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/client/user.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *clientContainer) Users(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/client/users.png")}, c.opts, opts)
	return node.New(nopts...)
}
