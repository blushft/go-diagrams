package outscale

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type securityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &securityContainer{path: "assets/outscale/security"}

func (c *securityContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/outscale/security/firewall.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdentityAndAccessManagement(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management", attr.AssetImage("assets/outscale/security/identity-and-access-management.png"), attr.Shape(attr.None))
}
