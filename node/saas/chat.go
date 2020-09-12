package saas

import "github.com/blushft/go-diagrams/node"

type chatContainer struct {
	path string
	opts []node.Option
}

var Chat = &chatContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/chat",
}

func (c *chatContainer) Slack(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/chat/slack.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *chatContainer) Telegram(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/chat/telegram.png")}, c.opts, opts)
	return node.New(nopts...)
}
