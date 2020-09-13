package aws

import "github.com/blushft/go-diagrams/diagram"

type businessContainer struct {
	path string
	opts []diagram.NodeOption
}

var Business = &businessContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/business",
}

func (c *businessContainer) AlexaForBusiness(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/business/alexa-for-business.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *businessContainer) Chime(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/business/chime.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *businessContainer) Workmail(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/business/workmail.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
