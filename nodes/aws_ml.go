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
	opts = append(opts, attr.AssetImage("assets/aws/ml/deepracer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("deepracer", opts...)
}

func (c *awsMlContainer) Forecast(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/forecast.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("forecast", opts...)
}

func (c *awsMlContainer) SagemakerNotebook(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/sagemaker-notebook.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sagemaker-notebook", opts...)
}

func (c *awsMlContainer) Sagemaker(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/sagemaker.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sagemaker", opts...)
}

func (c *awsMlContainer) TensorflowOnAws(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/tensorflow-on-aws.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tensorflow-on-aws", opts...)
}

func (c *awsMlContainer) DeepLearningAmis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/deep-learning-amis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("deep-learning-amis", opts...)
}

func (c *awsMlContainer) MachineLearning(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/machine-learning.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("machine-learning", opts...)
}

func (c *awsMlContainer) Personalize(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/personalize.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("personalize", opts...)
}

func (c *awsMlContainer) Rekognition(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/rekognition.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rekognition", opts...)
}

func (c *awsMlContainer) SagemakerTrainingJob(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/sagemaker-training-job.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sagemaker-training-job", opts...)
}

func (c *awsMlContainer) Translate(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/translate.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("translate", opts...)
}

func (c *awsMlContainer) Lex(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/lex.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("lex", opts...)
}

func (c *awsMlContainer) DeepLearningContainers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/deep-learning-containers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("deep-learning-containers", opts...)
}

func (c *awsMlContainer) ElasticInference(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/elastic-inference.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-inference", opts...)
}

func (c *awsMlContainer) SagemakerModel(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/sagemaker-model.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sagemaker-model", opts...)
}

func (c *awsMlContainer) Textract(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/textract.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("textract", opts...)
}

func (c *awsMlContainer) Transcribe(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/transcribe.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("transcribe", opts...)
}

func (c *awsMlContainer) ApacheMxnetOnAws(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/apache-mxnet-on-aws.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apache-mxnet-on-aws", opts...)
}

func (c *awsMlContainer) Deeplens(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/deeplens.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("deeplens", opts...)
}

func (c *awsMlContainer) Polly(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/polly.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("polly", opts...)
}

func (c *awsMlContainer) SagemakerGroundTruth(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/sagemaker-ground-truth.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sagemaker-ground-truth", opts...)
}

func (c *awsMlContainer) Comprehend(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/ml/comprehend.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("comprehend", opts...)
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
