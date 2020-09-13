package aws

import "github.com/blushft/go-diagrams/diagram"

type enduserContainer struct {
	path string
	opts []diagram.NodeOption
}

var Enduser = &enduserContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/enduser",
}

func (c *enduserContainer) Appstream20(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enduser/appstream-2-0.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enduserContainer) Workdocs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enduser/workdocs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enduserContainer) Worklink(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enduser/worklink.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enduserContainer) Workspaces(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enduser/workspaces.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
