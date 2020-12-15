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
	opts = append(opts, attr.AssetImage("assets/saas/chat/slack.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("slack", opts...)
}

func (c *saasChatContainer) Telegram(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/saas/chat/telegram.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("telegram", opts...)
}

func init() {
  const namespace = "saas.chat"
  Register(namespace, "Slack", SaasChat.Slack)
  Register(namespace, "Telegram", SaasChat.Telegram)
}
