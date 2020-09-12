package saas

import "github.com/blushft/go-diagrams/node"

type alertingContainer struct {
	path string
	opts []node.Option
}

var Alerting = &alertingContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/alerting",
}

func (c *alertingContainer) Opsgenie(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/alerting/opsgenie.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *alertingContainer) Pushover(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/alerting/pushover.png")}, c.opts, opts)
	return node.New(nopts...)
}
