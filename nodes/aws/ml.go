package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mlContainer struct {
	path  string
	attrs []attr.Attribute
}

var Ml = &mlContainer{path: "assets/aws/ml"}

func (c *mlContainer) Deepracer(opts ...attr.Attribute) *node.Node {
	return node.New("deepracer", attr.AssetImage("assets/aws/ml/deepracer.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Forecast(opts ...attr.Attribute) *node.Node {
	return node.New("forecast", attr.AssetImage("assets/aws/ml/forecast.png"), attr.Shape(attr.None))
}

func (c *mlContainer) SagemakerNotebook(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-notebook", attr.AssetImage("assets/aws/ml/sagemaker-notebook.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Sagemaker(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker", attr.AssetImage("assets/aws/ml/sagemaker.png"), attr.Shape(attr.None))
}

func (c *mlContainer) TensorflowOnAws(opts ...attr.Attribute) *node.Node {
	return node.New("tensorflow-on-aws", attr.AssetImage("assets/aws/ml/tensorflow-on-aws.png"), attr.Shape(attr.None))
}

func (c *mlContainer) DeepLearningAmis(opts ...attr.Attribute) *node.Node {
	return node.New("deep-learning-amis", attr.AssetImage("assets/aws/ml/deep-learning-amis.png"), attr.Shape(attr.None))
}

func (c *mlContainer) MachineLearning(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning", attr.AssetImage("assets/aws/ml/machine-learning.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Personalize(opts ...attr.Attribute) *node.Node {
	return node.New("personalize", attr.AssetImage("assets/aws/ml/personalize.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Rekognition(opts ...attr.Attribute) *node.Node {
	return node.New("rekognition", attr.AssetImage("assets/aws/ml/rekognition.png"), attr.Shape(attr.None))
}

func (c *mlContainer) SagemakerTrainingJob(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-training-job", attr.AssetImage("assets/aws/ml/sagemaker-training-job.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Translate(opts ...attr.Attribute) *node.Node {
	return node.New("translate", attr.AssetImage("assets/aws/ml/translate.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Lex(opts ...attr.Attribute) *node.Node {
	return node.New("lex", attr.AssetImage("assets/aws/ml/lex.png"), attr.Shape(attr.None))
}

func (c *mlContainer) DeepLearningContainers(opts ...attr.Attribute) *node.Node {
	return node.New("deep-learning-containers", attr.AssetImage("assets/aws/ml/deep-learning-containers.png"), attr.Shape(attr.None))
}

func (c *mlContainer) ElasticInference(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-inference", attr.AssetImage("assets/aws/ml/elastic-inference.png"), attr.Shape(attr.None))
}

func (c *mlContainer) SagemakerModel(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-model", attr.AssetImage("assets/aws/ml/sagemaker-model.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Textract(opts ...attr.Attribute) *node.Node {
	return node.New("textract", attr.AssetImage("assets/aws/ml/textract.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Transcribe(opts ...attr.Attribute) *node.Node {
	return node.New("transcribe", attr.AssetImage("assets/aws/ml/transcribe.png"), attr.Shape(attr.None))
}

func (c *mlContainer) ApacheMxnetOnAws(opts ...attr.Attribute) *node.Node {
	return node.New("apache-mxnet-on-aws", attr.AssetImage("assets/aws/ml/apache-mxnet-on-aws.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Deeplens(opts ...attr.Attribute) *node.Node {
	return node.New("deeplens", attr.AssetImage("assets/aws/ml/deeplens.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Polly(opts ...attr.Attribute) *node.Node {
	return node.New("polly", attr.AssetImage("assets/aws/ml/polly.png"), attr.Shape(attr.None))
}

func (c *mlContainer) SagemakerGroundTruth(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-ground-truth", attr.AssetImage("assets/aws/ml/sagemaker-ground-truth.png"), attr.Shape(attr.None))
}

func (c *mlContainer) Comprehend(opts ...attr.Attribute) *node.Node {
	return node.New("comprehend", attr.AssetImage("assets/aws/ml/comprehend.png"), attr.Shape(attr.None))
}
