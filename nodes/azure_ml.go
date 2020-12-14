package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuremlContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureMl = &awsMlContainer{path: "assets/azure/ml"}

func (c *azuremlContainer) CognitiveServices(opts ...attr.Attribute) *node.Node {
	return node.New("cognitive-services", attr.AssetImage("assets/azure/ml/cognitive-services.png"), attr.Shape(attr.None))
}

func (c *azuremlContainer) GenomicsAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("genomics-accounts", attr.AssetImage("assets/azure/ml/genomics-accounts.png"), attr.Shape(attr.None))
}

func (c *azuremlContainer) MachineLearningServiceWorkspaces(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-service-workspaces", attr.AssetImage("assets/azure/ml/machine-learning-service-workspaces.png"), attr.Shape(attr.None))
}

func (c *azuremlContainer) MachineLearningStudioWebServicePlans(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-studio-web-service-plans", attr.AssetImage("assets/azure/ml/machine-learning-studio-web-service-plans.png"), attr.Shape(attr.None))
}

func (c *azuremlContainer) MachineLearningStudioWebServices(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-studio-web-services", attr.AssetImage("assets/azure/ml/machine-learning-studio-web-services.png"), attr.Shape(attr.None))
}

func (c *azuremlContainer) MachineLearningStudioWorkspaces(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning-studio-workspaces", attr.AssetImage("assets/azure/ml/machine-learning-studio-workspaces.png"), attr.Shape(attr.None))
}

func (c *azuremlContainer) BatchAi(opts ...attr.Attribute) *node.Node {
	return node.New("batch-ai", attr.AssetImage("assets/azure/ml/batch-ai.png"), attr.Shape(attr.None))
}

func (c *azuremlContainer) BotServices(opts ...attr.Attribute) *node.Node {
	return node.New("bot-services", attr.AssetImage("assets/azure/ml/bot-services.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.ml"
  Register(namespace, "CognitiveServices", azureMl.CognitiveServices)
  Register(namespace, "GenomicsAccounts", azureMl.GenomicsAccounts)
  Register(namespace, "MachineLearningServiceWorkspaces", azureMl.MachineLearningServiceWorkspaces)
  Register(namespace, "MachineLearningStudioWebServicePlans", azureMl.MachineLearningStudioWebServicePlans)
  Register(namespace, "MachineLearningStudioWebServices", azureMl.MachineLearningStudioWebServices)
  Register(namespace, "MachineLearningStudioWorkspaces", azureMl.MachineLearningStudioWorkspaces)
  Register(namespace, "BatchAi", azureMl.BatchAi)
  Register(namespace, "BotServices", azureMl.BotServices)
  Register(namespace, "init", azureMl.init)
}
