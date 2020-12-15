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
	opts = append(opts, attr.AssetImage("assets/aws/general/marketplace.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("marketplace", opts...)
}

func (c *generalContainer) TradicionalServer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/tradicional-server.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tradicional-server", opts...)
}

func (c *generalContainer) GenericDatabase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/generic-database.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("generic-database", opts...)
}

func (c *generalContainer) GenericFirewall(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/generic-firewall.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("generic-firewall", opts...)
}

func (c *generalContainer) GenericOfficeBuilding(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/generic-office-building.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("generic-office-building", opts...)
}

func (c *generalContainer) GenericSamlToken(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/generic-saml-token.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("generic-saml-token", opts...)
}

func (c *generalContainer) GenericSdk(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/generic-sdk.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("generic-sdk", opts...)
}

func (c *generalContainer) User(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/user.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("user", opts...)
}

func (c *generalContainer) Disk(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/disk.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("disk", opts...)
}

func (c *generalContainer) General(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/general.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("general", opts...)
}

func (c *generalContainer) Users(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/general/users.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("users", opts...)
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
