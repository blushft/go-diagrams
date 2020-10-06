package saas

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mediaContainer struct {
	path  string
	attrs []attr.Attribute
}

var Media = &mediaContainer{path: "assets/saas/media"}

func (c *mediaContainer) Cloudinary(opts ...attr.Attribute) *node.Node {
	return node.New("cloudinary", attr.AssetImage("assets/saas/media/cloudinary.png"), attr.Shape(attr.None))
}
