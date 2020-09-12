package gcp

import "github.com/blushft/go-diagrams/node"

type mlContainer struct {
	path string
	opts []node.Option
}

var Ml = &mlContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/ml",
}

func (c *mlContainer) AiPlatformDataLabelingService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/ai-platform-data-labeling-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AiPlatform(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/ai-platform.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Automl(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/automl.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) DialogFlowEnterpriseEdition(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/dialog-flow-enterprise-edition.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) NaturalLanguageApi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/natural-language-api.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Tpu(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/tpu.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AutomlNaturalLanguage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/automl-natural-language.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AutomlVideoIntelligence(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/automl-video-intelligence.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) SpeechToText(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/speech-to-text.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AiHub(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/ai-hub.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AutomlTranslation(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/automl-translation.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) JobsApi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/jobs-api.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) RecommendationsAi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/recommendations-ai.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) TextToSpeech(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/text-to-speech.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) TranslationApi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/translation-api.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AdvancedSolutionsLab(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/advanced-solutions-lab.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AutomlTables(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/automl-tables.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) AutomlVision(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/automl-vision.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) InferenceApi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/inference-api.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) VideoIntelligenceApi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/video-intelligence-api.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) VisionApi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/ml/vision-api.png")}, c.opts, opts)
	return node.New(nopts...)
}
