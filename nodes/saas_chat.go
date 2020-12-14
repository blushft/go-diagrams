package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type chatContainer struct {
	path  string
	attrs []attr.Attribute
}

var Chat = &chatContainer{path: "assets/saas/chat"}

func (c *chatContainer) Slack(opts ...attr.Attribute) *node.Node {
	return node.New("slack", attr.AssetImage("assets/saas/chat/slack.png"), attr.Shape(attr.None))
}

func (c *chatContainer) Telegram(opts ...attr.Attribute) *node.Node {
	return node.New("telegram", attr.AssetImage("assets/saas/chat/telegram.png"), attr.Shape(attr.None))
}
