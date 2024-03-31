package aws

import "github.com/blushft/go-diagrams/diagram"

type devtoolsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Devtools = &devtoolsContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/devtools",
}

func (c *devtoolsContainer) Codebuild(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/codebuild.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Codedeploy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/codedeploy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Codestar(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/codestar.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) DeveloperTools(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/developer-tools.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) ToolsAndSdks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/tools-and-sdks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) CloudDevelopmentKit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/cloud-development-kit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Cloud9(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/cloud9.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Codecommit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/codecommit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Codepipeline(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/codepipeline.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) CommandLineInterface(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/command-line-interface.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) XRay(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/devtools/x-ray.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
