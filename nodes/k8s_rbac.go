package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sRbacContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sRbac = &k8sRbacContainer{path: "assets/k8s/rbac"}

func (c *k8sRbacContainer) Role(opts ...attr.Attribute) *node.Node {
	return node.New("role", attr.AssetImage("assets/k8s/rbac/role.png"), attr.Shape(attr.None))
}

func (c *k8sRbacContainer) Sa(opts ...attr.Attribute) *node.Node {
	return node.New("sa", attr.AssetImage("assets/k8s/rbac/sa.png"), attr.Shape(attr.None))
}

func (c *k8sRbacContainer) User(opts ...attr.Attribute) *node.Node {
	return node.New("user", attr.AssetImage("assets/k8s/rbac/user.png"), attr.Shape(attr.None))
}

func (c *k8sRbacContainer) CRole(opts ...attr.Attribute) *node.Node {
	return node.New("c-role", attr.AssetImage("assets/k8s/rbac/c-role.png"), attr.Shape(attr.None))
}

func (c *k8sRbacContainer) Crb(opts ...attr.Attribute) *node.Node {
	return node.New("crb", attr.AssetImage("assets/k8s/rbac/crb.png"), attr.Shape(attr.None))
}

func (c *k8sRbacContainer) Group(opts ...attr.Attribute) *node.Node {
	return node.New("group", attr.AssetImage("assets/k8s/rbac/group.png"), attr.Shape(attr.None))
}

func (c *k8sRbacContainer) Rb(opts ...attr.Attribute) *node.Node {
	return node.New("rb", attr.AssetImage("assets/k8s/rbac/rb.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.rbac"
  Register(namespace, "Role", K8sRbac.Role)
  Register(namespace, "Sa", K8sRbac.Sa)
  Register(namespace, "User", K8sRbac.User)
  Register(namespace, "CRole", K8sRbac.CRole)
  Register(namespace, "Crb", K8sRbac.Crb)
  Register(namespace, "Group", K8sRbac.Group)
  Register(namespace, "Rb", K8sRbac.Rb)
}
