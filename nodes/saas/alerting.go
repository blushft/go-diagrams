package saas

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alertingContainer struct {
	path  string
	attrs []attr.Attribute
}

var Alerting = &alertingContainer{path: "assets/saas/alerting"}

func (c *alertingContainer) Opsgenie(opts ...attr.Attribute) *node.Node {
	return node.New("opsgenie", attr.AssetImage("assets/saas/alerting/opsgenie.png"), attr.Shape(attr.None))
}

func (c *alertingContainer) Pushover(opts ...attr.Attribute) *node.Node {
	return node.New("pushover", attr.AssetImage("assets/saas/alerting/pushover.png"), attr.Shape(attr.None))
}
