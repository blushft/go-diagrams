package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type growContainer struct {
	path  string
	attrs []attr.Attribute
}

var Grow = &growContainer{path: "assets/firebase/grow"}

func (c *growContainer) AppIndexing(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/app-indexing.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-indexing", opts...)
}

func (c *growContainer) DynamicLinks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/dynamic-links.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dynamic-links", opts...)
}

func (c *growContainer) InAppMessaging(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/in-app-messaging.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("in-app-messaging", opts...)
}

func (c *growContainer) Invites(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/invites.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("invites", opts...)
}

func (c *growContainer) Messaging(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/messaging.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("messaging", opts...)
}

func (c *growContainer) Predictions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/predictions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("predictions", opts...)
}

func (c *growContainer) RemoteConfig(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/remote-config.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("remote-config", opts...)
}

func (c *growContainer) AbTesting(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/grow/ab-testing.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ab-testing", opts...)
}
