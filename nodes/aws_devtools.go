package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsDevToolsContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSDevTools = &awsDevToolsContainer{path: "assets/aws/devtools"}

func (c *awsDevToolsContainer) Codecommit(opts ...attr.Attribute) *node.Node {
	return node.New("codecommit", attr.AssetImage("assets/aws/devtools/codecommit.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) Codedeploy(opts ...attr.Attribute) *node.Node {
	return node.New("codedeploy", attr.AssetImage("assets/aws/devtools/codedeploy.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) DeveloperTools(opts ...attr.Attribute) *node.Node {
	return node.New("developer-tools", attr.AssetImage("assets/aws/devtools/developer-tools.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) ToolsAndSdks(opts ...attr.Attribute) *node.Node {
	return node.New("tools-and-sdks", attr.AssetImage("assets/aws/devtools/tools-and-sdks.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) XRay(opts ...attr.Attribute) *node.Node {
	return node.New("x-ray", attr.AssetImage("assets/aws/devtools/x-ray.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) CloudDevelopmentKit(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-development-kit", attr.AssetImage("assets/aws/devtools/cloud-development-kit.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) Codebuild(opts ...attr.Attribute) *node.Node {
	return node.New("codebuild", attr.AssetImage("assets/aws/devtools/codebuild.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) Codestar(opts ...attr.Attribute) *node.Node {
	return node.New("codestar", attr.AssetImage("assets/aws/devtools/codestar.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) CommandLineInterface(opts ...attr.Attribute) *node.Node {
	return node.New("command-line-interface", attr.AssetImage("assets/aws/devtools/command-line-interface.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) Cloud9(opts ...attr.Attribute) *node.Node {
	return node.New("cloud9", attr.AssetImage("assets/aws/devtools/cloud9.png"), attr.Shape(attr.None))
}

func (c *awsDevToolsContainer) Codepipeline(opts ...attr.Attribute) *node.Node {
	return node.New("codepipeline", attr.AssetImage("assets/aws/devtools/codepipeline.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.devtools"
  Register(namespace, "Codecommit", AWSDevTools.Codecommit)
  Register(namespace, "Codedeploy", AWSDevTools.Codedeploy)
  Register(namespace, "DeveloperTools", AWSDevTools.DeveloperTools)
  Register(namespace, "ToolsAndSdks", AWSDevTools.ToolsAndSdks)
  Register(namespace, "XRay", AWSDevTools.XRay)
  Register(namespace, "CloudDevelopmentKit", AWSDevTools.CloudDevelopmentKit)
  Register(namespace, "Codebuild", AWSDevTools.Codebuild)
  Register(namespace, "Codestar", AWSDevTools.Codestar)
  Register(namespace, "CommandLineInterface", AWSDevTools.CommandLineInterface)
  Register(namespace, "Cloud9", AWSDevTools.Cloud9)
  Register(namespace, "Codepipeline", AWSDevTools.Codepipeline)
}
