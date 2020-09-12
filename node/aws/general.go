package aws

import "github.com/blushft/go-diagrams/node"

type generalContainer struct {
	path string
	opts []node.Option
}

var General = &generalContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/general",
}

func (c *generalContainer) GenericSamlToken(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/generic-saml-token.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Marketplace(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/marketplace.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) TradicionalServer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/tradicional-server.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Users(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/users.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) General(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/general.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) GenericDatabase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/generic-database.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) GenericFirewall(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/generic-firewall.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) GenericOfficeBuilding(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/generic-office-building.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) GenericSdk(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/generic-sdk.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) User(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/user.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Disk(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/general/disk.png")}, c.opts, opts)
	return node.New(nopts...)
}
