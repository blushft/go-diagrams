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
	opts = append(opts, attr.AssetImage("assets/aws/devtools/codecommit.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("codecommit", opts...)
}

func (c *awsDevToolsContainer) Codedeploy(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/codedeploy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("codedeploy", opts...)
}

func (c *awsDevToolsContainer) DeveloperTools(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/developer-tools.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("developer-tools", opts...)
}

func (c *awsDevToolsContainer) ToolsAndSdks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/tools-and-sdks.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tools-and-sdks", opts...)
}

func (c *awsDevToolsContainer) XRay(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/x-ray.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("x-ray", opts...)
}

func (c *awsDevToolsContainer) CloudDevelopmentKit(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/cloud-development-kit.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-development-kit", opts...)
}

func (c *awsDevToolsContainer) Codebuild(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/codebuild.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("codebuild", opts...)
}

func (c *awsDevToolsContainer) Codestar(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/codestar.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("codestar", opts...)
}

func (c *awsDevToolsContainer) CommandLineInterface(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/command-line-interface.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("command-line-interface", opts...)
}

func (c *awsDevToolsContainer) Cloud9(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/cloud9.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud9", opts...)
}

func (c *awsDevToolsContainer) Codepipeline(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/devtools/codepipeline.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("codepipeline", opts...)
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
