package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mlContainer struct {
	path  string
	attrs []attr.Attribute
}

var GcpMl = &mlContainer{path: "assets/gcp/ml"}

func (c *mlContainer) AiPlatform(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/ai-platform.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ai-platform", opts...)
}

func (c *mlContainer) AutomlNaturalLanguage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/automl-natural-language.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("automl-natural-language", opts...)
}

func (c *mlContainer) RecommendationsAi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/recommendations-ai.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("recommendations-ai", opts...)
}

func (c *mlContainer) Tpu(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/tpu.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tpu", opts...)
}

func (c *mlContainer) TranslationApi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/translation-api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("translation-api", opts...)
}

func (c *mlContainer) VideoIntelligenceApi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/video-intelligence-api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("video-intelligence-api", opts...)
}

func (c *mlContainer) AiHub(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/ai-hub.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ai-hub", opts...)
}

func (c *mlContainer) AiPlatformDataLabelingService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/ai-platform-data-labeling-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ai-platform-data-labeling-service", opts...)
}

func (c *mlContainer) JobsApi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/jobs-api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("jobs-api", opts...)
}

func (c *mlContainer) TextToSpeech(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/text-to-speech.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("text-to-speech", opts...)
}

func (c *mlContainer) AutomlVision(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/automl-vision.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("automl-vision", opts...)
}

func (c *mlContainer) DialogFlowEnterpriseEdition(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/dialog-flow-enterprise-edition.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dialog-flow-enterprise-edition", opts...)
}

func (c *mlContainer) SpeechToText(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/speech-to-text.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("speech-to-text", opts...)
}

func (c *mlContainer) VisionApi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/vision-api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vision-api", opts...)
}

func (c *mlContainer) AutomlVideoIntelligence(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/automl-video-intelligence.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("automl-video-intelligence", opts...)
}

func (c *mlContainer) InferenceApi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/inference-api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("inference-api", opts...)
}

func (c *mlContainer) AutomlTranslation(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/automl-translation.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("automl-translation", opts...)
}

func (c *mlContainer) Automl(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/automl.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("automl", opts...)
}

func (c *mlContainer) NaturalLanguageApi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/natural-language-api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("natural-language-api", opts...)
}

func (c *mlContainer) AdvancedSolutionsLab(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/advanced-solutions-lab.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("advanced-solutions-lab", opts...)
}

func (c *mlContainer) AutomlTables(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/ml/automl-tables.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("automl-tables", opts...)
}

func init(){
	const namespace = "gcp.ml"
	Register(namespace, "AiPlatform", GcpMl.AiPlatform)
	Register(namespace, "AutomlNaturalLanguage", GcpMl.AutomlNaturalLanguage)
	Register(namespace, "RecommendationsAi", GcpMl.RecommendationsAi)
	Register(namespace, "Tpu", GcpMl.Tpu)
	Register(namespace, "TranslationApi", GcpMl.TranslationApi)
	Register(namespace, "VideoIntelligenceApi", GcpMl.VideoIntelligenceApi)
	Register(namespace, "AiHub", GcpMl.AiHub)
	Register(namespace, "AiPlatformDataLabelingService", GcpMl.AiPlatformDataLabelingService)
	Register(namespace, "JobsApi", GcpMl.JobsApi)
	Register(namespace, "TextToSpeech", GcpMl.TextToSpeech)
	Register(namespace, "AutomlVision", GcpMl.AutomlVision)
	Register(namespace, "DialogFlowEnterpriseEdition", GcpMl.DialogFlowEnterpriseEdition)
	Register(namespace, "SpeechToText", GcpMl.SpeechToText)
	Register(namespace, "VisionApi", GcpMl.VisionApi)
	Register(namespace, "AutomlVideoIntelligence", GcpMl.AutomlVideoIntelligence)
	Register(namespace, "InferenceApi", GcpMl.InferenceApi)
	Register(namespace, "AutomlTranslation", GcpMl.AutomlTranslation)
	Register(namespace, "Automl", GcpMl.Automl)
	Register(namespace, "NaturalLanguageApi", GcpMl.NaturalLanguageApi)
	Register(namespace, "AdvancedSolutionsLab", GcpMl.AdvancedSolutionsLab)
	Register(namespace, "AutomlTables", GcpMl.AutomlTables)
}