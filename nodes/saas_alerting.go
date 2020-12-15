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
	opts = append(opts, attr.AssetImage("assets/saas/alerting/opsgenie.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("opsgenie", opts...)
}

func (c *saasAlertingContainer) Pushover(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/saas/alerting/pushover.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pushover", opts...)
}

func init() {
  const namespace = "saas.alerting"
  Register(namespace, "Opsgenie", SaasAlerting.Opsgenie)
  Register(namespace, "Pushover", SaasAlerting.Pushover)
}
