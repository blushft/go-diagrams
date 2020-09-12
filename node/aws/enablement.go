package aws

import "github.com/blushft/go-diagrams/node"

type enablementContainer struct {
	path string
	opts []node.Option
}

var Enablement = &enablementContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/enablement",
}

func (c *enablementContainer) Iq(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enablement/iq.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enablementContainer) ManagedServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enablement/managed-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enablementContainer) ProfessionalServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enablement/professional-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enablementContainer) Support(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/enablement/support.png")}, c.opts, opts)
	return node.New(nopts...)
}
