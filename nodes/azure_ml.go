package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuremlContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureMl = &azuremlContainer{path: "assets/azure/ml"}

func (c *azuremlContainer) CognitiveServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/cognitive-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cognitive-services", opts...)
}

func (c *azuremlContainer) GenomicsAccounts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/genomics-accounts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("genomics-accounts", opts...)
}

func (c *azuremlContainer) MachineLearningServiceWorkspaces(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/machine-learning-service-workspaces.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("machine-learning-service-workspaces", opts...)
}

func (c *azuremlContainer) MachineLearningStudioWebServicePlans(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/machine-learning-studio-web-service-plans.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("machine-learning-studio-web-service-plans", opts...)
}

func (c *azuremlContainer) MachineLearningStudioWebServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/machine-learning-studio-web-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("machine-learning-studio-web-services", opts...)
}

func (c *azuremlContainer) MachineLearningStudioWorkspaces(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/machine-learning-studio-workspaces.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("machine-learning-studio-workspaces", opts...)
}

func (c *azuremlContainer) BatchAi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/batch-ai.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("batch-ai", opts...)
}

func (c *azuremlContainer) BotServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/ml/bot-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bot-services", opts...)
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
}
