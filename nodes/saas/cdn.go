package saas

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type cdnContainer struct {
	path  string
	attrs []attr.Attribute
}

var Cdn = &cdnContainer{path: "assets/saas/cdn"}

func (c *cdnContainer) Cloudflare(opts ...attr.Attribute) *node.Node {
	return node.New("cloudflare", attr.AssetImage("assets/saas/cdn/cloudflare.png"), attr.Shape(attr.None))
}
