package oci

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/compute",
}

func (c *computeContainer) InstancePools(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/instance-pools.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) OkeWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/oke-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) AutoscaleWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/autoscale-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) FunctionsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/functions-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Ocir(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/ocir.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Oke(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/oke.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) VmWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/vm-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Autoscale(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/autoscale.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Bm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/bm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) BmWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/bm-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Vm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/vm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Functions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/functions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) InstancePoolsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/instance-pools-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) OcirWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/ocir-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ContainerWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/container-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Container(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/compute/container.png")}, c.opts, opts)
	return node.New(nopts...)
}
