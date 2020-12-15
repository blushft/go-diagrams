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
	opts = append(opts, attr.AssetImage("assets/saas/identity/auth0.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("auth0", opts...)
}

func (c *saasIdentityContainer) Okta(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/saas/identity/okta.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("okta", opts...)
}

func init() {
  const namespace = "saas.identity"
  Register(namespace, "Auth0", SaasIdentity.Auth0)
  Register(namespace, "Okta", SaasIdentity.Okta)
}
