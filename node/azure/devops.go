package azure

import "github.com/blushft/go-diagrams/node"

type devopsContainer struct {
	path string
	opts []node.Option
}

var Devops = &devopsContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/devops",
}

func (c *devopsContainer) ApplicationInsights(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/application-insights.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) Artifacts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/artifacts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) Boards(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/boards.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) Devops(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/devops.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) DevtestLabs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/devtest-labs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) Pipelines(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/pipelines.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) Repos(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/repos.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) TestPlans(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/devops/test-plans.png")}, c.opts, opts)
	return node.New(nopts...)
}
