package aws

import "github.com/blushft/go-diagrams/node"

type engagementContainer struct {
	path string
	opts []node.Option
}

var Engagement = &engagementContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/engagement",
}

func (c *engagementContainer) Pinpoint(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/engagement/pinpoint.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *engagementContainer) SimpleEmailServiceSes(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/engagement/simple-email-service-ses.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *engagementContainer) Connect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/engagement/connect.png")}, c.opts, opts)
	return node.New(nopts...)
}
