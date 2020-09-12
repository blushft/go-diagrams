package k8s

import "github.com/blushft/go-diagrams/node"

type rbacContainer struct {
	path string
	opts []node.Option
}

var Rbac = &rbacContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/rbac",
}

func (c *rbacContainer) Group(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/rbac/group.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *rbacContainer) Rb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/rbac/rb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *rbacContainer) Role(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/rbac/role.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *rbacContainer) Sa(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/rbac/sa.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *rbacContainer) User(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/rbac/user.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *rbacContainer) CRole(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/rbac/c-role.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *rbacContainer) Crb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/rbac/crb.png")}, c.opts, opts)
	return node.New(nopts...)
}
