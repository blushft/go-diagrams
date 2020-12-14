package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsMlContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSMl = &awsMlContainer{path: "assets/aws/ml"}

func (c *awsMlContainer) Deepracer(opts ...attr.Attribute) *node.Node {
	return node.New("deepracer", attr.AssetImage("assets/aws/ml/deepracer.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Forecast(opts ...attr.Attribute) *node.Node {
	return node.New("forecast", attr.AssetImage("assets/aws/ml/forecast.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) SagemakerNotebook(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-notebook", attr.AssetImage("assets/aws/ml/sagemaker-notebook.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Sagemaker(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker", attr.AssetImage("assets/aws/ml/sagemaker.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) TensorflowOnAws(opts ...attr.Attribute) *node.Node {
	return node.New("tensorflow-on-aws", attr.AssetImage("assets/aws/ml/tensorflow-on-aws.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) DeepLearningAmis(opts ...attr.Attribute) *node.Node {
	return node.New("deep-learning-amis", attr.AssetImage("assets/aws/ml/deep-learning-amis.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) MachineLearning(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning", attr.AssetImage("assets/aws/ml/machine-learning.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Personalize(opts ...attr.Attribute) *node.Node {
	return node.New("personalize", attr.AssetImage("assets/aws/ml/personalize.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Rekognition(opts ...attr.Attribute) *node.Node {
	return node.New("rekognition", attr.AssetImage("assets/aws/ml/rekognition.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) SagemakerTrainingJob(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-training-job", attr.AssetImage("assets/aws/ml/sagemaker-training-job.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Translate(opts ...attr.Attribute) *node.Node {
	return node.New("translate", attr.AssetImage("assets/aws/ml/translate.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Lex(opts ...attr.Attribute) *node.Node {
	return node.New("lex", attr.AssetImage("assets/aws/ml/lex.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) DeepLearningContainers(opts ...attr.Attribute) *node.Node {
	return node.New("deep-learning-containers", attr.AssetImage("assets/aws/ml/deep-learning-containers.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) ElasticInference(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-inference", attr.AssetImage("assets/aws/ml/elastic-inference.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) SagemakerModel(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-model", attr.AssetImage("assets/aws/ml/sagemaker-model.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Textract(opts ...attr.Attribute) *node.Node {
	return node.New("textract", attr.AssetImage("assets/aws/ml/textract.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Transcribe(opts ...attr.Attribute) *node.Node {
	return node.New("transcribe", attr.AssetImage("assets/aws/ml/transcribe.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) ApacheMxnetOnAws(opts ...attr.Attribute) *node.Node {
	return node.New("apache-mxnet-on-aws", attr.AssetImage("assets/aws/ml/apache-mxnet-on-aws.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Deeplens(opts ...attr.Attribute) *node.Node {
	return node.New("deeplens", attr.AssetImage("assets/aws/ml/deeplens.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Polly(opts ...attr.Attribute) *node.Node {
	return node.New("polly", attr.AssetImage("assets/aws/ml/polly.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) SagemakerGroundTruth(opts ...attr.Attribute) *node.Node {
	return node.New("sagemaker-ground-truth", attr.AssetImage("assets/aws/ml/sagemaker-ground-truth.png"), attr.Shape(attr.None))
}

func (c *awsMlContainer) Comprehend(opts ...attr.Attribute) *node.Node {
	return node.New("comprehend", attr.AssetImage("assets/aws/ml/comprehend.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.ml"
  Register(namespace, "Deepracer", AWSMl.Deepracer)
  Register(namespace, "Forecast", AWSMl.Forecast)
  Register(namespace, "SagemakerNotebook", AWSMl.SagemakerNotebook)
  Register(namespace, "Sagemaker", AWSMl.Sagemaker)
  Register(namespace, "TensorflowOnAws", AWSMl.TensorflowOnAws)
  Register(namespace, "DeepLearningAmis", AWSMl.DeepLearningAmis)
  Register(namespace, "MachineLearning", AWSMl.MachineLearning)
  Register(namespace, "Personalize", AWSMl.Personalize)
  Register(namespace, "Rekognition", AWSMl.Rekognition)
  Register(namespace, "SagemakerTrainingJob", AWSMl.SagemakerTrainingJob)
  Register(namespace, "Translate", AWSMl.Translate)
  Register(namespace, "Lex", AWSMl.Lex)
  Register(namespace, "DeepLearningContainers", AWSMl.DeepLearningContainers)
  Register(namespace, "ElasticInference", AWSMl.ElasticInference)
  Register(namespace, "SagemakerModel", AWSMl.SagemakerModel)
  Register(namespace, "Textract", AWSMl.Textract)
  Register(namespace, "Transcribe", AWSMl.Transcribe)
  Register(namespace, "ApacheMxnetOnAws", AWSMl.ApacheMxnetOnAws)
  Register(namespace, "Deeplens", AWSMl.Deeplens)
  Register(namespace, "Polly", AWSMl.Polly)
  Register(namespace, "SagemakerGroundTruth", AWSMl.SagemakerGroundTruth)
  Register(namespace, "Comprehend", AWSMl.Comprehend)
}
