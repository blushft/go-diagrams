package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasSocialContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasSocial =&saasSocialContainer{path: "assets/saas/social"}

func (c *saasSocialContainer) Facebook(opts ...attr.Attribute) *node.Node {
	return node.New("facebook", attr.AssetImage("assets/saas/social/facebook.png"), attr.Shape(attr.None))
}

func (c *saasSocialContainer) Twitter(opts ...attr.Attribute) *node.Node {
	return node.New("twitter", attr.AssetImage("assets/saas/social/twitter.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.social"
  Register(namespace, "Facebook", SaasSocial.Facebook)
  Register(namespace, "Twitter", SaasSocial.Twitter)
}
