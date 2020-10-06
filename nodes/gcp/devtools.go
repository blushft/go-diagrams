package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type devtoolsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Devtools = &devtoolsContainer{path: "assets/gcp/devtools"}

func (c *devtoolsContainer) CodeForIntellij(opts ...attr.Attribute) *node.Node {
	return node.New("code-for-intellij", attr.AssetImage("assets/gcp/devtools/code-for-intellij.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) MavenAppEnginePlugin(opts ...attr.Attribute) *node.Node {
	return node.New("maven-app-engine-plugin", attr.AssetImage("assets/gcp/devtools/maven-app-engine-plugin.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) GradleAppEnginePlugin(opts ...attr.Attribute) *node.Node {
	return node.New("gradle-app-engine-plugin", attr.AssetImage("assets/gcp/devtools/gradle-app-engine-plugin.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Sdk(opts ...attr.Attribute) *node.Node {
	return node.New("sdk", attr.AssetImage("assets/gcp/devtools/sdk.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) TestLab(opts ...attr.Attribute) *node.Node {
	return node.New("test-lab", attr.AssetImage("assets/gcp/devtools/test-lab.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) ToolsForVisualStudio(opts ...attr.Attribute) *node.Node {
	return node.New("tools-for-visual-studio", attr.AssetImage("assets/gcp/devtools/tools-for-visual-studio.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Code(opts ...attr.Attribute) *node.Node {
	return node.New("code", attr.AssetImage("assets/gcp/devtools/code.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) IdePlugins(opts ...attr.Attribute) *node.Node {
	return node.New("ide-plugins", attr.AssetImage("assets/gcp/devtools/ide-plugins.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Scheduler(opts ...attr.Attribute) *node.Node {
	return node.New("scheduler", attr.AssetImage("assets/gcp/devtools/scheduler.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Tasks(opts ...attr.Attribute) *node.Node {
	return node.New("tasks", attr.AssetImage("assets/gcp/devtools/tasks.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) ToolsForEclipse(opts ...attr.Attribute) *node.Node {
	return node.New("tools-for-eclipse", attr.AssetImage("assets/gcp/devtools/tools-for-eclipse.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Build(opts ...attr.Attribute) *node.Node {
	return node.New("build", attr.AssetImage("assets/gcp/devtools/build.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) ContainerRegistry(opts ...attr.Attribute) *node.Node {
	return node.New("container-registry", attr.AssetImage("assets/gcp/devtools/container-registry.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) SourceRepositories(opts ...attr.Attribute) *node.Node {
	return node.New("source-repositories", attr.AssetImage("assets/gcp/devtools/source-repositories.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) ToolsForPowershell(opts ...attr.Attribute) *node.Node {
	return node.New("tools-for-powershell", attr.AssetImage("assets/gcp/devtools/tools-for-powershell.png"), attr.Shape(attr.None))
}
