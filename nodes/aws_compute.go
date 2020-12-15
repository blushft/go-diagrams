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
	opts = append(opts, attr.AssetImage("assets/aws/compute/lambda.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("lambda", opts...)
}

func (c *awsComputeContainer) Ec2ContainerRegistry(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/ec2-container-registry.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ec2-container-registry", opts...)
}

func (c *awsComputeContainer) ElasticContainerService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/elastic-container-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-container-service", opts...)
}

func (c *awsComputeContainer) LightsailRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/lightsail-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("lightsail-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxDeadline(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-deadline.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-deadline", opts...)
}

func (c *awsComputeContainer) ThinkboxDraft(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-draft.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-draft", opts...)
}

func (c *awsComputeContainer) ApplicationAutoScaling(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/application-auto-scaling.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("application-auto-scaling", opts...)
}

func (c *awsComputeContainer) Batch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/batch.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("batch", opts...)
}

func (c *awsComputeContainer) OutpostsRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/outposts-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("outposts-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxSequoia(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-sequoia.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-sequoia", opts...)
}

func (c *awsComputeContainer) ElasticBeanstalk(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/elastic-beanstalk.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-beanstalk", opts...)
}

func (c *awsComputeContainer) Fargate(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/fargate.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fargate", opts...)
}

func (c *awsComputeContainer) ThinkboxStokeRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-stoke-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-stoke-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxXmesh(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-xmesh.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-xmesh", opts...)
}

func (c *awsComputeContainer) ServerlessApplicationRepository(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/serverless-application-repository.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("serverless-application-repository", opts...)
}

func (c *awsComputeContainer) ThinkboxFrost(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-frost.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-frost", opts...)
}

func (c *awsComputeContainer) ElasticContainerServiceRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/elastic-container-service-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-container-service-rounded", opts...)
}

func (c *awsComputeContainer) FargateRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/fargate-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fargate-rounded", opts...)
}

func (c *awsComputeContainer) Lightsail(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/lightsail.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("lightsail", opts...)
}

func (c *awsComputeContainer) ThinkboxDeadlineRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-deadline-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-deadline-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxKrakatoa(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-krakatoa.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-krakatoa", opts...)
}

func (c *awsComputeContainer) Ec2Rounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/ec2-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ec2-rounded", opts...)
}

func (c *awsComputeContainer) ElasticBeanstalkRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/elastic-beanstalk-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-beanstalk-rounded", opts...)
}

func (c *awsComputeContainer) ElasticKubernetesServiceRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/elastic-kubernetes-service-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-kubernetes-service-rounded", opts...)
}

func (c *awsComputeContainer) ElasticKubernetesService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/elastic-kubernetes-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-kubernetes-service", opts...)
}

func (c *awsComputeContainer) ThinkboxSequoiaRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-sequoia-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-sequoia-rounded", opts...)
}

func (c *awsComputeContainer) VmwareCloudOnAws(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/vmware-cloud-on-aws.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vmware-cloud-on-aws", opts...)
}

func (c *awsComputeContainer) BatchRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/batch-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("batch-rounded", opts...)
}

func (c *awsComputeContainer) Ec2ContainerRegistryRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/ec2-container-registry-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ec2-container-registry-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxFrostRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-frost-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-frost-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxXmeshRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-xmesh-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-xmesh-rounded", opts...)
}

func (c *awsComputeContainer) VmwareCloudOnAwsRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/vmware-cloud-on-aws-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vmware-cloud-on-aws-rounded", opts...)
}

func (c *awsComputeContainer) ComputeRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/compute-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("compute-rounded", opts...)
}

func (c *awsComputeContainer) Compute(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/compute.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("compute", opts...)
}

func (c *awsComputeContainer) LambdaRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/lambda-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("lambda-rounded", opts...)
}

func (c *awsComputeContainer) Outposts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/outposts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("outposts", opts...)
}

func (c *awsComputeContainer) ServerlessApplicationRepositoryRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/serverless-application-repository-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("serverless-application-repository-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxDraftRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-draft-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-draft-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxKrakatoaRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-krakatoa-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-krakatoa-rounded", opts...)
}

func (c *awsComputeContainer) ThinkboxStoke(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/thinkbox-stoke.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("thinkbox-stoke", opts...)
}

func (c *awsComputeContainer) ApplicationAutoScalingRounded(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/application-auto-scaling-rounded.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("application-auto-scaling-rounded", opts...)
}

func (c *awsComputeContainer) Ec2(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/compute/ec2.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ec2", opts...)
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
