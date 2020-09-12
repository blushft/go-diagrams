package aws

import "github.com/blushft/go-diagrams/node"

type businessContainer struct {
	path string
	opts []node.Option
}

var Business = &businessContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/business",
}

func (c *businessContainer) AlexaForBusiness(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/business/alexa-for-business.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *businessContainer) Chime(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/business/chime.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *businessContainer) Workmail(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/business/workmail.png")}, c.opts, opts)
	return node.New(nopts...)
}
