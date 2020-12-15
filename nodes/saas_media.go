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
	opts = append(opts, attr.AssetImage("assets/saas/media/cloudinary.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudinary", opts...)
}

func init() {
  const namespace = "saas.media"
  Register(namespace, "Cloudinary", SaasMedia.Cloudinary)
}
