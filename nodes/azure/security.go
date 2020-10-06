package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type securityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &securityContainer{path: "assets/azure/security"}

func (c *securityContainer) KeyVaults(opts ...attr.Attribute) *node.Node {
	return node.New("key-vaults", attr.AssetImage("assets/azure/security/key-vaults.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SecurityCenter(opts ...attr.Attribute) *node.Node {
	return node.New("security-center", attr.AssetImage("assets/azure/security/security-center.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Sentinel(opts ...attr.Attribute) *node.Node {
	return node.New("sentinel", attr.AssetImage("assets/azure/security/sentinel.png"), attr.Shape(attr.None))
}
