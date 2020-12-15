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
	opts = append(opts, attr.AssetImage("assets/aws/enduser/appstream-2-0.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("appstream-2-0", opts...)
}

func (c *enduserContainer) Workdocs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/enduser/workdocs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("workdocs", opts...)
}

func (c *enduserContainer) Worklink(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/enduser/worklink.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("worklink", opts...)
}

func (c *enduserContainer) Workspaces(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/enduser/workspaces.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("workspaces", opts...)
}

func init() {
  const namespace = "aws.enduser"
  Register(namespace, "Appstream20", Enduser.Appstream20)
  Register(namespace, "Workdocs", Enduser.Workdocs)
  Register(namespace, "Worklink", Enduser.Worklink)
  Register(namespace, "Workspaces", Enduser.Workspaces)
}
