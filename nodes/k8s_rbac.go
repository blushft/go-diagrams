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
	opts = append(opts, attr.AssetImage("assets/k8s/rbac/role.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("role", opts...)
}

func (c *k8sRbacContainer) Sa(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/rbac/sa.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sa", opts...)
}

func (c *k8sRbacContainer) User(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/rbac/user.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("user", opts...)
}

func (c *k8sRbacContainer) CRole(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/rbac/c-role.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("c-role", opts...)
}

func (c *k8sRbacContainer) Crb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/rbac/crb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("crb", opts...)
}

func (c *k8sRbacContainer) Group(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/rbac/group.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("group", opts...)
}

func (c *k8sRbacContainer) Rb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/rbac/rb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rb", opts...)
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
