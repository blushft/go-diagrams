package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasAlertingContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasAlerting =&saasAlertingContainer{path: "assets/saas/alerting"}

func (c *saasAlertingContainer) Opsgenie(opts ...attr.Attribute) *node.Node {
	return node.New("opsgenie", attr.AssetImage("assets/saas/alerting/opsgenie.png"), attr.Shape(attr.None))
}

func (c *saasAlertingContainer) Pushover(opts ...attr.Attribute) *node.Node {
	return node.New("pushover", attr.AssetImage("assets/saas/alerting/pushover.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.alerting"
  Register(namespace, "Opsgenie", SaasAlerting.Opsgenie)
  Register(namespace, "Pushover", SaasAlerting.Pushover)
}
