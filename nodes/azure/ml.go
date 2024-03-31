package azure

import "github.com/blushft/go-diagrams/diagram"

type mlContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ml = &mlContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/ml",
}

func (c *mlContainer) MachineLearningServiceWorkspaces(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/machine-learning-service-workspaces.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) MachineLearningStudioWebServicePlans(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/machine-learning-studio-web-service-plans.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) MachineLearningStudioWebServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/machine-learning-studio-web-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) MachineLearningStudioWorkspaces(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/machine-learning-studio-workspaces.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) BatchAi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/batch-ai.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) BotServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/bot-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) CognitiveServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/cognitive-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) GenomicsAccounts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/ml/genomics-accounts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
