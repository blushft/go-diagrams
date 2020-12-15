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
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/operation-orchestration-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("operation-orchestration-service", opts...)
}

func (c *alibabaCloudComputeContainer) WebAppService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/web-app-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("web-app-service", opts...)
}

func (c *alibabaCloudComputeContainer) ContainerService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/container-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container-service", opts...)
}

func (c *alibabaCloudComputeContainer) ServerlessAppEngine(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/serverless-app-engine.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("serverless-app-engine", opts...)
}

func (c *alibabaCloudComputeContainer) AutoScaling(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/auto-scaling.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("auto-scaling", opts...)
}

func (c *alibabaCloudComputeContainer) BatchCompute(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/batch-compute.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("batch-compute", opts...)
}

func (c *alibabaCloudComputeContainer) FunctionCompute(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/function-compute.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("function-compute", opts...)
}

func (c *alibabaCloudComputeContainer) ResourceOrchestrationService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/resource-orchestration-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("resource-orchestration-service", opts...)
}

func (c *alibabaCloudComputeContainer) ContainerRegistry(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/container-registry.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container-registry", opts...)
}

func (c *alibabaCloudComputeContainer) ElasticComputeService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/elastic-compute-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-compute-service", opts...)
}

func (c *alibabaCloudComputeContainer) ElasticContainerInstance(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/elastic-container-instance.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-container-instance", opts...)
}

func (c *alibabaCloudComputeContainer) ElasticHighPerformanceComputing(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/elastic-high-performance-computing.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-high-performance-computing", opts...)
}

func (c *alibabaCloudComputeContainer) ElasticSearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/elastic-search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-search", opts...)
}

func (c *alibabaCloudComputeContainer) ServerLoadBalancer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/server-load-balancer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("server-load-balancer", opts...)
}

func (c *alibabaCloudComputeContainer) SimpleApplicationServer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/compute/simple-application-server.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("simple-application-server", opts...)
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