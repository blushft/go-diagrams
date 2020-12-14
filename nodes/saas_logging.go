package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasLoggingContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasLogging =&saasLoggingContainer{path: "assets/saas/logging"}

func (c *saasLoggingContainer) Datadog(opts ...attr.Attribute) *node.Node {
	return node.New("datadog", attr.AssetImage("assets/saas/logging/datadog.png"), attr.Shape(attr.None))
}

func (c *saasLoggingContainer) Papertrail(opts ...attr.Attribute) *node.Node {
	return node.New("papertrail", attr.AssetImage("assets/saas/logging/papertrail.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.logging"
  Register(namespace, "Datadog", SaasLogging.Datadog)
  Register(namespace, "Papertrail", SaasLogging.Papertrail)
}
