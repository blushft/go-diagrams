package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/compute",
}

func (c *computeContainer) AutoScaling(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/auto-scaling.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ContainerRegistry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/container-registry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticComputeService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/elastic-compute-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) FunctionCompute(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/function-compute.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) OperationOrchestrationService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/operation-orchestration-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ServerLoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/server-load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ServerlessAppEngine(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/serverless-app-engine.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ContainerService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/container-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticHighPerformanceComputing(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/elastic-high-performance-computing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticSearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/elastic-search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ResourceOrchestrationService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/resource-orchestration-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) BatchCompute(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/batch-compute.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticContainerInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/elastic-container-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) SimpleApplicationServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/simple-application-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) WebAppService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/compute/web-app-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
