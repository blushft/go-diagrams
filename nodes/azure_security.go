package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuresecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureSecurity = &azuresecurityContainer{path: "assets/azure/security"}

func (c *azuresecurityContainer) KeyVaults(opts ...attr.Attribute) *node.Node {
	return node.New("key-vaults", attr.AssetImage("assets/azure/security/key-vaults.png"), attr.Shape(attr.None))
}

func (c *azuresecurityContainer) SecurityCenter(opts ...attr.Attribute) *node.Node {
	return node.New("security-center", attr.AssetImage("assets/azure/security/security-center.png"), attr.Shape(attr.None))
}

func (c *azuresecurityContainer) Sentinel(opts ...attr.Attribute) *node.Node {
	return node.New("sentinel", attr.AssetImage("assets/azure/security/sentinel.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.security"
  Register(namespace, "KeyVaults", azureSecurity.KeyVaults)
  Register(namespace, "SecurityCenter", azureSecurity.SecurityCenter)
  Register(namespace, "Sentinel", azureSecurity.Sentinel)
}
