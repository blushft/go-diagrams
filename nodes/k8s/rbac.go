package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type rbacContainer struct {
	path  string
	attrs []attr.Attribute
}

var Rbac = &rbacContainer{path: "assets/k8s/rbac"}

func (c *rbacContainer) Sa(opts ...attr.Attribute) *node.Node {
	return node.New("sa", attr.AssetImage("assets/k8s/rbac/sa.png"), attr.Shape(attr.None))
}

func (c *rbacContainer) User(opts ...attr.Attribute) *node.Node {
	return node.New("user", attr.AssetImage("assets/k8s/rbac/user.png"), attr.Shape(attr.None))
}

func (c *rbacContainer) CRole(opts ...attr.Attribute) *node.Node {
	return node.New("c-role", attr.AssetImage("assets/k8s/rbac/c-role.png"), attr.Shape(attr.None))
}

func (c *rbacContainer) Crb(opts ...attr.Attribute) *node.Node {
	return node.New("crb", attr.AssetImage("assets/k8s/rbac/crb.png"), attr.Shape(attr.None))
}

func (c *rbacContainer) Group(opts ...attr.Attribute) *node.Node {
	return node.New("group", attr.AssetImage("assets/k8s/rbac/group.png"), attr.Shape(attr.None))
}

func (c *rbacContainer) Rb(opts ...attr.Attribute) *node.Node {
	return node.New("rb", attr.AssetImage("assets/k8s/rbac/rb.png"), attr.Shape(attr.None))
}

func (c *rbacContainer) Role(opts ...attr.Attribute) *node.Node {
	return node.New("role", attr.AssetImage("assets/k8s/rbac/role.png"), attr.Shape(attr.None))
}
