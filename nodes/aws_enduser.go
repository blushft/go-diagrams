package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type enduserContainer struct {
	path  string
	attrs []attr.Attribute
}

var Enduser = &enduserContainer{path: "assets/aws/enduser"}

func (c *enduserContainer) Appstream20(opts ...attr.Attribute) *node.Node {
	return node.New("appstream-2-0", attr.AssetImage("assets/aws/enduser/appstream-2-0.png"), attr.Shape(attr.None))
}

func (c *enduserContainer) Workdocs(opts ...attr.Attribute) *node.Node {
	return node.New("workdocs", attr.AssetImage("assets/aws/enduser/workdocs.png"), attr.Shape(attr.None))
}

func (c *enduserContainer) Worklink(opts ...attr.Attribute) *node.Node {
	return node.New("worklink", attr.AssetImage("assets/aws/enduser/worklink.png"), attr.Shape(attr.None))
}

func (c *enduserContainer) Workspaces(opts ...attr.Attribute) *node.Node {
	return node.New("workspaces", attr.AssetImage("assets/aws/enduser/workspaces.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.enduser"
  Register(namespace, "Appstream20", Enduser.Appstream20)
  Register(namespace, "Workdocs", Enduser.Workdocs)
  Register(namespace, "Worklink", Enduser.Worklink)
  Register(namespace, "Workspaces", Enduser.Workspaces)
}
