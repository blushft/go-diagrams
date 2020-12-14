package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type generalContainer struct {
	path  string
	attrs []attr.Attribute
}

var General = &generalContainer{path: "assets/aws/general"}

func (c *generalContainer) Marketplace(opts ...attr.Attribute) *node.Node {
	return node.New("marketplace", attr.AssetImage("assets/aws/general/marketplace.png"), attr.Shape(attr.None))
}

func (c *generalContainer) TradicionalServer(opts ...attr.Attribute) *node.Node {
	return node.New("tradicional-server", attr.AssetImage("assets/aws/general/tradicional-server.png"), attr.Shape(attr.None))
}

func (c *generalContainer) GenericDatabase(opts ...attr.Attribute) *node.Node {
	return node.New("generic-database", attr.AssetImage("assets/aws/general/generic-database.png"), attr.Shape(attr.None))
}

func (c *generalContainer) GenericFirewall(opts ...attr.Attribute) *node.Node {
	return node.New("generic-firewall", attr.AssetImage("assets/aws/general/generic-firewall.png"), attr.Shape(attr.None))
}

func (c *generalContainer) GenericOfficeBuilding(opts ...attr.Attribute) *node.Node {
	return node.New("generic-office-building", attr.AssetImage("assets/aws/general/generic-office-building.png"), attr.Shape(attr.None))
}

func (c *generalContainer) GenericSamlToken(opts ...attr.Attribute) *node.Node {
	return node.New("generic-saml-token", attr.AssetImage("assets/aws/general/generic-saml-token.png"), attr.Shape(attr.None))
}

func (c *generalContainer) GenericSdk(opts ...attr.Attribute) *node.Node {
	return node.New("generic-sdk", attr.AssetImage("assets/aws/general/generic-sdk.png"), attr.Shape(attr.None))
}

func (c *generalContainer) User(opts ...attr.Attribute) *node.Node {
	return node.New("user", attr.AssetImage("assets/aws/general/user.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Disk(opts ...attr.Attribute) *node.Node {
	return node.New("disk", attr.AssetImage("assets/aws/general/disk.png"), attr.Shape(attr.None))
}

func (c *generalContainer) General(opts ...attr.Attribute) *node.Node {
	return node.New("general", attr.AssetImage("assets/aws/general/general.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Users(opts ...attr.Attribute) *node.Node {
	return node.New("users", attr.AssetImage("assets/aws/general/users.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.general"
  Register(namespace, "Marketplace", General.Marketplace)
  Register(namespace, "TradicionalServer", General.TradicionalServer)
  Register(namespace, "GenericDatabase", General.GenericDatabase)
  Register(namespace, "GenericFirewall", General.GenericFirewall)
  Register(namespace, "GenericOfficeBuilding", General.GenericOfficeBuilding)
  Register(namespace, "GenericSamlToken", General.GenericSamlToken)
  Register(namespace, "GenericSdk", General.GenericSdk)
  Register(namespace, "User", General.User)
  Register(namespace, "Disk", General.Disk)
  Register(namespace, "General", General.General)
  Register(namespace, "Users", General.Users)
}
