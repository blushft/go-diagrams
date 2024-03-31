package oci

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/compute",
}

func (c *computeContainer) AutoscaleWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/autoscale-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) InstancePoolsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/instance-pools-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) InstancePools(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/instance-pools.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) OcirWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/ocir-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Autoscale(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/autoscale.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) BmWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/bm-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) FunctionsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/functions-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Functions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/functions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ocir(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/ocir.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Vm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/vm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) OkeWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/oke-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Oke(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/oke.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) VmWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/vm-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Bm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/bm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ContainerWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/container-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Container(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/compute/container.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
