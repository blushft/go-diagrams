package k8s

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/compute",
}

func (c *computeContainer) Cronjob(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/compute/cronjob.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Deploy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/compute/deploy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ds(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/compute/ds.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Job(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/compute/job.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Pod(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/compute/pod.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Rs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/compute/rs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Sts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/compute/sts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
