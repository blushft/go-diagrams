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
	return node.New("recombee", attr.AssetImage("assets/saas/recommendation/recombee.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.recommendation"
  Register(namespace, "Recombee", SaasRecommendation.Recombee)
}