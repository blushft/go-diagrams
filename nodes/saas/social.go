package saas

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type socialContainer struct {
	path  string
	attrs []attr.Attribute
}

var Social = &socialContainer{path: "assets/saas/social"}

func (c *socialContainer) Facebook(opts ...attr.Attribute) *node.Node {
	return node.New("facebook", attr.AssetImage("assets/saas/social/facebook.png"), attr.Shape(attr.None))
}

func (c *socialContainer) Twitter(opts ...attr.Attribute) *node.Node {
	return node.New("twitter", attr.AssetImage("assets/saas/social/twitter.png"), attr.Shape(attr.None))
}
