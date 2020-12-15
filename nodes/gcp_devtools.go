package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type devtoolsContainer struct {
	path  string
	attrs []attr.Attribute
}

var GcpDevTools = &devtoolsContainer{path: "assets/gcp/devtools"}

func (c *devtoolsContainer) CodeForIntellij(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/code-for-intellij.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("code-for-intellij", opts...)
}

func (c *devtoolsContainer) MavenAppEnginePlugin(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/maven-app-engine-plugin.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("maven-app-engine-plugin", opts...)
}

func (c *devtoolsContainer) GradleAppEnginePlugin(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/gradle-app-engine-plugin.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("gradle-app-engine-plugin", opts...)
}

func (c *devtoolsContainer) Sdk(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/sdk.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sdk", opts...)
}

func (c *devtoolsContainer) TestLab(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/test-lab.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("test-lab", opts...)
}

func (c *devtoolsContainer) ToolsForVisualStudio(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/tools-for-visual-studio.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tools-for-visual-studio", opts...)
}

func (c *devtoolsContainer) Code(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/code.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("code", opts...)
}

func (c *devtoolsContainer) IdePlugins(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/ide-plugins.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ide-plugins", opts...)
}

func (c *devtoolsContainer) Scheduler(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/scheduler.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("scheduler", opts...)
}

func (c *devtoolsContainer) Tasks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/tasks.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tasks", opts...)
}

func (c *devtoolsContainer) ToolsForEclipse(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/tools-for-eclipse.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tools-for-eclipse", opts...)
}

func (c *devtoolsContainer) Build(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/build.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("build", opts...)
}

func (c *devtoolsContainer) ContainerRegistry(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/container-registry.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container-registry", opts...)
}

func (c *devtoolsContainer) SourceRepositories(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/source-repositories.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("source-repositories", opts...)
}

func (c *devtoolsContainer) ToolsForPowershell(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/devtools/tools-for-powershell.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tools-for-powershell", opts...)
}

func init(){
	const namespace = "gcp.devtools"
	Register(namespace, "CodeForIntellij", GcpDevTools.CodeForIntellij)
	Register(namespace, "MavenAppEnginePlugin", GcpDevTools.MavenAppEnginePlugin)
	Register(namespace, "GradleAppEnginePlugin", GcpDevTools.GradleAppEnginePlugin)
	Register(namespace, "Sdk", GcpDevTools.Sdk)
	Register(namespace, "TestLab", GcpDevTools.TestLab)
	Register(namespace, "ToolsForVisualStudio", GcpDevTools.ToolsForVisualStudio)
	Register(namespace, "Code", GcpDevTools.Code)
	Register(namespace, "IdePlugins", GcpDevTools.IdePlugins)
	Register(namespace, "Scheduler", GcpDevTools.Scheduler)
	Register(namespace, "Tasks", GcpDevTools.Tasks)
	Register(namespace, "ToolsForEclipse", GcpDevTools.ToolsForEclipse)
	Register(namespace, "Build", GcpDevTools.Build)
	Register(namespace, "ContainerRegistry", GcpDevTools.ContainerRegistry)
	Register(namespace, "SourceRepositories", GcpDevTools.SourceRepositories)
	Register(namespace, "ToolsForPowershell", GcpDevTools.ToolsForPowershell)
}