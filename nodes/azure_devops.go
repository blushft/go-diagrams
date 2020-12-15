package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuredevopsContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureDevops = &azuredevopsContainer{path: "assets/azure/devops"}

func (c *azuredevopsContainer) Devops(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/devops.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("devops", opts...)
}

func (c *azuredevopsContainer) DevtestLabs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/devtest-labs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("devtest-labs", opts...)
}

func (c *azuredevopsContainer) Pipelines(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/pipelines.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pipelines", opts...)
}

func (c *azuredevopsContainer) Repos(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/repos.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("repos", opts...)
}

func (c *azuredevopsContainer) TestPlans(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/test-plans.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("test-plans", opts...)
}

func (c *azuredevopsContainer) ApplicationInsights(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/application-insights.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("application-insights", opts...)
}

func (c *azuredevopsContainer) Artifacts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/artifacts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("artifacts", opts...)
}

func (c *azuredevopsContainer) Boards(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/devops/boards.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("boards", opts...)
}

func init() {
  const namespace = "azure.devops"
  Register(namespace, "Devops", azureDevops.Devops)
  Register(namespace, "DevtestLabs", azureDevops.DevtestLabs)
  Register(namespace, "Pipelines", azureDevops.Pipelines)
  Register(namespace, "Repos", azureDevops.Repos)
  Register(namespace, "TestPlans", azureDevops.TestPlans)
  Register(namespace, "ApplicationInsights", azureDevops.ApplicationInsights)
  Register(namespace, "Artifacts", azureDevops.Artifacts)
  Register(namespace, "Boards", azureDevops.Boards)
}
