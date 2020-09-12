package k8s

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/compute",
}

func (c *computeContainer) Job(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/compute/job.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Pod(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/compute/pod.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Rs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/compute/rs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Sts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/compute/sts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Cronjob(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/compute/cronjob.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Deploy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/compute/deploy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Ds(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/compute/ds.png")}, c.opts, opts)
	return node.New(nopts...)
}
