package saas

import "github.com/blushft/go-diagrams/diagram"

type chatContainer struct {
	path string
	opts []diagram.NodeOption
}

var Chat = &chatContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/chat",
}

func (c *chatContainer) Slack(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/chat/slack.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *chatContainer) Telegram(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/chat/telegram.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
