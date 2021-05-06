package alibabacloud

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/alibabacloud/compute"}

func (c *computeContainer) AutoScaling(opts ...attr.Attribute) *node.Node {
	return node.New("auto-scaling", attr.AssetImage("assets/alibabacloud/compute/auto-scaling.png"), attr.Shape(attr.None))
}

func (c *computeContainer) FunctionCompute(opts ...attr.Attribute) *node.Node {
	return node.New("function-compute", attr.AssetImage("assets/alibabacloud/compute/function-compute.png"), attr.Shape(attr.None))
}

func (c *computeContainer) OperationOrchestrationService(opts ...attr.Attribute) *node.Node {
	return node.New("operation-orchestration-service", attr.AssetImage("assets/alibabacloud/compute/operation-orchestration-service.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ServerLoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("server-load-balancer", attr.AssetImage("assets/alibabacloud/compute/server-load-balancer.png"), attr.Shape(attr.None))
}

func (c *computeContainer) WebAppService(opts ...attr.Attribute) *node.Node {
	return node.New("web-app-service", attr.AssetImage("assets/alibabacloud/compute/web-app-service.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ContainerService(opts ...attr.Attribute) *node.Node {
	return node.New("container-service", attr.AssetImage("assets/alibabacloud/compute/container-service.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticSearch(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-search", attr.AssetImage("assets/alibabacloud/compute/elastic-search.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ServerlessAppEngine(opts ...attr.Attribute) *node.Node {
	return node.New("serverless-app-engine", attr.AssetImage("assets/alibabacloud/compute/serverless-app-engine.png"), attr.Shape(attr.None))
}

func (c *computeContainer) SimpleApplicationServer(opts ...attr.Attribute) *node.Node {
	return node.New("simple-application-server", attr.AssetImage("assets/alibabacloud/compute/simple-application-server.png"), attr.Shape(attr.None))
}

func (c *computeContainer) BatchCompute(opts ...attr.Attribute) *node.Node {
	return node.New("batch-compute", attr.AssetImage("assets/alibabacloud/compute/batch-compute.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ContainerRegistry(opts ...attr.Attribute) *node.Node {
	return node.New("container-registry", attr.AssetImage("assets/alibabacloud/compute/container-registry.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticComputeService(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-compute-service", attr.AssetImage("assets/alibabacloud/compute/elastic-compute-service.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ResourceOrchestrationService(opts ...attr.Attribute) *node.Node {
	return node.New("resource-orchestration-service", attr.AssetImage("assets/alibabacloud/compute/resource-orchestration-service.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticContainerInstance(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-container-instance", attr.AssetImage("assets/alibabacloud/compute/elastic-container-instance.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ElasticHighPerformanceComputing(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-high-performance-computing", attr.AssetImage("assets/alibabacloud/compute/elastic-high-performance-computing.png"), attr.Shape(attr.None))
}
