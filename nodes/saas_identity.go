package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type saasIdentityContainer struct {
	path  string
	attrs []attr.Attribute
}

var SaasIdentity =&saasIdentityContainer{path: "assets/saas/identity"}

func (c *saasIdentityContainer) Auth0(opts ...attr.Attribute) *node.Node {
	return node.New("auth0", attr.AssetImage("assets/saas/identity/auth0.png"), attr.Shape(attr.None))
}

func (c *saasIdentityContainer) Okta(opts ...attr.Attribute) *node.Node {
	return node.New("okta", attr.AssetImage("assets/saas/identity/okta.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "saas.identity"
  Register(namespace, "Auth0", SaasIdentity.Auth0)
  Register(namespace, "Okta", SaasIdentity.Okta)
}
