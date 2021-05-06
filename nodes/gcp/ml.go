package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mlContainer struct {
	path  string
	attrs []attr.Attribute
}

var Ml = &mlContainer{path: "assets/gcp/ml"}

func (c *mlContainer) JobsApi(opts ...attr.Attribute) *node.Node {
	return node.New("jobs-api", attr.AssetImage("assets/gcp/ml/jobs-api.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AiPlatformDataLabelingService(opts ...attr.Attribute) *node.Node {
	return node.New("ai-platform-data-labeling-service", attr.AssetImage("assets/gcp/ml/ai-platform-data-labeling-service.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AiPlatform(opts ...attr.Attribute) *node.Node {
	return node.New("ai-platform", attr.AssetImage("assets/gcp/ml/ai-platform.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AutomlNaturalLanguage(opts ...attr.Attribute) *node.Node {
	return node.New("automl-natural-language", attr.AssetImage("assets/gcp/ml/automl-natural-language.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AutomlTables(opts ...attr.Attribute) *node.Node {
	return node.New("automl-tables", attr.AssetImage("assets/gcp/ml/automl-tables.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AdvancedSolutionsLab(opts ...attr.Attribute) *node.Node {
	return node.New("advanced-solutions-lab", attr.AssetImage("assets/gcp/ml/advanced-solutions-lab.png"), attr.Shape(attr.None))
}

func (c *mlContainer) InferenceApi(opts ...attr.Attribute) *node.Node {
	return node.New("inference-api", attr.AssetImage("assets/gcp/ml/inference-api.png"), attr.Shape(attr.None))
}

func (c *mlContainer) NaturalLanguageApi(opts ...attr.Attribute) *node.Node {
	return node.New("natural-language-api", attr.AssetImage("assets/gcp/ml/natural-language-api.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Automl(opts ...attr.Attribute) *node.Node {
	return node.New("automl", attr.AssetImage("assets/gcp/ml/automl.png"), attr.Shape(attr.None))
}

func (c *mlContainer) SpeechToText(opts ...attr.Attribute) *node.Node {
	return node.New("speech-to-text", attr.AssetImage("assets/gcp/ml/speech-to-text.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Tpu(opts ...attr.Attribute) *node.Node {
	return node.New("tpu", attr.AssetImage("assets/gcp/ml/tpu.png"), attr.Shape(attr.None))
}

func (c *mlContainer) VisionApi(opts ...attr.Attribute) *node.Node {
	return node.New("vision-api", attr.AssetImage("assets/gcp/ml/vision-api.png"), attr.Shape(attr.None))
}

func (c *mlContainer) DialogFlowEnterpriseEdition(opts ...attr.Attribute) *node.Node {
	return node.New("dialog-flow-enterprise-edition", attr.AssetImage("assets/gcp/ml/dialog-flow-enterprise-edition.png"), attr.Shape(attr.None))
}

func (c *mlContainer) RecommendationsAi(opts ...attr.Attribute) *node.Node {
	return node.New("recommendations-ai", attr.AssetImage("assets/gcp/ml/recommendations-ai.png"), attr.Shape(attr.None))
}

func (c *mlContainer) TextToSpeech(opts ...attr.Attribute) *node.Node {
	return node.New("text-to-speech", attr.AssetImage("assets/gcp/ml/text-to-speech.png"), attr.Shape(attr.None))
}

func (c *mlContainer) TranslationApi(opts ...attr.Attribute) *node.Node {
	return node.New("translation-api", attr.AssetImage("assets/gcp/ml/translation-api.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AiHub(opts ...attr.Attribute) *node.Node {
	return node.New("ai-hub", attr.AssetImage("assets/gcp/ml/ai-hub.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AutomlTranslation(opts ...attr.Attribute) *node.Node {
	return node.New("automl-translation", attr.AssetImage("assets/gcp/ml/automl-translation.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AutomlVideoIntelligence(opts ...attr.Attribute) *node.Node {
	return node.New("automl-video-intelligence", attr.AssetImage("assets/gcp/ml/automl-video-intelligence.png"), attr.Shape(attr.None))
}

func (c *mlContainer) AutomlVision(opts ...attr.Attribute) *node.Node {
	return node.New("automl-vision", attr.AssetImage("assets/gcp/ml/automl-vision.png"), attr.Shape(attr.None))
}

func (c *mlContainer) VideoIntelligenceApi(opts ...attr.Attribute) *node.Node {
	return node.New("video-intelligence-api", attr.AssetImage("assets/gcp/ml/video-intelligence-api.png"), attr.Shape(attr.None))
}
