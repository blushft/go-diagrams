package apps

import "github.com/blushft/go-diagrams/diagram"

type clientContainer struct {
	path string
	opts []diagram.NodeOption
}

var Client = &clientContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/client",
}

func (c *clientContainer) Client(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/client/client.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *clientContainer) User(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/client/user.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *clientContainer) Users(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/client/users.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
