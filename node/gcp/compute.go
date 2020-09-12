package gcp

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/compute",
}

func (c *computeContainer) ContainerOptimizedOs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/container-optimized-os.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Functions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/functions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) GkeOnPrem(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/gke-on-prem.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Gpu(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/gpu.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) KubernetesEngine(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/kubernetes-engine.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Run(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/run.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) AppEngine(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/app-engine.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ComputeEngine(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/compute/compute-engine.png")}, c.opts, opts)
	return node.New(nopts...)
}
