package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type clientContainer struct {
	path  string
	attrs []attr.Attribute
}

var Client = &clientContainer{path: "assets/apps/client"}

func (c *clientContainer) Client(opts ...attr.Attribute) *node.Node {
	return node.New("client", attr.AssetImage("assets/apps/client/client.png"), attr.Shape(attr.None))
}

func (c *clientContainer) User(opts ...attr.Attribute) *node.Node {
	return node.New("user", attr.AssetImage("assets/apps/client/user.png"), attr.Shape(attr.None))
}

func (c *clientContainer) Users(opts ...attr.Attribute) *node.Node {
	return node.New("users", attr.AssetImage("assets/apps/client/users.png"), attr.Shape(attr.None))
}
