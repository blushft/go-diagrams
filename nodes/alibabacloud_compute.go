package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabaCloudCompute = &alibabaCloudComputeContainer{path: "assets/alibabacloud/compute"}

func (c *alibabaCloudComputeContainer) OperationOrchestrationService(opts ...attr.Attribute) *node.Node {
	return node.New("operation-orchestration-service", attr.AssetImage("assets/alibabacloud/compute/operation-orchestration-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) WebAppService(opts ...attr.Attribute) *node.Node {
	return node.New("web-app-service", attr.AssetImage("assets/alibabacloud/compute/web-app-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ContainerService(opts ...attr.Attribute) *node.Node {
	return node.New("container-service", attr.AssetImage("assets/alibabacloud/compute/container-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ServerlessAppEngine(opts ...attr.Attribute) *node.Node {
	return node.New("serverless-app-engine", attr.AssetImage("assets/alibabacloud/compute/serverless-app-engine.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) AutoScaling(opts ...attr.Attribute) *node.Node {
	return node.New("auto-scaling", attr.AssetImage("assets/alibabacloud/compute/auto-scaling.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) BatchCompute(opts ...attr.Attribute) *node.Node {
	return node.New("batch-compute", attr.AssetImage("assets/alibabacloud/compute/batch-compute.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) FunctionCompute(opts ...attr.Attribute) *node.Node {
	return node.New("function-compute", attr.AssetImage("assets/alibabacloud/compute/function-compute.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ResourceOrchestrationService(opts ...attr.Attribute) *node.Node {
	return node.New("resource-orchestration-service", attr.AssetImage("assets/alibabacloud/compute/resource-orchestration-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ContainerRegistry(opts ...attr.Attribute) *node.Node {
	return node.New("container-registry", attr.AssetImage("assets/alibabacloud/compute/container-registry.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ElasticComputeService(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-compute-service", attr.AssetImage("assets/alibabacloud/compute/elastic-compute-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ElasticContainerInstance(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-container-instance", attr.AssetImage("assets/alibabacloud/compute/elastic-container-instance.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ElasticHighPerformanceComputing(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-high-performance-computing", attr.AssetImage("assets/alibabacloud/compute/elastic-high-performance-computing.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ElasticSearch(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-search", attr.AssetImage("assets/alibabacloud/compute/elastic-search.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) ServerLoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("server-load-balancer", attr.AssetImage("assets/alibabacloud/compute/server-load-balancer.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudComputeContainer) SimpleApplicationServer(opts ...attr.Attribute) *node.Node {
	return node.New("simple-application-server", attr.AssetImage("assets/alibabacloud/compute/simple-application-server.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "alibabacloud.compute"
	Register(namespace, "OperationOrchestrationService", AlibabaCloudCompute.OperationOrchestrationService)
	Register(namespace, "WebAppService", AlibabaCloudCompute.WebAppService)
	Register(namespace, "ContainerService", AlibabaCloudCompute.ContainerService)
	Register(namespace, "ServerlessAppEngine", AlibabaCloudCompute.ServerlessAppEngine)
	Register(namespace, "AutoScaling", AlibabaCloudCompute.AutoScaling)
	Register(namespace, "BatchCompute", AlibabaCloudCompute.BatchCompute)
	Register(namespace, "FunctionCompute", AlibabaCloudCompute.FunctionCompute)
	Register(namespace, "ResourceOrchestrationService", AlibabaCloudCompute.ResourceOrchestrationService)
	Register(namespace, "ContainerRegistry", AlibabaCloudCompute.ContainerRegistry)
	Register(namespace, "ElasticComputeService", AlibabaCloudCompute.ElasticComputeService)
	Register(namespace, "ElasticContainerInstance", AlibabaCloudCompute.ElasticContainerInstance)
	Register(namespace, "ElasticHighPerformanceComputing", AlibabaCloudCompute.ElasticHighPerformanceComputing)
	Register(namespace, "ElasticSearch", AlibabaCloudCompute.ElasticSearch)
	Register(namespace, "ServerLoadBalancer", AlibabaCloudCompute.ServerLoadBalancer)
	Register(namespace, "SimpleApplicationServer", AlibabaCloudCompute.SimpleApplicationServer)

}