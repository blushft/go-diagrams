package aws

import "github.com/blushft/go-diagrams/diagram"

type generalContainer struct {
	path string
	opts []diagram.NodeOption
}

var General = &generalContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/general",
}

func (c *generalContainer) Marketplace(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/marketplace.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Users(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/users.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) General(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/general.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericFirewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericOfficeBuilding(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-office-building.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericSdk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-sdk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) TradicionalServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/tradicional-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) User(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/user.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Disk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/disk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericDatabase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-database.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericSamlToken(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-saml-token.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
