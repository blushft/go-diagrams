package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasReccomendationConatiner struct {
	path  string
	attrs []attr.Attribute
}

var SaasRecommendation =&saasReccomendationConatiner{path: "assets/saas/recommendation"}

func (c *saasReccomendationConatiner) Recombee(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/saas/recommendation/recombee.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("recombee", opts...)
}

func init() {
  const namespace = "saas.recommendation"
  Register(namespace, "Recombee", SaasRecommendation.Recombee)
}
