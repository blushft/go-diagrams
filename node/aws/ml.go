package aws

import "github.com/blushft/go-diagrams/node"

type mlContainer struct {
	path string
	opts []node.Option
}

var Ml = &mlContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/ml",
}

func (c *mlContainer) Polly(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/polly.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) SagemakerGroundTruth(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/sagemaker-ground-truth.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) SagemakerNotebook(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/sagemaker-notebook.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) SagemakerTrainingJob(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/sagemaker-training-job.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Sagemaker(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/sagemaker.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) TensorflowOnAws(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/tensorflow-on-aws.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Comprehend(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/comprehend.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) ElasticInference(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/elastic-inference.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Textract(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/textract.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Translate(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/translate.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Deeplens(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/deeplens.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) MachineLearning(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/machine-learning.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Personalize(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/personalize.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Rekognition(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/rekognition.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Forecast(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/forecast.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Lex(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/lex.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) DeepLearningContainers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/deep-learning-containers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Deepracer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/deepracer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) SagemakerModel(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/sagemaker-model.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) Transcribe(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/transcribe.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) ApacheMxnetOnAws(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/apache-mxnet-on-aws.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mlContainer) DeepLearningAmis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/ml/deep-learning-amis.png")}, c.opts, opts)
	return node.New(nopts...)
}
