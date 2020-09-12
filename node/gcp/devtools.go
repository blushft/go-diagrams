package gcp

import "github.com/blushft/go-diagrams/node"

type devtoolsContainer struct {
	path string
	opts []node.Option
}

var Devtools = &devtoolsContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/devtools",
}

func (c *devtoolsContainer) CodeForIntellij(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/code-for-intellij.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) ContainerRegistry(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/container-registry.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) TestLab(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/test-lab.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) MavenAppEnginePlugin(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/maven-app-engine-plugin.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Scheduler(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/scheduler.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Tasks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/tasks.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) ToolsForVisualStudio(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/tools-for-visual-studio.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Code(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/code.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) GradleAppEnginePlugin(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/gradle-app-engine-plugin.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) IdePlugins(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/ide-plugins.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Sdk(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/sdk.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) SourceRepositories(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/source-repositories.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) ToolsForEclipse(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/tools-for-eclipse.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) Build(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/build.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devtoolsContainer) ToolsForPowershell(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/devtools/tools-for-powershell.png")}, c.opts, opts)
	return node.New(nopts...)
}
