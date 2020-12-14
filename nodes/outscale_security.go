package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type outscaleSecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var OutscaleSecurity =&outscaleSecurityContainer{path: "assets/outscale/security"}

func (c *outscaleSecurityContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/outscale/security/firewall.png"), attr.Shape(attr.None))
}

func (c *outscaleSecurityContainer) IdentityAndAccessManagement(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management", attr.AssetImage("assets/outscale/security/identity-and-access-management.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "outscale.security"
  Register(namespace, "Firewall", OutscaleSecurity.Firewall)
  Register(namespace, "IdentityAndAccessManagement", OutscaleSecurity.IdentityAndAccessManagement)
}
