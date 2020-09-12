package aws

import "github.com/blushft/go-diagrams/node"

type devtoolsContainer struct {
	path string
	opts []node.Option
}

var Devtools = &devtoolsContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/devtools",
}

func (c *devtoolsContainer) Codebuild(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/codebuild.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Codecommit(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/codecommit.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Codedeploy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/codedeploy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Codestar(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/codestar.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) XRay(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/x-ray.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) CloudDevelopmentKit(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/cloud-development-kit.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Cloud9(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/cloud9.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) DeveloperTools(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/developer-tools.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) ToolsAndSdks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/tools-and-sdks.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Codepipeline(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/codepipeline.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) CommandLineInterface(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/devtools/command-line-interface.png")}, c.opts, opts)
	return node.New(nopts...)
}
