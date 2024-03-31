package aws

import "github.com/blushft/go-diagrams/diagram"

type mlContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ml = &mlContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/ml",
}

func (c *mlContainer) ApacheMxnetOnAws(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/apache-mxnet-on-aws.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) ElasticInference(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/elastic-inference.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Lex(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/lex.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Personalize(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/personalize.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) SagemakerTrainingJob(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/sagemaker-training-job.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) MachineLearning(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/machine-learning.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) SagemakerNotebook(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/sagemaker-notebook.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) SagemakerModel(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/sagemaker-model.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Sagemaker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/sagemaker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Transcribe(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/transcribe.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) DeepLearningContainers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/deep-learning-containers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Deepracer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/deepracer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Forecast(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/forecast.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Polly(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/polly.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) SagemakerGroundTruth(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/sagemaker-ground-truth.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Textract(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/textract.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Translate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/translate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Comprehend(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/comprehend.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) DeepLearningAmis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/deep-learning-amis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Deeplens(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/deeplens.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Rekognition(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/rekognition.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) TensorflowOnAws(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/ml/tensorflow-on-aws.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
