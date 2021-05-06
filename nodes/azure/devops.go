package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type devopsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Devops = &devopsContainer{path: "assets/azure/devops"}

func (c *devopsContainer) Pipelines(opts ...attr.Attribute) *node.Node {
	return node.New("pipelines", attr.AssetImage("assets/azure/devops/pipelines.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) Repos(opts ...attr.Attribute) *node.Node {
	return node.New("repos", attr.AssetImage("assets/azure/devops/repos.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) TestPlans(opts ...attr.Attribute) *node.Node {
	return node.New("test-plans", attr.AssetImage("assets/azure/devops/test-plans.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) ApplicationInsights(opts ...attr.Attribute) *node.Node {
	return node.New("application-insights", attr.AssetImage("assets/azure/devops/application-insights.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) Artifacts(opts ...attr.Attribute) *node.Node {
	return node.New("artifacts", attr.AssetImage("assets/azure/devops/artifacts.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) Boards(opts ...attr.Attribute) *node.Node {
	return node.New("boards", attr.AssetImage("assets/azure/devops/boards.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) Devops(opts ...attr.Attribute) *node.Node {
	return node.New("devops", attr.AssetImage("assets/azure/devops/devops.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) DevtestLabs(opts ...attr.Attribute) *node.Node {
	return node.New("devtest-labs", attr.AssetImage("assets/azure/devops/devtest-labs.png"), attr.Shape(attr.None))
}
