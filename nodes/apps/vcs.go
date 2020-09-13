package apps

import "github.com/blushft/go-diagrams/diagram"

type vcsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Vcs = &vcsContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/vcs",
}

func (c *vcsContainer) Git(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/vcs/git.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *vcsContainer) Github(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/vcs/github.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *vcsContainer) Gitlab(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/vcs/gitlab.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
