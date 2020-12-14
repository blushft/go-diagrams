package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type engagementContainer struct {
	path  string
	attrs []attr.Attribute
}

var Engagement = &engagementContainer{path: "assets/aws/engagement"}

func (c *engagementContainer) Connect(opts ...attr.Attribute) *node.Node {
	return node.New("connect", attr.AssetImage("assets/aws/engagement/connect.png"), attr.Shape(attr.None))
}

func (c *engagementContainer) Pinpoint(opts ...attr.Attribute) *node.Node {
	return node.New("pinpoint", attr.AssetImage("assets/aws/engagement/pinpoint.png"), attr.Shape(attr.None))
}

func (c *engagementContainer) SimpleEmailServiceSes(opts ...attr.Attribute) *node.Node {
	return node.New("simple-email-service-ses", attr.AssetImage("assets/aws/engagement/simple-email-service-ses.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.engagement"
  Register(namespace, "Connect", Engagement.Connect)
  Register(namespace, "Pinpoint", Engagement.Pinpoint)
  Register(namespace, "SimpleEmailServiceSes", Engagement.SimpleEmailServiceSes)
}
