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
	opts = append(opts, attr.AssetImage("assets/apps/vcs/github.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("github", opts...)
}

func (c *vcsContainer) Gitlab(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/vcs/gitlab.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("gitlab", opts...)
}

func (c *vcsContainer) Git(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/vcs/git.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("git", opts...)
}

func init() {
	const namespace = "apps.vcs"
	Register(namespace, "Github", Vcs.Github)
	Register(namespace, "Gitlab", Vcs.Gitlab)
	Register(namespace, "Git", Vcs.Git)
}
