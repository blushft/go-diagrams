package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasCdnContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasCdn =&saasCdnContainer{path: "assets/saas/cdn"}

func (c *saasCdnContainer) Cloudflare(opts ...attr.Attribute) *node.Node {
	return node.New("cloudflare", attr.AssetImage("assets/saas/cdn/cloudflare.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.cdn"
  Register(namespace, "Cloudflare", SaasCdn.Cloudflare)
}
