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
	opts = append(opts, attr.AssetImage("assets/azure/security/key-vaults.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("key-vaults", opts...)
}

func (c *azuresecurityContainer) SecurityCenter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/security/security-center.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-center", opts...)
}

func (c *azuresecurityContainer) Sentinel(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/security/sentinel.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sentinel", opts...)
}

func init() {
  const namespace = "azure.security"
  Register(namespace, "KeyVaults", azureSecurity.KeyVaults)
  Register(namespace, "SecurityCenter", azureSecurity.SecurityCenter)
  Register(namespace, "Sentinel", azureSecurity.Sentinel)
}
