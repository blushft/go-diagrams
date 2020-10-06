package saas

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type loggingContainer struct {
	path  string
	attrs []attr.Attribute
}

var Logging = &loggingContainer{path: "assets/saas/logging"}

func (c *loggingContainer) Datadog(opts ...attr.Attribute) *node.Node {
	return node.New("datadog", attr.AssetImage("assets/saas/logging/datadog.png"), attr.Shape(attr.None))
}

func (c *loggingContainer) Papertrail(opts ...attr.Attribute) *node.Node {
	return node.New("papertrail", attr.AssetImage("assets/saas/logging/papertrail.png"), attr.Shape(attr.None))
}
