package aws

import "github.com/blushft/go-diagrams/node"

type enduserContainer struct {
	path string
	opts []node.Option
}

var Enduser = &enduserContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/enduser",
}

func (c *enduserContainer) Appstream20(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enduser/appstream-2-0.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enduserContainer) Workdocs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enduser/workdocs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enduserContainer) Worklink(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enduser/worklink.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enduserContainer) Workspaces(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enduser/workspaces.png")}, c.opts, opts)
	return node.New(nopts...)
}
