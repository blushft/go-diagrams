package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasChatContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasChat =&saasChatContainer{path: "assets/saas/chat"}

func (c *saasChatContainer) Slack(opts ...attr.Attribute) *node.Node {
	return node.New("slack", attr.AssetImage("assets/saas/chat/slack.png"), attr.Shape(attr.None))
}

func (c *saasChatContainer) Telegram(opts ...attr.Attribute) *node.Node {
	return node.New("telegram", attr.AssetImage("assets/saas/chat/telegram.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.chat"
  Register(namespace, "Slack", SaasChat.Slack)
  Register(namespace, "Telegram", SaasChat.Telegram)
}
