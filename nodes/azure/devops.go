package azure

import "github.com/blushft/go-diagrams/diagram"

type devopsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Devops = &devopsContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/devops",
}

func (c *devopsContainer) Artifacts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/artifacts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) Boards(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/boards.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) Devops(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/devops.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) DevtestLabs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/devtest-labs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) Pipelines(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/pipelines.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) Repos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/repos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) TestPlans(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/test-plans.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ApplicationInsights(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/devops/application-insights.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
