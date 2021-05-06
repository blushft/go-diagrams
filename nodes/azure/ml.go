package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mlContainer struct {
	path  string
	attrs []attr.Attribute
}

var Ml = &mlContainer{path: "assets/azure/ml"}

func (c *mlContainer) BotServices(opts ...attr.Attribute) *node.Node {
	return node.New("bot-services", attr.AssetImage("assets/azure/ml/bot-services.png"), attr.Shape(attr.None))
}

func (c *mlContainer) CognitiveServices(opts ...attr.Attribute) *node.Node {
	return node.New("cognitive-services", attr.AssetImage("assets/azure/ml/cognitive-services.png"), attr.Shape(attr.None))
}

func (c *mlContainer) GenomicsAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("genomics-accounts", attr.AssetImage("assets/azure/ml/genomics-accounts.png"), attr.Shape(attr.None))
}

func (c *mlContainer) MachineLearningServiceWorkspaces(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-service-workspaces", attr.AssetImage("assets/azure/ml/machine-learning-service-workspaces.png"), attr.Shape(attr.None))
}

func (c *mlContainer) MachineLearningStudioWebServicePlans(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-studio-web-service-plans", attr.AssetImage("assets/azure/ml/machine-learning-studio-web-service-plans.png"), attr.Shape(attr.None))
}

func (c *mlContainer) MachineLearningStudioWebServices(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-studio-web-services", attr.AssetImage("assets/azure/ml/machine-learning-studio-web-services.png"), attr.Shape(attr.None))
}

func (c *mlContainer) MachineLearningStudioWorkspaces(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-studio-workspaces", attr.AssetImage("assets/azure/ml/machine-learning-studio-workspaces.png"), attr.Shape(attr.None))
}

func (c *mlContainer) BatchAi(opts ...attr.Attribute) *node.Node {
	return node.New("batch-ai", attr.AssetImage("assets/azure/ml/batch-ai.png"), attr.Shape(attr.None))
}
