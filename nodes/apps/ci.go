package apps

import "github.com/blushft/go-diagrams/diagram"

type ciContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ci = &ciContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/ci",
}

func (c *ciContainer) Teamcity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/teamcity.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Travisci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/travisci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Zuulci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/zuulci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Circleci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/circleci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Concourseci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/concourseci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Droneci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/droneci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Gitlabci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/gitlabci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Jenkins(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/jenkins.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Screwdrivercd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/screwdrivercd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Bitbucket(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/ci/bitbucket.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
