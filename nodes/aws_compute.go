package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/aws/compute"}

func (c *computeContainer) Lambda(opts ...attr.Attribute) *node.Node {
	return node.New("lambda", attr.AssetImage("assets/aws/compute/lambda.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Ec2ContainerRegistry(opts ...attr.Attribute) *node.Node {
	return node.New("ec2-container-registry", attr.AssetImage("assets/aws/compute/ec2-container-registry.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticContainerService(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-container-service", attr.AssetImage("assets/aws/compute/elastic-container-service.png"), attr.Shape(attr.None))
}

func (c *computeContainer) LightsailRounded(opts ...attr.Attribute) *node.Node {
	return node.New("lightsail-rounded", attr.AssetImage("assets/aws/compute/lightsail-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxDeadline(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-deadline", attr.AssetImage("assets/aws/compute/thinkbox-deadline.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxDraft(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-draft", attr.AssetImage("assets/aws/compute/thinkbox-draft.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ApplicationAutoScaling(opts ...attr.Attribute) *node.Node {
	return node.New("application-auto-scaling", attr.AssetImage("assets/aws/compute/application-auto-scaling.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Batch(opts ...attr.Attribute) *node.Node {
	return node.New("batch", attr.AssetImage("assets/aws/compute/batch.png"), attr.Shape(attr.None))
}

func (c *computeContainer) OutpostsRounded(opts ...attr.Attribute) *node.Node {
	return node.New("outposts-rounded", attr.AssetImage("assets/aws/compute/outposts-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxSequoia(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-sequoia", attr.AssetImage("assets/aws/compute/thinkbox-sequoia.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticBeanstalk(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-beanstalk", attr.AssetImage("assets/aws/compute/elastic-beanstalk.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Fargate(opts ...attr.Attribute) *node.Node {
	return node.New("fargate", attr.AssetImage("assets/aws/compute/fargate.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxStokeRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-stoke-rounded", attr.AssetImage("assets/aws/compute/thinkbox-stoke-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxXmesh(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-xmesh", attr.AssetImage("assets/aws/compute/thinkbox-xmesh.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ServerlessApplicationRepository(opts ...attr.Attribute) *node.Node {
	return node.New("serverless-application-repository", attr.AssetImage("assets/aws/compute/serverless-application-repository.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxFrost(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-frost", attr.AssetImage("assets/aws/compute/thinkbox-frost.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticContainerServiceRounded(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-container-service-rounded", attr.AssetImage("assets/aws/compute/elastic-container-service-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) FargateRounded(opts ...attr.Attribute) *node.Node {
	return node.New("fargate-rounded", attr.AssetImage("assets/aws/compute/fargate-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Lightsail(opts ...attr.Attribute) *node.Node {
	return node.New("lightsail", attr.AssetImage("assets/aws/compute/lightsail.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxDeadlineRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-deadline-rounded", attr.AssetImage("assets/aws/compute/thinkbox-deadline-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxKrakatoa(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-krakatoa", attr.AssetImage("assets/aws/compute/thinkbox-krakatoa.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Ec2Rounded(opts ...attr.Attribute) *node.Node {
	return node.New("ec2-rounded", attr.AssetImage("assets/aws/compute/ec2-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticBeanstalkRounded(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-beanstalk-rounded", attr.AssetImage("assets/aws/compute/elastic-beanstalk-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticKubernetesServiceRounded(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-kubernetes-service-rounded", attr.AssetImage("assets/aws/compute/elastic-kubernetes-service-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticKubernetesService(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-kubernetes-service", attr.AssetImage("assets/aws/compute/elastic-kubernetes-service.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxSequoiaRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-sequoia-rounded", attr.AssetImage("assets/aws/compute/thinkbox-sequoia-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) VmwareCloudOnAws(opts ...attr.Attribute) *node.Node {
	return node.New("vmware-cloud-on-aws", attr.AssetImage("assets/aws/compute/vmware-cloud-on-aws.png"), attr.Shape(attr.None))
}

func (c *computeContainer) BatchRounded(opts ...attr.Attribute) *node.Node {
	return node.New("batch-rounded", attr.AssetImage("assets/aws/compute/batch-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Ec2ContainerRegistryRounded(opts ...attr.Attribute) *node.Node {
	return node.New("ec2-container-registry-rounded", attr.AssetImage("assets/aws/compute/ec2-container-registry-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxFrostRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-frost-rounded", attr.AssetImage("assets/aws/compute/thinkbox-frost-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxXmeshRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-xmesh-rounded", attr.AssetImage("assets/aws/compute/thinkbox-xmesh-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) VmwareCloudOnAwsRounded(opts ...attr.Attribute) *node.Node {
	return node.New("vmware-cloud-on-aws-rounded", attr.AssetImage("assets/aws/compute/vmware-cloud-on-aws-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ComputeRounded(opts ...attr.Attribute) *node.Node {
	return node.New("compute-rounded", attr.AssetImage("assets/aws/compute/compute-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Compute(opts ...attr.Attribute) *node.Node {
	return node.New("compute", attr.AssetImage("assets/aws/compute/compute.png"), attr.Shape(attr.None))
}

func (c *computeContainer) LambdaRounded(opts ...attr.Attribute) *node.Node {
	return node.New("lambda-rounded", attr.AssetImage("assets/aws/compute/lambda-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Outposts(opts ...attr.Attribute) *node.Node {
	return node.New("outposts", attr.AssetImage("assets/aws/compute/outposts.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ServerlessApplicationRepositoryRounded(opts ...attr.Attribute) *node.Node {
	return node.New("serverless-application-repository-rounded", attr.AssetImage("assets/aws/compute/serverless-application-repository-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxDraftRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-draft-rounded", attr.AssetImage("assets/aws/compute/thinkbox-draft-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxKrakatoaRounded(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-krakatoa-rounded", attr.AssetImage("assets/aws/compute/thinkbox-krakatoa-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ThinkboxStoke(opts ...attr.Attribute) *node.Node {
	return node.New("thinkbox-stoke", attr.AssetImage("assets/aws/compute/thinkbox-stoke.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ApplicationAutoScalingRounded(opts ...attr.Attribute) *node.Node {
	return node.New("application-auto-scaling-rounded", attr.AssetImage("assets/aws/compute/application-auto-scaling-rounded.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Ec2(opts ...attr.Attribute) *node.Node {
	return node.New("ec2", attr.AssetImage("assets/aws/compute/ec2.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.compute"
  Register(namespace, "Lambda", Compute.Lambda)
  Register(namespace, "Ec2ContainerRegistry", Compute.Ec2ContainerRegistry)
  Register(namespace, "ElasticContainerService", Compute.ElasticContainerService)
  Register(namespace, "LightsailRounded", Compute.LightsailRounded)
  Register(namespace, "ThinkboxDeadline", Compute.ThinkboxDeadline)
  Register(namespace, "ThinkboxDraft", Compute.ThinkboxDraft)
  Register(namespace, "ApplicationAutoScaling", Compute.ApplicationAutoScaling)
  Register(namespace, "Batch", Compute.Batch)
  Register(namespace, "OutpostsRounded", Compute.OutpostsRounded)
  Register(namespace, "ThinkboxSequoia", Compute.ThinkboxSequoia)
  Register(namespace, "ElasticBeanstalk", Compute.ElasticBeanstalk)
  Register(namespace, "Fargate", Compute.Fargate)
  Register(namespace, "ThinkboxStokeRounded", Compute.ThinkboxStokeRounded)
  Register(namespace, "ThinkboxXmesh", Compute.ThinkboxXmesh)
  Register(namespace, "ServerlessApplicationRepository", Compute.ServerlessApplicationRepository)
  Register(namespace, "ThinkboxFrost", Compute.ThinkboxFrost)
  Register(namespace, "ElasticContainerServiceRounded", Compute.ElasticContainerServiceRounded)
  Register(namespace, "FargateRounded", Compute.FargateRounded)
  Register(namespace, "Lightsail", Compute.Lightsail)
  Register(namespace, "ThinkboxDeadlineRounded", Compute.ThinkboxDeadlineRounded)
  Register(namespace, "ThinkboxKrakatoa", Compute.ThinkboxKrakatoa)
  Register(namespace, "Ec2Rounded", Compute.Ec2Rounded)
  Register(namespace, "ElasticBeanstalkRounded", Compute.ElasticBeanstalkRounded)
  Register(namespace, "ElasticKubernetesServiceRounded", Compute.ElasticKubernetesServiceRounded)
  Register(namespace, "ElasticKubernetesService", Compute.ElasticKubernetesService)
  Register(namespace, "ThinkboxSequoiaRounded", Compute.ThinkboxSequoiaRounded)
  Register(namespace, "VmwareCloudOnAws", Compute.VmwareCloudOnAws)
  Register(namespace, "BatchRounded", Compute.BatchRounded)
  Register(namespace, "Ec2ContainerRegistryRounded", Compute.Ec2ContainerRegistryRounded)
  Register(namespace, "ThinkboxFrostRounded", Compute.ThinkboxFrostRounded)
  Register(namespace, "ThinkboxXmeshRounded", Compute.ThinkboxXmeshRounded)
  Register(namespace, "VmwareCloudOnAwsRounded", Compute.VmwareCloudOnAwsRounded)
  Register(namespace, "ComputeRounded", Compute.ComputeRounded)
  Register(namespace, "Compute", Compute.Compute)
  Register(namespace, "LambdaRounded", Compute.LambdaRounded)
  Register(namespace, "Outposts", Compute.Outposts)
  Register(namespace, "ServerlessApplicationRepositoryRounded", Compute.ServerlessApplicationRepositoryRounded)
  Register(namespace, "ThinkboxDraftRounded", Compute.ThinkboxDraftRounded)
  Register(namespace, "ThinkboxKrakatoaRounded", Compute.ThinkboxKrakatoaRounded)
  Register(namespace, "ThinkboxStoke", Compute.ThinkboxStoke)
  Register(namespace, "ApplicationAutoScalingRounded", Compute.ApplicationAutoScalingRounded)
  Register(namespace, "Ec2", Compute.Ec2)
}
