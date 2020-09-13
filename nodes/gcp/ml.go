package gcp

import "github.com/blushft/go-diagrams/diagram"

type mlContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ml = &mlContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/ml",
}

func (c *mlContainer) InferenceApi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/inference-api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) NaturalLanguageApi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/natural-language-api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) RecommendationsAi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/recommendations-ai.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) SpeechToText(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/speech-to-text.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AdvancedSolutionsLab(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/advanced-solutions-lab.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AiHub(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/ai-hub.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Automl(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/automl.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Tpu(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/tpu.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) VideoIntelligenceApi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/video-intelligence-api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AiPlatform(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/ai-platform.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AutomlTranslation(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/automl-translation.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) JobsApi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/jobs-api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) TranslationApi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/translation-api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AiPlatformDataLabelingService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/ai-platform-data-labeling-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AutomlNaturalLanguage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/automl-natural-language.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) DialogFlowEnterpriseEdition(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/dialog-flow-enterprise-edition.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) TextToSpeech(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/text-to-speech.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) VisionApi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/vision-api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AutomlTables(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/automl-tables.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AutomlVideoIntelligence(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/automl-video-intelligence.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) AutomlVision(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/ml/automl-vision.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
