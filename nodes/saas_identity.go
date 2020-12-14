package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type identityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Identity = &identityContainer{path: "assets/saas/identity"}

func (c *identityContainer) Auth0(opts ...attr.Attribute) *node.Node {
	return node.New("auth0", attr.AssetImage("assets/saas/identity/auth0.png"), attr.Shape(attr.None))
}

func (c *identityContainer) Okta(opts ...attr.Attribute) *node.Node {
	return node.New("okta", attr.AssetImage("assets/saas/identity/okta.png"), attr.Shape(attr.None))
}
