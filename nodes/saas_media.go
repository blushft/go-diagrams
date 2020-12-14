package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasMediaContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasMedia =&saasMediaContainer{path: "assets/saas/media"}

func (c *saasMediaContainer) Cloudinary(opts ...attr.Attribute) *node.Node {
	return node.New("cloudinary", attr.AssetImage("assets/saas/media/cloudinary.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.media"
  Register(namespace, "Cloudinary", SaasMedia.Cloudinary)
}
