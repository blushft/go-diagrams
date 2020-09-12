package onprem

import "github.com/blushft/go-diagrams/node"

type ciContainer struct {
	path string
	opts []node.Option
}

var Ci = &ciContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/ci",
}

func (c *ciContainer) Droneci(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/droneci.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ciContainer) Gitlabci(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/gitlabci.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ciContainer) Jenkins(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/jenkins.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ciContainer) Teamcity(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/teamcity.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ciContainer) Travisci(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/travisci.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ciContainer) Zuulci(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/zuulci.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ciContainer) Circleci(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/circleci.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *ciContainer) Concourseci(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/ci/concourseci.png")}, c.opts, opts)
	return node.New(nopts...)
}
