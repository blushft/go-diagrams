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
	opts = append(opts, attr.AssetImage("assets/saas/cdn/cloudflare.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudflare", opts...)
}

func init() {
  const namespace = "saas.cdn"
  Register(namespace, "Cloudflare", SaasCdn.Cloudflare)
}
