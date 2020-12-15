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
	opts = append(opts, attr.AssetImage("assets/apps/client/client.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("client", opts...)
}

func (c *appsClientcontainer) User(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/client/user.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("user", opts...)
}

func (c *appsClientcontainer) Users(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/client/users.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("users", opts...)
}

func init() {
	const namespace = "apps.client"
	Register(namespace, "Client", Client.Client)
	Register(namespace, "User", Client.User)
	Register(namespace, "Users", Client.Users)
}
