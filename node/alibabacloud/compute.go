package alibabacloud

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/compute",
}

func (c *computeContainer) ElasticHighPerformanceComputing(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/elastic-high-performance-computing.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) OperationOrchestrationService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/operation-orchestration-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) WebAppService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/web-app-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) BatchCompute(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/batch-compute.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ContainerService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/container-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticComputeService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/elastic-compute-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) FunctionCompute(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/function-compute.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ServerLoadBalancer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/server-load-balancer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) SimpleApplicationServer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/simple-application-server.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) AutoScaling(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/auto-scaling.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ContainerRegistry(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/container-registry.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ResourceOrchestrationService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/resource-orchestration-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticContainerInstance(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/elastic-container-instance.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ElasticSearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/elastic-search.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ServerlessAppEngine(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/compute/serverless-app-engine.png")}, c.opts, opts)
	return node.New(nopts...)
}
