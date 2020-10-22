package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type vcsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Vcs = &vcsContainer{path: "assets/apps/vcs"}

func (c *vcsContainer) Github(opts ...attr.Attribute) *node.Node {
	return node.New("github", attr.AssetImage("assets/apps/vcs/github.png"), attr.Shape(attr.None))
}

func (c *vcsContainer) Gitlab(opts ...attr.Attribute) *node.Node {
	return node.New("gitlab", attr.AssetImage("assets/apps/vcs/gitlab.png"), attr.Shape(attr.None))
}

func (c *vcsContainer) Git(opts ...attr.Attribute) *node.Node {
	return node.New("git", attr.AssetImage("assets/apps/vcs/git.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.vcs"
	Register(namespace, "Github", Vcs.Github)
	Register(namespace, "Gitlab", Vcs.Gitlab)
	Register(namespace, "Git", Vcs.Git)
}
