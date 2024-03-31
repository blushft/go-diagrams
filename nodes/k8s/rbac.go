package k8s

import "github.com/blushft/go-diagrams/diagram"

type rbacContainer struct {
	path string
	opts []diagram.NodeOption
}

var Rbac = &rbacContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/rbac",
}

func (c *rbacContainer) Group(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/rbac/group.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *rbacContainer) Rb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/rbac/rb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *rbacContainer) Role(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/rbac/role.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *rbacContainer) Sa(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/rbac/sa.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *rbacContainer) User(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/rbac/user.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *rbacContainer) CRole(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/rbac/c-role.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *rbacContainer) Crb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/rbac/crb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
