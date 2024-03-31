package gcp

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/compute",
}

func (c *computeContainer) KubernetesEngine(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/kubernetes-engine.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Run(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/run.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) AppEngine(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/app-engine.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ComputeEngine(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/compute-engine.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ContainerOptimizedOs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/container-optimized-os.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Functions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/functions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) GkeOnPrem(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/gke-on-prem.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Gpu(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/compute/gpu.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
