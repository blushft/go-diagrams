package aws

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/compute",
}

func (c *computeContainer) ThinkboxDeadlineRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-deadline-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxDeadline(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-deadline.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxStokeRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-stoke-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxStoke(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-stoke.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Ec2ContainerRegistryRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/ec2-container-registry-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Ec2ContainerRegistry(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/ec2-container-registry.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticKubernetesServiceRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/elastic-kubernetes-service-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) FargateRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/fargate-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) VmwareCloudOnAws(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/vmware-cloud-on-aws.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticContainerService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/elastic-container-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxFrostRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-frost-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxKrakatoaRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-krakatoa-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxXmesh(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-xmesh.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) OutpostsRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/outposts-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ServerlessApplicationRepositoryRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/serverless-application-repository-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxKrakatoa(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-krakatoa.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ApplicationAutoScalingRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/application-auto-scaling-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ComputeRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/compute-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Ec2Rounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/ec2-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) LambdaRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/lambda-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticBeanstalk(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/elastic-beanstalk.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticContainerServiceRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/elastic-container-service-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Outposts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/outposts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxSequoiaRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-sequoia-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxSequoia(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-sequoia.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Lambda(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/lambda.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) LightsailRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/lightsail-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ServerlessApplicationRepository(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/serverless-application-repository.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxDraftRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-draft-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ApplicationAutoScaling(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/application-auto-scaling.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxFrost(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-frost.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxDraft(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-draft.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ThinkboxXmeshRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/thinkbox-xmesh-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) VmwareCloudOnAwsRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/vmware-cloud-on-aws-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) BatchRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/batch-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Ec2(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/ec2.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Fargate(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/fargate.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Lightsail(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/lightsail.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Batch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/batch.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Compute(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/compute.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticBeanstalkRounded(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/elastic-beanstalk-rounded.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticKubernetesService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/compute/elastic-kubernetes-service.png")}, c.opts, opts)
	return node.New(nopts...)
}
