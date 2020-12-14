package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSCompute = &awsComputeContainer{path: "assets/aws/compute"}

func (c *awsComputeContainer) Lambda(opts ...attr.Attribute) *node.Node {
	return node.New("lambda", attr.AssetImage("assets/aws/compute/lambda.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Ec2ContainerRegistry(opts ...attr.Attribute) *node.Node {
	return node.New("ec2-container-registry", attr.AssetImage("assets/aws/compute/ec2-container-registry.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ElasticContainerService(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-container-service", attr.AssetImage("assets/aws/compute/elastic-container-service.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) LightsailRounded(opts ...attr.Attribute) *node.Node {
	return node.New("lightsail-rounded", attr.AssetImage("assets/aws/compute/lightsail-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxDeadline(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-deadline", attr.AssetImage("assets/aws/compute/thinkbox-deadline.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxDraft(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-draft", attr.AssetImage("assets/aws/compute/thinkbox-draft.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ApplicationAutoScaling(opts ...attr.Attribute) *node.Node {
	return node.New("application-auto-scaling", attr.AssetImage("assets/aws/compute/application-auto-scaling.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Batch(opts ...attr.Attribute) *node.Node {
	return node.New("batch", attr.AssetImage("assets/aws/compute/batch.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) OutpostsRounded(opts ...attr.Attribute) *node.Node {
	return node.New("outposts-rounded", attr.AssetImage("assets/aws/compute/outposts-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxSequoia(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-sequoia", attr.AssetImage("assets/aws/compute/thinkbox-sequoia.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ElasticBeanstalk(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-beanstalk", attr.AssetImage("assets/aws/compute/elastic-beanstalk.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Fargate(opts ...attr.Attribute) *node.Node {
	return node.New("fargate", attr.AssetImage("assets/aws/compute/fargate.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxStokeRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-stoke-rounded", attr.AssetImage("assets/aws/compute/thinkbox-stoke-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxXmesh(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-xmesh", attr.AssetImage("assets/aws/compute/thinkbox-xmesh.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ServerlessApplicationRepository(opts ...attr.Attribute) *node.Node {
	return node.New("serverless-application-repository", attr.AssetImage("assets/aws/compute/serverless-application-repository.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxFrost(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-frost", attr.AssetImage("assets/aws/compute/thinkbox-frost.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ElasticContainerServiceRounded(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-container-service-rounded", attr.AssetImage("assets/aws/compute/elastic-container-service-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) FargateRounded(opts ...attr.Attribute) *node.Node {
	return node.New("fargate-rounded", attr.AssetImage("assets/aws/compute/fargate-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Lightsail(opts ...attr.Attribute) *node.Node {
	return node.New("lightsail", attr.AssetImage("assets/aws/compute/lightsail.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxDeadlineRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-deadline-rounded", attr.AssetImage("assets/aws/compute/thinkbox-deadline-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxKrakatoa(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-krakatoa", attr.AssetImage("assets/aws/compute/thinkbox-krakatoa.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Ec2Rounded(opts ...attr.Attribute) *node.Node {
	return node.New("ec2-rounded", attr.AssetImage("assets/aws/compute/ec2-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ElasticBeanstalkRounded(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-beanstalk-rounded", attr.AssetImage("assets/aws/compute/elastic-beanstalk-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ElasticKubernetesServiceRounded(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-kubernetes-service-rounded", attr.AssetImage("assets/aws/compute/elastic-kubernetes-service-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ElasticKubernetesService(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-kubernetes-service", attr.AssetImage("assets/aws/compute/elastic-kubernetes-service.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxSequoiaRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-sequoia-rounded", attr.AssetImage("assets/aws/compute/thinkbox-sequoia-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) VmwareCloudOnAws(opts ...attr.Attribute) *node.Node {
	return node.New("vmware-cloud-on-aws", attr.AssetImage("assets/aws/compute/vmware-cloud-on-aws.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) BatchRounded(opts ...attr.Attribute) *node.Node {
	return node.New("batch-rounded", attr.AssetImage("assets/aws/compute/batch-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Ec2ContainerRegistryRounded(opts ...attr.Attribute) *node.Node {
	return node.New("ec2-container-registry-rounded", attr.AssetImage("assets/aws/compute/ec2-container-registry-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxFrostRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-frost-rounded", attr.AssetImage("assets/aws/compute/thinkbox-frost-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxXmeshRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-xmesh-rounded", attr.AssetImage("assets/aws/compute/thinkbox-xmesh-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) VmwareCloudOnAwsRounded(opts ...attr.Attribute) *node.Node {
	return node.New("vmware-cloud-on-aws-rounded", attr.AssetImage("assets/aws/compute/vmware-cloud-on-aws-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ComputeRounded(opts ...attr.Attribute) *node.Node {
	return node.New("compute-rounded", attr.AssetImage("assets/aws/compute/compute-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Compute(opts ...attr.Attribute) *node.Node {
	return node.New("compute", attr.AssetImage("assets/aws/compute/compute.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) LambdaRounded(opts ...attr.Attribute) *node.Node {
	return node.New("lambda-rounded", attr.AssetImage("assets/aws/compute/lambda-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Outposts(opts ...attr.Attribute) *node.Node {
	return node.New("outposts", attr.AssetImage("assets/aws/compute/outposts.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ServerlessApplicationRepositoryRounded(opts ...attr.Attribute) *node.Node {
	return node.New("serverless-application-repository-rounded", attr.AssetImage("assets/aws/compute/serverless-application-repository-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxDraftRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-draft-rounded", attr.AssetImage("assets/aws/compute/thinkbox-draft-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxKrakatoaRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-krakatoa-rounded", attr.AssetImage("assets/aws/compute/thinkbox-krakatoa-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ThinkboxStoke(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-stoke", attr.AssetImage("assets/aws/compute/thinkbox-stoke.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) ApplicationAutoScalingRounded(opts ...attr.Attribute) *node.Node {
	return node.New("application-auto-scaling-rounded", attr.AssetImage("assets/aws/compute/application-auto-scaling-rounded.png"), attr.Shape(attr.None))
}

func (c *awsComputeContainer) Ec2(opts ...attr.Attribute) *node.Node {
	return node.New("ec2", attr.AssetImage("assets/aws/compute/ec2.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.compute"
  Register(namespace, "Lambda", AWSCompute.Lambda)
  Register(namespace, "Ec2ContainerRegistry", AWSCompute.Ec2ContainerRegistry)
  Register(namespace, "ElasticContainerService", AWSCompute.ElasticContainerService)
  Register(namespace, "LightsailRounded", AWSCompute.LightsailRounded)
  Register(namespace, "ThinkboxDeadline", AWSCompute.ThinkboxDeadline)
  Register(namespace, "ThinkboxDraft", AWSCompute.ThinkboxDraft)
  Register(namespace, "ApplicationAutoScaling", AWSCompute.ApplicationAutoScaling)
  Register(namespace, "Batch", AWSCompute.Batch)
  Register(namespace, "OutpostsRounded", AWSCompute.OutpostsRounded)
  Register(namespace, "ThinkboxSequoia", AWSCompute.ThinkboxSequoia)
  Register(namespace, "ElasticBeanstalk", AWSCompute.ElasticBeanstalk)
  Register(namespace, "Fargate", AWSCompute.Fargate)
  Register(namespace, "ThinkboxStokeRounded", AWSCompute.ThinkboxStokeRounded)
  Register(namespace, "ThinkboxXmesh", AWSCompute.ThinkboxXmesh)
  Register(namespace, "ServerlessApplicationRepository", AWSCompute.ServerlessApplicationRepository)
  Register(namespace, "ThinkboxFrost", AWSCompute.ThinkboxFrost)
  Register(namespace, "ElasticContainerServiceRounded", AWSCompute.ElasticContainerServiceRounded)
  Register(namespace, "FargateRounded", AWSCompute.FargateRounded)
  Register(namespace, "Lightsail", AWSCompute.Lightsail)
  Register(namespace, "ThinkboxDeadlineRounded", AWSCompute.ThinkboxDeadlineRounded)
  Register(namespace, "ThinkboxKrakatoa", AWSCompute.ThinkboxKrakatoa)
  Register(namespace, "Ec2Rounded", AWSCompute.Ec2Rounded)
  Register(namespace, "ElasticBeanstalkRounded", AWSCompute.ElasticBeanstalkRounded)
  Register(namespace, "ElasticKubernetesServiceRounded", AWSCompute.ElasticKubernetesServiceRounded)
  Register(namespace, "ElasticKubernetesService", AWSCompute.ElasticKubernetesService)
  Register(namespace, "ThinkboxSequoiaRounded", AWSCompute.ThinkboxSequoiaRounded)
  Register(namespace, "VmwareCloudOnAws", AWSCompute.VmwareCloudOnAws)
  Register(namespace, "BatchRounded", AWSCompute.BatchRounded)
  Register(namespace, "Ec2ContainerRegistryRounded", AWSCompute.Ec2ContainerRegistryRounded)
  Register(namespace, "ThinkboxFrostRounded", AWSCompute.ThinkboxFrostRounded)
  Register(namespace, "ThinkboxXmeshRounded", AWSCompute.ThinkboxXmeshRounded)
  Register(namespace, "VmwareCloudOnAwsRounded", AWSCompute.VmwareCloudOnAwsRounded)
  Register(namespace, "ComputeRounded", AWSCompute.ComputeRounded)
  Register(namespace, "Compute", AWSCompute.Compute)
  Register(namespace, "LambdaRounded", AWSCompute.LambdaRounded)
  Register(namespace, "Outposts", AWSCompute.Outposts)
  Register(namespace, "ServerlessApplicationRepositoryRounded", AWSCompute.ServerlessApplicationRepositoryRounded)
  Register(namespace, "ThinkboxDraftRounded", AWSCompute.ThinkboxDraftRounded)
  Register(namespace, "ThinkboxKrakatoaRounded", AWSCompute.ThinkboxKrakatoaRounded)
  Register(namespace, "ThinkboxStoke", AWSCompute.ThinkboxStoke)
  Register(namespace, "ApplicationAutoScalingRounded", AWSCompute.ApplicationAutoScalingRounded)
  Register(namespace, "Ec2", AWSCompute.Ec2)
}
