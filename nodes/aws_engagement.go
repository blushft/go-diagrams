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
	opts = append(opts, attr.AssetImage("assets/aws/engagement/connect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("connect", opts...)
}

func (c *engagementContainer) Pinpoint(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/engagement/pinpoint.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pinpoint", opts...)
}

func (c *engagementContainer) SimpleEmailServiceSes(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/engagement/simple-email-service-ses.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("simple-email-service-ses", opts...)
}

func init() {
  const namespace = "aws.engagement"
  Register(namespace, "Connect", Engagement.Connect)
  Register(namespace, "Pinpoint", Engagement.Pinpoint)
  Register(namespace, "SimpleEmailServiceSes", Engagement.SimpleEmailServiceSes)
}
