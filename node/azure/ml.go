package azure

import "github.com/blushft/go-diagrams/node"

type mlContainer struct {
	path string
	opts []node.Option
}

var Ml = &mlContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/ml",
}

func (c *mlContainer) MachineLearningStudioWebServicePlans(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/machine-learning-studio-web-service-plans.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) MachineLearningStudioWebServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/machine-learning-studio-web-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) MachineLearningStudioWorkspaces(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/machine-learning-studio-workspaces.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) BatchAi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/batch-ai.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) BotServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/bot-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) CognitiveServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/cognitive-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) GenomicsAccounts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/genomics-accounts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) MachineLearningServiceWorkspaces(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/ml/machine-learning-service-workspaces.png")}, c.opts, opts)
	return node.New(nopts...)
}
