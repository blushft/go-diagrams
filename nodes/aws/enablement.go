package aws

import "github.com/blushft/go-diagrams/diagram"

type enablementContainer struct {
	path string
	opts []diagram.NodeOption
}

var Enablement = &enablementContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/enablement",
}

func (c *enablementContainer) Iq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enablement/iq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enablementContainer) ManagedServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enablement/managed-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enablementContainer) ProfessionalServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enablement/professional-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enablementContainer) Support(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/enablement/support.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
