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
	return node.New("devops", attr.AssetImage("assets/azure/devops/devops.png"), attr.Shape(attr.None))
}

func (c *azuredevopsContainer) DevtestLabs(opts ...attr.Attribute) *node.Node {
	return node.New("devtest-labs", attr.AssetImage("assets/azure/devops/devtest-labs.png"), attr.Shape(attr.None))
}

func (c *azuredevopsContainer) Pipelines(opts ...attr.Attribute) *node.Node {
	return node.New("pipelines", attr.AssetImage("assets/azure/devops/pipelines.png"), attr.Shape(attr.None))
}

func (c *azuredevopsContainer) Repos(opts ...attr.Attribute) *node.Node {
	return node.New("repos", attr.AssetImage("assets/azure/devops/repos.png"), attr.Shape(attr.None))
}

func (c *azuredevopsContainer) TestPlans(opts ...attr.Attribute) *node.Node {
	return node.New("test-plans", attr.AssetImage("assets/azure/devops/test-plans.png"), attr.Shape(attr.None))
}

func (c *azuredevopsContainer) ApplicationInsights(opts ...attr.Attribute) *node.Node {
	return node.New("application-insights", attr.AssetImage("assets/azure/devops/application-insights.png"), attr.Shape(attr.None))
}

func (c *azuredevopsContainer) Artifacts(opts ...attr.Attribute) *node.Node {
	return node.New("artifacts", attr.AssetImage("assets/azure/devops/artifacts.png"), attr.Shape(attr.None))
}

func (c *azuredevopsContainer) Boards(opts ...attr.Attribute) *node.Node {
	return node.New("boards", attr.AssetImage("assets/azure/devops/boards.png"), attr.Shape(attr.None))
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
