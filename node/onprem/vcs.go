package onprem

import "github.com/blushft/go-diagrams/node"

type vcsContainer struct {
	path string
	opts []node.Option
}

var Vcs = &vcsContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/vcs",
}

func (c *vcsContainer) Github(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/vcs/github.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *vcsContainer) Gitlab(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/vcs/gitlab.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *vcsContainer) Git(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/vcs/git.png")}, c.opts, opts)
	return node.New(nopts...)
}
