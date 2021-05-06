package firebase

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type growContainer struct {
	path  string
	attrs []attr.Attribute
}

var Grow = &growContainer{path: "assets/firebase/grow"}

func (c *growContainer) Messaging(opts ...attr.Attribute) *node.Node {
	return node.New("messaging", attr.AssetImage("assets/firebase/grow/messaging.png"), attr.Shape(attr.None))
}

func (c *growContainer) Predictions(opts ...attr.Attribute) *node.Node {
	return node.New("predictions", attr.AssetImage("assets/firebase/grow/predictions.png"), attr.Shape(attr.None))
}

func (c *growContainer) RemoteConfig(opts ...attr.Attribute) *node.Node {
	return node.New("remote-config", attr.AssetImage("assets/firebase/grow/remote-config.png"), attr.Shape(attr.None))
}

func (c *growContainer) AbTesting(opts ...attr.Attribute) *node.Node {
	return node.New("ab-testing", attr.AssetImage("assets/firebase/grow/ab-testing.png"), attr.Shape(attr.None))
}

func (c *growContainer) AppIndexing(opts ...attr.Attribute) *node.Node {
	return node.New("app-indexing", attr.AssetImage("assets/firebase/grow/app-indexing.png"), attr.Shape(attr.None))
}

func (c *growContainer) DynamicLinks(opts ...attr.Attribute) *node.Node {
	return node.New("dynamic-links", attr.AssetImage("assets/firebase/grow/dynamic-links.png"), attr.Shape(attr.None))
}

func (c *growContainer) InAppMessaging(opts ...attr.Attribute) *node.Node {
	return node.New("in-app-messaging", attr.AssetImage("assets/firebase/grow/in-app-messaging.png"), attr.Shape(attr.None))
}

func (c *growContainer) Invites(opts ...attr.Attribute) *node.Node {
	return node.New("invites", attr.AssetImage("assets/firebase/grow/invites.png"), attr.Shape(attr.None))
}
