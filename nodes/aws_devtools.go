package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type devtoolsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Devtools = &devtoolsContainer{path: "assets/aws/devtools"}

func (c *devtoolsContainer) Codecommit(opts ...attr.Attribute) *node.Node {
	return node.New("codecommit", attr.AssetImage("assets/aws/devtools/codecommit.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Codedeploy(opts ...attr.Attribute) *node.Node {
	return node.New("codedeploy", attr.AssetImage("assets/aws/devtools/codedeploy.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) DeveloperTools(opts ...attr.Attribute) *node.Node {
	return node.New("developer-tools", attr.AssetImage("assets/aws/devtools/developer-tools.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) ToolsAndSdks(opts ...attr.Attribute) *node.Node {
	return node.New("tools-and-sdks", attr.AssetImage("assets/aws/devtools/tools-and-sdks.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) XRay(opts ...attr.Attribute) *node.Node {
	return node.New("x-ray", attr.AssetImage("assets/aws/devtools/x-ray.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) CloudDevelopmentKit(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-development-kit", attr.AssetImage("assets/aws/devtools/cloud-development-kit.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Codebuild(opts ...attr.Attribute) *node.Node {
	return node.New("codebuild", attr.AssetImage("assets/aws/devtools/codebuild.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Codestar(opts ...attr.Attribute) *node.Node {
	return node.New("codestar", attr.AssetImage("assets/aws/devtools/codestar.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) CommandLineInterface(opts ...attr.Attribute) *node.Node {
	return node.New("command-line-interface", attr.AssetImage("assets/aws/devtools/command-line-interface.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Cloud9(opts ...attr.Attribute) *node.Node {
	return node.New("cloud9", attr.AssetImage("assets/aws/devtools/cloud9.png"), attr.Shape(attr.None))
}

func (c *devtoolsContainer) Codepipeline(opts ...attr.Attribute) *node.Node {
	return node.New("codepipeline", attr.AssetImage("assets/aws/devtools/codepipeline.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.devtools"
  Register(namespace, "Codecommit", Devtools.Codecommit)
  Register(namespace, "Codedeploy", Devtools.Codedeploy)
  Register(namespace, "DeveloperTools", Devtools.DeveloperTools)
  Register(namespace, "ToolsAndSdks", Devtools.ToolsAndSdks)
  Register(namespace, "XRay", Devtools.XRay)
  Register(namespace, "CloudDevelopmentKit", Devtools.CloudDevelopmentKit)
  Register(namespace, "Codebuild", Devtools.Codebuild)
  Register(namespace, "Codestar", Devtools.Codestar)
  Register(namespace, "CommandLineInterface", Devtools.CommandLineInterface)
  Register(namespace, "Cloud9", Devtools.Cloud9)
  Register(namespace, "Codepipeline", Devtools.Codepipeline)
}
