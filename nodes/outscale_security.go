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
	opts = append(opts, attr.AssetImage("assets/outscale/security/firewall.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firewall", opts...)
}

func (c *outscaleSecurityContainer) IdentityAndAccessManagement(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/security/identity-and-access-management.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("identity-and-access-management", opts...)
}

func init() {
  const namespace = "outscale.security"
  Register(namespace, "Firewall", OutscaleSecurity.Firewall)
  Register(namespace, "IdentityAndAccessManagement", OutscaleSecurity.IdentityAndAccessManagement)
}
