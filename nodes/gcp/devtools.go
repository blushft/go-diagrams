package gcp

import "github.com/blushft/go-diagrams/diagram"

type devtoolsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Devtools = &devtoolsContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/devtools",
}

func (c *devtoolsContainer) ContainerRegistry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/container-registry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) IdePlugins(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/ide-plugins.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) SourceRepositories(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/source-repositories.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Code(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/code.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) GradleAppEnginePlugin(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/gradle-app-engine-plugin.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) TestLab(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/test-lab.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) ToolsForVisualStudio(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/tools-for-visual-studio.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) CodeForIntellij(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/code-for-intellij.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Scheduler(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/scheduler.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Sdk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/sdk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) ToolsForPowershell(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/tools-for-powershell.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Build(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/build.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) MavenAppEnginePlugin(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/maven-app-engine-plugin.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) Tasks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/tasks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devtoolsContainer) ToolsForEclipse(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/devtools/tools-for-eclipse.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
