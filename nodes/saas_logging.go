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
	opts = append(opts, attr.AssetImage("assets/saas/logging/datadog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("datadog", opts...)
}

func (c *saasLoggingContainer) Papertrail(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/saas/logging/papertrail.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("papertrail", opts...)
}

func init() {
  const namespace = "saas.logging"
  Register(namespace, "Datadog", SaasLogging.Datadog)
  Register(namespace, "Papertrail", SaasLogging.Papertrail)
}
