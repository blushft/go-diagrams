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
	opts = append(opts, attr.AssetImage("assets/saas/social/facebook.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("facebook", opts...)
}

func (c *saasSocialContainer) Twitter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/saas/social/twitter.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("twitter", opts...)
}

func init() {
  const namespace = "saas.social"
  Register(namespace, "Facebook", SaasSocial.Facebook)
  Register(namespace, "Twitter", SaasSocial.Twitter)
}
