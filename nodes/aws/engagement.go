package aws

import "github.com/blushft/go-diagrams/diagram"

type engagementContainer struct {
	path string
	opts []diagram.NodeOption
}

var Engagement = &engagementContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/engagement",
}

func (c *engagementContainer) Pinpoint(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/engagement/pinpoint.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *engagementContainer) SimpleEmailServiceSes(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/engagement/simple-email-service-ses.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *engagementContainer) Connect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/engagement/connect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
