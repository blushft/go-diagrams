package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsClientcontainer struct {
	path  string
	attrs []attr.Attribute
}

var Client = &appsClientcontainer{path: "assets/apps/client"}

func (c *appsClientcontainer) Client(opts ...attr.Attribute) *node.Node {
	return node.New("client", attr.AssetImage("assets/apps/client/client.png"), attr.Shape(attr.None))
}

func (c *appsClientcontainer) User(opts ...attr.Attribute) *node.Node {
	return node.New("user", attr.AssetImage("assets/apps/client/user.png"), attr.Shape(attr.None))
}

func (c *appsClientcontainer) Users(opts ...attr.Attribute) *node.Node {
	return node.New("users", attr.AssetImage("assets/apps/client/users.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.client"
	Register(namespace, "Client", Client.Client)
	Register(namespace, "User", Client.User)
	Register(namespace, "Users", Client.Users)
}
